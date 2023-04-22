package row_data_gateway_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flowck/patterns_of_enterprise_application_architecture_golang/row_data_gateway"
)

func TestPersonFinder_FindById(t *testing.T) {
	finder := row_data_gateway.NewPersonFinder(db)

	person := fixturePersonGateway()
	require.Nil(t, person.Insert(ctx, db))

	personFound, err := finder.FindById(ctx, person.Id)
	require.Nil(t, err)

	assert.Equal(t, person.Id, personFound.Id)
	assert.Equal(t, person.FirstName, personFound.FirstName)
	assert.Equal(t, person.LastName, personFound.LastName)
	assert.Equal(t, person.NumberOfDependents, personFound.NumberOfDependents)
}

func TestPersonFinder_People(t *testing.T) {
	maxFixtures := 5
	people := make([]*row_data_gateway.PersonGateway, maxFixtures)

	for i := 0; i < maxFixtures; i++ {
		people[i] = fixturePersonGateway()
		require.Nil(t, people[i].Insert(ctx, db))
	}

	finder := row_data_gateway.NewPersonFinder(db)
	peopleFound, err := finder.People(ctx)
	require.Nil(t, err)

	// slow assertion: O^n runtime complexity
	for _, personFound := range peopleFound {
		for _, person := range people {
			if person.Id != personFound.Id {
				continue
			}

			assertPerson(t, person, personFound)
		}
	}
}

func assertPerson(t *testing.T, personA, personB *row_data_gateway.PersonGateway) {
	assert.Equal(t, personA.Id, personB.Id)
	assert.Equal(t, personA.FirstName, personB.FirstName)
	assert.Equal(t, personA.LastName, personB.LastName)
	assert.Equal(t, personA.NumberOfDependents, personB.NumberOfDependents)
}
