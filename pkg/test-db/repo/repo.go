package repo

import (
	"context"

	"github.com/go-pg/pg"
)

type wrapper struct {
	db *pg.DB
}

func CreateWrapperRepo(db *pg.DB) *wrapper {
	return &wrapper{db: db}
}

type WrapperRepo interface {
	Load(ctx context.Context) ([]Technology, error)
}

func (l *wrapper) Load(ctx context.Context) ([]Technology, error) {
	db := l.db.WithContext(ctx)

	selectQuery := `
	SELECT *
	FROM   technologies 
	`

	var queryResult []Technology
	_, err := db.Query(&queryResult, selectQuery)
	if err != nil {
		if err != pg.ErrNoRows {
			return nil, err
		}
	}

	return queryResult, nil
}
