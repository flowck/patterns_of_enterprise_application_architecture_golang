package row_data_gateway

import (
	"context"
	"database/sql"
	"time"
)

type PersonGateway struct {
	Id                 string
	FirstName          string
	LastName           string
	NumberOfDependents int16
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func (p *PersonGateway) Update(ctx context.Context, db *sql.DB) (int64, error) {
	return 0, nil
}

func (p *PersonGateway) Insert(ctx context.Context, db *sql.DB) error {
	_, err := db.QueryContext(ctx, `
		INSERT INTO people (id, first_name, last_name, number_of_dependents, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		p.Id, p.FirstName, p.LastName, p.NumberOfDependents, p.CreatedAt, p.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
