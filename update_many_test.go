package mgorepo

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Drafteame/mgorepo/clock"
	"github.com/Drafteame/mgorepo/internal/seed"
	ptesting "github.com/Drafteame/mgorepo/internal/testing"
)

func randomNumber() int {
	return rand.Intn(100)
}

func TestRepository_UpdateMany(t *testing.T) {
	t.Run("success update many", func(t *testing.T) {
		d := getTestDriver(t)
		db := d.Client().Database(d.DbName())

		c := clock.NewTest(time.Now()).ForceUTC()

		daos := make([]any, 0, 100)

		for i := 0; i < 100; i++ {
			oid := primitive.NewObjectID()

			dao := testDAO{
				ID:         oid,
				Identifier: "test",
				Sortable:   randomNumber(),
			}
			daos = append(daos, dao)
		}

		seed.InsertMany(t, db, collection, daos...)

		filters := newSearchFilters().WithSortableGreaterThan(5)
		data := newUpdateFields().WithIdentifier("test2")

		repo := newTestRepository(d).SetClock(c)

		total, err := repo.Count(context.Background(), filters)
		if err != nil {
			t.Fatal(err)
		}

		updated, errUpdate := repo.UpdateMany(context.Background(), filters, data)

		assert.NoError(t, errUpdate)
		assert.Equal(t, total, updated)

		allDocs, errFind := repo.Search(context.Background(), NewSearchOptions(newSearchFilters(), NewSearchOrders()).WithLimit(100))
		if errFind != nil {
			t.Fatal(errFind)
		}

		var totalTest2 int64

		for _, doc := range allDocs {
			if doc.Identifier == "test2" {
				ptesting.AssertDate(t, c.Now(), doc.UpdatedAt)
				totalTest2++
			}
		}

		assert.Equal(t, total, totalTest2)
	})

	t.Run("update many error no fields to update", func(t *testing.T) {
		d := getTestDriver(t)
		db := d.Client().Database(d.DbName())

		c := clock.NewTest(time.Now()).ForceUTC()

		daos := make([]any, 0, 100)

		for i := 0; i < 100; i++ {
			oid := primitive.NewObjectID()

			dao := testDAO{
				ID:         oid,
				Identifier: "test",
				Sortable:   randomNumber(),
			}
			daos = append(daos, dao)
		}

		seed.InsertMany(t, db, collection, daos...)

		filters := newSearchFilters().WithSortableGreaterThan(5)
		data := newUpdateFields()

		repo := newTestRepository(d).SetClock(c)
		updated, errUpdate := repo.UpdateMany(context.Background(), filters, data)

		assert.Error(t, errUpdate)
		assert.ErrorIs(t, errUpdate, ErrEmptyUpdate)
		assert.Equal(t, int64(0), updated)

		allDocs, errFind := repo.Search(context.Background(), NewSearchOptions(newSearchFilters(), NewSearchOrders()).WithLimit(100))
		if errFind != nil {
			t.Fatal(errFind)
		}

		for _, doc := range allDocs {
			ptesting.AssertEmptyDate(t, doc.UpdatedAt)
		}
	})

	t.Run("update many error no filters", func(t *testing.T) {
		d := getTestDriver(t)
		db := d.Client().Database(d.DbName())

		c := clock.NewTest(time.Now()).ForceUTC()

		daos := make([]any, 0, 100)

		for i := 0; i < 100; i++ {
			oid := primitive.NewObjectID()

			dao := testDAO{
				ID:         oid,
				Identifier: "test",
				Sortable:   randomNumber(),
			}
			daos = append(daos, dao)
		}

		seed.InsertMany(t, db, collection, daos...)

		filters := newSearchFilters()
		data := newUpdateFields().WithIdentifier("test2")

		repo := newTestRepository(d).SetClock(c)

		updated, errUpdate := repo.UpdateMany(context.Background(), filters, data)

		assert.Error(t, errUpdate)
		assert.ErrorIs(t, errUpdate, ErrEmptyFilters)
		assert.Equal(t, int64(0), updated)

		allDocs, errFind := repo.Search(context.Background(), NewSearchOptions(newSearchFilters(), NewSearchOrders()).WithLimit(100))
		if errFind != nil {
			t.Fatal(errFind)
		}

		for _, doc := range allDocs {
			ptesting.AssertEmptyDate(t, doc.UpdatedAt)
		}
	})

	t.Run("update many with no timestamps", func(t *testing.T) {
		d := getTestDriver(t)
		db := d.Client().Database(d.DbName())

		daos := make([]any, 0, 100)

		for i := 0; i < 100; i++ {
			oid := primitive.NewObjectID()
			dao := testDAO{
				ID:         oid,
				Identifier: "test",
				Sortable:   randomNumber(),
			}
			daos = append(daos, dao)
		}

		seed.InsertMany(t, db, collection, daos...)

		filters := newSearchFilters().WithSortableGreaterThan(5)
		data := newUpdateFields().WithIdentifier("test2")

		repo := newTestRepository(d).WithTimestamps(false)

		expTotal, err := repo.Count(context.Background(), filters)
		if err != nil {
			t.Fatal(err)
		}

		updated, errUpdate := repo.UpdateMany(context.Background(), filters, data)

		assert.NoError(t, errUpdate)
		assert.Equal(t, expTotal, updated)

		allDocs, errFind := repo.Search(context.Background(), NewSearchOptions(newSearchFilters(), NewSearchOrders()).WithLimit(100))
		if errFind != nil {
			t.Fatal(errFind)
		}

		var total int64

		for _, doc := range allDocs {
			if doc.Identifier == "test2" {
				total++
			}

			ptesting.AssertEmptyDate(t, doc.UpdatedAt)
		}

		assert.Equal(t, expTotal, total)
	})
}
