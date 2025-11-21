package database

import (
	"context"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/nanoteck137/pyrin/ember"
	"github.com/nanoteck137/storebook/types"
	"github.com/nanoteck137/storebook/utils"
)

type Image struct {
	RowId int `db:"rowid"`

	CollectionId string `db:"collection_id"`
	Hash         string `db:"hash"`

	Filename string `db:"filename"`

	Created int64 `db:"created"`
	Updated int64 `db:"updated"`
}

// TODO(patrik): Use goqu.T more
func ImageQuery() *goqu.SelectDataset {
	query := dialect.From("images").
		Select(
			"images.rowid",

			"images.collection_id",
			"images.hash",

			"images.filename",

			"images.created",
			"images.updated",
		)

	return query
}

func (db DB) GetPagedImages(ctx context.Context, opts FetchOptions) ([]Image, types.Page, error) {
	query := ImageQuery()

	countQuery := query.
		Select(goqu.COUNT("images.id"))

	if opts.PerPage > 0 {
		query = query.
			Limit(uint(opts.PerPage)).
			Offset(uint(opts.Page * opts.PerPage))
	}

	totalItems, err := ember.Single[int](db.db, ctx, countQuery)
	if err != nil {
		return nil, types.Page{}, err
	}

	totalPages := utils.TotalPages(opts.PerPage, totalItems)
	page := types.Page{
		Page:       opts.Page,
		PerPage:    opts.PerPage,
		TotalItems: totalItems,
		TotalPages: totalPages,
	}

	items, err := ember.Multiple[Image](db.db, ctx, query)
	if err != nil {
		return nil, types.Page{}, err
	}

	return items, page, nil
}

func (db DB) GetPagedImagesByCollectionId(ctx context.Context, collectionId string, opts FetchOptions) ([]Image, types.Page, error) {
	query := ImageQuery().
		Where(
			goqu.I("images.collection_id").Eq(collectionId),
		)

	countQuery := query.
		Select(goqu.COUNT("images.id"))

	if opts.PerPage > 0 {
		query = query.
			Limit(uint(opts.PerPage)).
			Offset(uint(opts.Page * opts.PerPage))
	}

	totalItems, err := ember.Single[int](db.db, ctx, countQuery)
	if err != nil {
		return nil, types.Page{}, err
	}

	totalPages := utils.TotalPages(opts.PerPage, totalItems)
	page := types.Page{
		Page:       opts.Page,
		PerPage:    opts.PerPage,
		TotalItems: totalItems,
		TotalPages: totalPages,
	}

	items, err := ember.Multiple[Image](db.db, ctx, query)
	if err != nil {
		return nil, types.Page{}, err
	}

	return items, page, nil
}

func (db DB) GetAllImages(ctx context.Context) ([]Image, error) {
	query := ImageQuery()
	return ember.Multiple[Image](db.db, ctx, query)
}

func (db DB) GetAllImagesByCollectionId(ctx context.Context, collectionId string) ([]Image, error) {
	query := ImageQuery().
		Where(
			goqu.I("images.collection_id").Eq(collectionId),
		).
		Order(
			goqu.I("images.created").Asc(),
		)

	return ember.Multiple[Image](db.db, ctx, query)
}

func (db DB) GetImageById(ctx context.Context, collectionId, id string) (Image, error) {
	query := ImageQuery().
		Where(
			goqu.I("images.collection_id").Eq(collectionId),
			goqu.I("images.id").Eq(id),
		)

	return ember.Single[Image](db.db, ctx, query)
}

type CreateImageParams struct {
	CollectionId string
	Hash         string

	Filename string

	Created int64
	Updated int64
}

func (db DB) CreateImage(ctx context.Context, params CreateImageParams) error {
	t := time.Now().UnixMilli()
	created := params.Created
	updated := params.Updated

	if created == 0 && updated == 0 {
		created = t
		updated = t
	}

	query := dialect.Insert("images").Rows(goqu.Record{
		"collection_id": params.CollectionId,
		"hash":          params.Hash,

		"filename": params.Filename,

		"created": created,
		"updated": updated,
	})

	_, err := db.db.Exec(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

// type ImageChanges struct {
// 	Title Change[string]
//
// 	Created Change[int64]
// }
//
// func (db DB) UpdateImage(ctx context.Context, id string, changes ImageChanges) error {
// 	record := goqu.Record{}
//
// 	addToRecord(record, "title", changes.Title)
//
// 	addToRecord(record, "created", changes.Created)
//
// 	if len(record) == 0 {
// 		return nil
// 	}
//
// 	record["updated"] = time.Now().UnixMilli()
//
// 	query := dialect.Update("images").
// 		Set(record).
// 		Where(goqu.I("images.id").Eq(id))
//
// 	_, err := db.db.Exec(ctx, query)
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }

func (db DB) RemoveImage(ctx context.Context, collectionId, id string) error {
	query := dialect.Delete("images").
		Where(
			goqu.I("images.collection_id").Eq(collectionId),
			goqu.I("images.id").Eq(id),
		)

	_, err := db.db.Exec(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
