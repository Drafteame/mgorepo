package mgorepo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Drafteame/mgorepo/driver"
	"github.com/Drafteame/mgorepo/internal/seed"
)

func TestRepository_Count(t *testing.T) {
	d, driverErr := driver.NewTest(t)
	if driverErr != nil {
		t.Fatal(driverErr)
	}

	db := d.Client().Database(d.DbName())

	seed.InsertOne(t, db, collection, testDAO{})

	opt := newSearchFilters()
	repo := newTestRepository(d)
	count, err := repo.Count(context.Background(), opt)

	assert.Equal(t, int64(1), count)
	assert.Nil(t, err)
}
