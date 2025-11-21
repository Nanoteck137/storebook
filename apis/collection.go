package apis

import (
	"archive/zip"
	"context"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"
	"sort"
	"strconv"

	"github.com/maruel/natural"
	"github.com/nanoteck137/pyrin"
	"github.com/nanoteck137/pyrin/anvil"
	"github.com/nanoteck137/storebook/core"
	"github.com/nanoteck137/storebook/database"
	"github.com/nanoteck137/storebook/types"
	"github.com/nanoteck137/storebook/utils"
	"github.com/nanoteck137/validate"
)

type Collection struct {
	Id string `json:"id"`

	Title string `json:"title"`
}

type GetCollection struct {
	Page        types.Page   `json:"page"`
	Collections []Collection `json:"collections"`
}

type GetCollectionById struct {
	Collection
}

// TODO(patrik): Move
func getPageOptions(q url.Values) database.FetchOptions {
	perPage := 100
	page := 0

	if s := q.Get("perPage"); s != "" {
		i, _ := strconv.Atoi(s)
		if i > 0 {
			perPage = i
		}
	}

	if s := q.Get("page"); s != "" {
		i, _ := strconv.Atoi(s)
		page = i
	}

	return database.FetchOptions{
		PerPage: perPage,
		Page:    page,
	}
}

func ConvertDBCollection(c pyrin.Context, collection database.Collection) Collection {
	return Collection{
		Id:    collection.Id,
		Title: collection.Title,
	}
}

type CreateCollection struct {
	Id string `json:"id"`
}

type CreateCollectionBody struct {
	Title string `json:"title"`
}

func (b *CreateCollectionBody) Transform() {
	b.Title = anvil.String(b.Title)
}

func (b CreateCollectionBody) Validate() error {
	return validate.ValidateStruct(&b,
		validate.Field(&b.Title, validate.Required),
	)
}

type EditCollectionBody struct {
	Title *string `json:"title,omitempty"`
}

func (b *EditCollectionBody) Transform() {
	b.Title = anvil.StringPtr(b.Title)
}

func (b EditCollectionBody) Validate() error {
	return validate.ValidateStruct(&b,
		validate.Field(&b.Title, validate.Required.When(b.Title != nil)),
	)
}

type CollectionImage struct {
	CollectionId string `json:"collectionId"`
	Hash         string `json:"hash"`
	Filename     string `json:"filename"`
	Url          string `json:"url"`
}

func ConvertDBCollectionImage(c pyrin.Context, image database.Image) CollectionImage {
	url := ConvertURL(c, fmt.Sprintf("/files/collections/%s/images/%s", image.CollectionId, image.Filename))

	return CollectionImage{
		CollectionId: image.CollectionId,
		Hash:         image.Hash,
		Filename:     image.Filename,
		Url:          url,
	}
}

type GetCollectionImages struct {
	Images []CollectionImage `json:"images"`
}

func InstallCollectionHandlers(app core.App, group pyrin.Group) {
	group.Register(
		pyrin.ApiHandler{
			Name:         "GetCollections",
			Method:       http.MethodGet,
			Path:         "/collections",
			ResponseType: GetCollection{},
			HandlerFunc: func(c pyrin.Context) (any, error) {
				q := c.Request().URL.Query()
				opts := getPageOptions(q)

				ctx := context.TODO()

				collection, p, err := app.DB().GetPagedCollection(ctx, opts)
				if err != nil {
					return nil, err
				}

				res := GetCollection{
					Page:        p,
					Collections: make([]Collection, len(collection)),
				}

				for i, m := range collection {
					res.Collections[i] = ConvertDBCollection(c, m)
				}

				return res, nil
			},
		},

		pyrin.ApiHandler{
			Name:         "GetCollectionById",
			Method:       http.MethodGet,
			Path:         "/collections/:id",
			ResponseType: GetCollectionById{},
			HandlerFunc: func(c pyrin.Context) (any, error) {
				id := c.Param("id")

				collection, err := app.DB().GetCollectionById(c.Request().Context(), id)
				if err != nil {
					if errors.Is(err, database.ErrItemNotFound) {
						return nil, CollectionNotFound()
					}

					return nil, err
				}

				return GetCollectionById{
					Collection: ConvertDBCollection(c, collection),
				}, nil
			},
		},

		pyrin.ApiHandler{
			Name:         "GetCollectionImages",
			Method:       http.MethodGet,
			Path:         "/collections/:id/images",
			ResponseType: GetCollectionImages{},
			HandlerFunc: func(c pyrin.Context) (any, error) {
				id := c.Param("id")

				ctx := c.Request().Context()

				dbCollection, err := app.DB().GetCollectionById(ctx, id)
				if err != nil {
					if errors.Is(err, database.ErrItemNotFound) {
						return nil, CollectionNotFound()
					}

					return nil, err
				}

				images, err := app.DB().GetAllImagesByCollectionId(ctx, dbCollection.Id)
				if err != nil {
					if errors.Is(err, database.ErrItemNotFound) {
						return nil, CollectionNotFound()
					}

					return nil, err
				}

				res := GetCollectionImages{
					Images: make([]CollectionImage, len(images)),
				}

				for i, image := range images {
					res.Images[i] = ConvertDBCollectionImage(c, image)
				}

				return res, nil
			},
		},

		pyrin.ApiHandler{
			Name:         "CreateCollection",
			Method:       http.MethodPost,
			Path:         "/collections",
			ResponseType: CreateCollection{},
			BodyType:     CreateCollectionBody{},
			HandlerFunc: func(c pyrin.Context) (any, error) {
				body, err := pyrin.Body[CreateCollectionBody](c)
				if err != nil {
					return nil, err
				}

				ctx := context.Background()

				id := utils.CreateCollectionId()

				collectionDir := app.WorkDir().CollectionDirById(id)
				err = collectionDir.Create()
				if err != nil {
					return nil, err
				}

				_, err = app.DB().CreateCollection(ctx, database.CreateCollectionParams{
					Id:    id,
					Title: body.Title,
				})
				if err != nil {
					return nil, err
				}

				return CreateCollection{
					Id: id,
				}, nil
			},
		},

		pyrin.ApiHandler{
			Name:         "EditCollection",
			Method:       http.MethodPatch,
			Path:         "/collections/:id",
			ResponseType: nil,
			BodyType:     EditCollectionBody{},
			HandlerFunc: func(c pyrin.Context) (any, error) {
				id := c.Param("id")

				body, err := pyrin.Body[EditCollectionBody](c)
				if err != nil {
					return nil, err
				}

				ctx := context.Background()

				dbCollection, err := app.DB().GetCollectionById(ctx, id)
				if err != nil {
					if errors.Is(err, database.ErrItemNotFound) {
						return nil, CollectionNotFound()
					}

					return nil, err
				}

				changes := database.CollectionChanges{}

				if body.Title != nil {
					changes.Title = database.Change[string]{
						Value:   *body.Title,
						Changed: *body.Title != dbCollection.Title,
					}
				}

				err = app.DB().UpdateCollection(ctx, dbCollection.Id, changes)
				if err != nil {
					return nil, err
				}

				return nil, nil
			},
		},

		pyrin.ApiHandler{
			Name:         "DeleteCollection",
			Method:       http.MethodDelete,
			Path:         "/collections/:id",
			ResponseType: nil,
			HandlerFunc: func(c pyrin.Context) (any, error) {
				id := c.Param("id")

				ctx := context.Background()

				dbCollection, err := app.DB().GetCollectionById(ctx, id)
				if err != nil {
					if errors.Is(err, database.ErrItemNotFound) {
						return nil, CollectionNotFound()
					}

					return nil, err
				}

				err = app.DB().RemoveCollection(ctx, dbCollection.Id)
				if err != nil {
					return nil, err
				}

				dir := app.WorkDir().CollectionDirById(dbCollection.Id)
				err = os.RemoveAll(dir.String())
				if err != nil {
					return nil, err
				}

				return nil, nil
			},
		},

		pyrin.FormApiHandler{
			Name:   "UploadToCollection",
			Method: http.MethodPost,
			Path:   "/collections/:id/upload",
			Spec: pyrin.FormSpec{
				Files: map[string]pyrin.FormFileSpec{
					"file": {
						NumExpected: 1,
					},
				},
			},
			HandlerFunc: func(c pyrin.Context) (any, error) {
				id := c.Param("id")

				ctx := context.Background()

				dbCollection, err := app.DB().GetCollectionById(ctx, id)
				if err != nil {
					if errors.Is(err, database.ErrItemNotFound) {
						return nil, CollectionNotFound()
					}

					return nil, err
				}

				collectionDir := app.WorkDir().CollectionDirById(dbCollection.Id)
				err = collectionDir.Create()
				if err != nil {
					return nil, err
				}

				files, err := pyrin.FormFiles(c, "file")
				if err != nil {
					return nil, err
				}

				for _, f := range files {
					fmt.Printf("f.Filename: %v\n", f.Filename)
					contentType := f.Header.Get("Content-Type")
					mediaType, _, err := mime.ParseMediaType(contentType)
					if err != nil {
						return nil, err
					}

					readZip := func() error {
						file, err := f.Open()
						if err != nil {
							return err
						}

						r, err := zip.NewReader(file, f.Size)
						if err != nil {
							return err
						}

						sort.SliceStable(r.File, func(i, j int) bool {
							return natural.Less(r.File[i].Name, r.File[j].Name)
						})

						importFile := func(db *database.DB, zf *zip.File) error {
							r, err := zf.Open()
							if err != nil {
								return err
							}

							defer r.Close()

							data, err := io.ReadAll(r)
							if err != nil {
								return err
							}

							ext := path.Ext(zf.Name)
							out, hash, err := utils.WriteHashedFile(data, collectionDir.Images(), ext)
							if err != nil {
								return err
							}

							fmt.Printf("out: %v\n", out)

							err = db.CreateImage(ctx, database.CreateImageParams{
								CollectionId: dbCollection.Id,
								Hash:         hash,
								Filename:     path.Base(out),
							})
							if err != nil {
								return err
							}

							return nil
						}

						for _, zf := range r.File {
							err := importFile(&app.DB().DB, zf)
							if err != nil {
								return err
							}
						}

						return nil
					}

					switch mediaType {
					case "application/zip":
						err := readZip()
						if err != nil {
							return nil, err
						}
					}
				}

				return nil, nil
			},
		},
	)
}
