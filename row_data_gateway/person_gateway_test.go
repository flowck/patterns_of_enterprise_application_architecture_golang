package row_data_gateway_test

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/flowck/patterns_of_enterprise_application_architecture_golang/row_data_gateway"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

var (
	db  *sql.DB
	ctx context.Context
)

func TestPersonGateway_Insert(t *testing.T) {
	createdAtMin := time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC)
	row := row_data_gateway.PersonGateway{
		Id:                 gofakeit.UUID(),
		FirstName:          gofakeit.FirstName(),
		LastName:           gofakeit.LastName(),
		NumberOfDependents: int16(gofakeit.Number(0, 50)),
		CreatedAt:          gofakeit.DateRange(createdAtMin, time.Now()),
		UpdatedAt:          time.Now(),
	}

	require.Nil(t, row.Insert(ctx, db), "require no error during INSERT")
}

func TestMain(m *testing.M) {
	// Random seed data
	gofakeit.Seed(0)

	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Minute*1)
	defer cancel()

	var err error
	db, err = sql.Open("postgres", "postgres://postgres:password@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("could not open a connection to postgres: %v\n", err)
	}

	if err = db.PingContext(ctx); err != nil {
		log.Fatalf("could not ping postgres: %v\n", err)
	}

	os.Exit(m.Run())
}
