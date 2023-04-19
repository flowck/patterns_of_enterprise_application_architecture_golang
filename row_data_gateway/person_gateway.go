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
	result, err := db.ExecContext(ctx, `
		UPDATE people
		SET first_name = $2, last_name = $3, number_of_dependents = $4, updated_at = $5
		WHERE id = $1
	`, p.Id, p.FirstName, p.LastName, p.NumberOfDependents, p.UpdatedAt)

	if err != nil {
		return 0, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (p *PersonGateway) Insert(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, `
		INSERT INTO people (id, first_name, last_name, number_of_dependents, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		p.Id, p.FirstName, p.LastName, p.NumberOfDependents, p.CreatedAt, p.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
