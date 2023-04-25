package optimistic_offline_lock_test

import (
	"context"
	"database/sql"
	"log"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flowck/patterns_of_enterprise_application_architecture_golang/optimistic_offline_lock"
)

var (
	db  *sql.DB
	ctx context.Context
)

// Attempt to update two costumers concurrently, thus causing version conflict
func TestService_EditCustomer(t *testing.T) {
	service := optimistic_offline_lock.NewService(db)
	customer1 := fixtureCustomer(t)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	errorCount := atomic.Int32{}
	for i := 0; i < 2; i++ {
		go func(wg *sync.WaitGroup) {
			c, err := service.GetCustomer(ctx, customer1.ID)
			require.Nil(t, err)
			require.Equal(t, customer1.ID, c.ID)
			t.Logf("Customer version: %d", c.Version())

			c.Edit(gofakeit.FirstName(), gofakeit.LastName(), gofakeit.Email())
			simulateLatency(time.Millisecond * 100)

			err = service.EditCustomer(ctx, c)
			if err != nil {
				errorCount.Add(1)
			}

			wg.Done()
		}(wg)
	}

	wg.Wait()

	expectedErrorCount := 1
	assert.Equal(t, expectedErrorCount, int(errorCount.Load()), "expect one of the updated to fail due to version conflict")
}

func simulateLatency(duration time.Duration) {
	time.Sleep(duration)
}

func fixtureCustomer(t *testing.T) *optimistic_offline_lock.Customer {
	c := &optimistic_offline_lock.Customer{
		ID:        gofakeit.UUID(),
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
	}

	_, err := db.ExecContext(
		ctx,
		`INSERT INTO customers (id, first_name, last_name, email) VALUES ($1, $2, $3, $4)`,
		c.ID, c.FirstName, c.LastName, c.Email,
	)
	require.Nil(t, err)

	return c
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
