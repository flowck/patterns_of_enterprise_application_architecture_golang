package optimistic_offline_lock

import (
	"context"
	"database/sql"
	"errors"
)

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}

func (s *Service) GetCustomer(ctx context.Context, ID string) (*Customer, error) {
	row := s.db.QueryRowContext(
		ctx,
		`SELECT id, first_name, last_name, email, version from customers WHERE id = $1`,
		ID,
	)
	if row.Err() != nil {
		return nil, row.Err()
	}

	c := &Customer{}
	err := row.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Email, &c.version)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *Service) EditCustomer(ctx context.Context, c *Customer) error {
	// The update should only be made if to the row that its id and version match the id and version loaded
	// in the session.
	res, err := s.db.ExecContext(
		ctx,
		`
			UPDATE customers SET first_name = $1, last_name = $2, email = $3, version = $4 + 1
			WHERE id = $5 AND version = $4;
		`,
		c.FirstName, c.LastName, c.Email, c.version, c.ID,
	)
	if err != nil {
		return err
	}

	updatedCount, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if updatedCount == 0 {
		return errors.New("there has been a version conflict")
	}

	return nil
}
