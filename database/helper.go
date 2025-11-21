package database

import (
	"github.com/doug-martin/goqu/v9"
)

type Change[T any] struct {
	Value   T
	Changed bool
}

func addToRecord[T any](record goqu.Record, name string, change Change[T]) {
	if change.Changed {
		record[name] = change.Value
	}
}
