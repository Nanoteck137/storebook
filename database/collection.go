package database

import (
	"context"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/nanoteck137/pyrin/ember"
	"github.com/nanoteck137/storebook/types"
	"github.com/nanoteck137/storebook/utils"
)

type Collection struct {
	RowId int `db:"rowid"`

	Id string `db:"id"`

	Title string `db:"title"`

	Created int64 `db:"created"`
	Updated int64 `db:"updated"`
}

// TODO(patrik): Use goqu.T more
func CollectionQuery() *goqu.SelectDataset {
	query := dialect.From("collections").
		Select(
			"collections.rowid",

			"collections.id",

			"collections.title",

			"collections.created",
			"collections.updated",
		)

	return query
}

// Cleanup
type FetchOptions struct {
	PerPage int
	Page    int
}

func (db DB) GetPagedCollection(ctx context.Context, opts FetchOptions) ([]Collection, types.Page, error) {
	query := CollectionQuery()

	countQuery := query.
		Select(goqu.COUNT("collections.id"))

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

	items, err := ember.Multiple[Collection](db.db, ctx, query)
	if err != nil {
		return nil, types.Page{}, err
	}

	return items, page, nil
}

func (db DB) GetAllCollection(ctx context.Context) ([]Collection, error) {
	query := CollectionQuery()
	return ember.Multiple[Collection](db.db, ctx, query)
}

func (db DB) GetCollectionById(ctx context.Context, id string) (Collection, error) {
	query := CollectionQuery().
		Where(goqu.I("collections.id").Eq(id))

	return ember.Single[Collection](db.db, ctx, query)
}

type CreateCollectionParams struct {
	Id   string

	Title string

	Created int64
	Updated int64
}

func (db DB) CreateCollection(ctx context.Context, params CreateCollectionParams) (string, error) {
	t := time.Now().UnixMilli()
	created := params.Created
	updated := params.Updated

	if created == 0 && updated == 0 {
		created = t
		updated = t
	}

	id := params.Id
	if id == "" {
		id = utils.CreateCollectionId()
	}

	query := dialect.Insert("collections").Rows(goqu.Record{
		"id":   id,

		"title":       params.Title,

		"created": created,
		"updated": updated,
	}).
		Returning("id")

	return ember.Single[string](db.db, ctx, query)
}

type CollectionChanges struct {
	Title       Change[string]

	Created Change[int64]
}

func (db DB) UpdateCollection(ctx context.Context, id string, changes CollectionChanges) error {
	record := goqu.Record{}

	addToRecord(record, "title", changes.Title)

	addToRecord(record, "created", changes.Created)

	if len(record) == 0 {
		return nil
	}

	record["updated"] = time.Now().UnixMilli()

	query := dialect.Update("collections").
		Set(record).
		Where(goqu.I("collections.id").Eq(id))

	_, err := db.db.Exec(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (db DB) RemoveCollection(ctx context.Context, id string) error {
	query := dialect.Delete("collections").
		Where(goqu.I("collections.id").Eq(id))

	_, err := db.db.Exec(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
