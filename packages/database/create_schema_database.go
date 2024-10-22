package database

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/vinibgoulart/gitbook-postgresql-vectorize/packages/content"
	"github.com/vinibgoulart/gitbook-postgresql-vectorize/packages/page"
	"github.com/vinibgoulart/gitbook-postgresql-vectorize/packages/space"
)

func CreateSchemaDatabase(db *bun.DB, ctx context.Context) error {
	models := []interface{}{
		(*space.Space)(nil),
		(*content.Content)(nil),
		(*page.Page)(nil),
	}

	for _, model := range models {
		_, err := db.NewCreateTable().Model(model).IfNotExists().Exec(ctx)

		if err != nil {
			return err
		}
	}

	return nil
}
