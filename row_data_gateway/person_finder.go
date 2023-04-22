package row_data_gateway

import (
	"context"
	"database/sql"
)

type PersonFinder struct {
	db *sql.DB
}

type customRow interface {
	Scan(dest ...any) error
}

func NewPersonFinder(db *sql.DB) PersonFinder {
	return PersonFinder{db: db}
}

func (p PersonFinder) FindById(ctx context.Context, id string) (*PersonGateway, error) {
	row := p.db.QueryRowContext(
		ctx,
		`SELECT id, first_name, last_name, number_of_dependents, created_at, updated_at from people WHERE id = $1`,
		id,
	)
	if row.Err() != nil {
		return nil, row.Err()
	}

	person := &PersonGateway{}
	err := scanToPerson(row, person)
	if err != nil {
		return nil, err
	}

	return person, nil
}

func (p PersonFinder) People(ctx context.Context) ([]*PersonGateway, error) {
	rows, err := p.db.QueryContext(
		ctx,
		`SELECT id, first_name, last_name, number_of_dependents, created_at, updated_at from people`,
	)
	defer func() { rows.Close() }()
	if err != nil {
		return nil, err
	}

	var people []*PersonGateway
	for rows.Next() {
		person := &PersonGateway{}
		err = scanToPerson(rows, person)
		if err != nil {
			return nil, err
		}

		people = append(people, person)
	}

	return people, nil
}

func scanToPerson(row customRow, person *PersonGateway) error {
	return row.Scan(
		&person.Id,
		&person.FirstName,
		&person.LastName,
		&person.NumberOfDependents,
		&person.CreatedAt,
		&person.UpdatedAt,
	)
}
