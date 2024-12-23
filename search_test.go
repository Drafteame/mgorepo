package mgorepo

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Drafteame/mgorepo/internal/seed"
	ptesting "github.com/Drafteame/mgorepo/internal/testing"
)

func TestRepository_Search(t *testing.T) {
	t.Run("success search", func(t *testing.T) {
		d := getTestDriver(t)
		db := d.Client().Database(d.DbName())

		now := time.Now().UTC()

		dao := testDAO{
			Identifier: "identifier",
			CreatedAt:  primitive.NewDateTimeFromTime(now),
			UpdatedAt:  primitive.NewDateTimeFromTime(now),
		}

		seed.InsertMany(t, db, collection, dao, dao)

		repo := newTestRepository(d)

		opt := NewSearchOptions(newSearchFilters().WithIdentifier("identifier"), NewSearchOrders())
		models, err := repo.Search(context.Background(), opt)

		assert.NoError(t, err)
		assert.Len(t, models, 2)
	})

	t.Run("empty search", func(t *testing.T) {
		d := getTestDriver(t)
		repo := newTestRepository(d)

		opt := NewSearchOptions(newSearchFilters(), NewSearchOrders())
		models, err := repo.Search(context.Background(), opt)

		assert.NoError(t, err)
		assert.Len(t, models, 0)
	})

	t.Run("success search with limit", func(t *testing.T) {
		d := getTestDriver(t)
		db := d.Client().Database(d.DbName())
		now := time.Now().UTC()

		dao := testDAO{
			Identifier: "identifier",
			CreatedAt:  primitive.NewDateTimeFromTime(now),
			UpdatedAt:  primitive.NewDateTimeFromTime(now),
		}

		seed.InsertMany(t, db, collection, dao, dao)

		repo := newTestRepository(d)

		opt := NewSearchOptions(newSearchFilters(), NewSearchOrders()).WithLimit(1)
		models, err := repo.Search(context.Background(), opt)

		assert.NoError(t, err)
		assert.Len(t, models, 1)
	})

	t.Run("success search with offset", func(t *testing.T) {
		d := getTestDriver(t)
		db := d.Client().Database(d.DbName())
		now := time.Now().UTC()

		dao := testDAO{
			Identifier: "identifier",
			Sortable:   1,
			CreatedAt:  primitive.NewDateTimeFromTime(now),
			UpdatedAt:  primitive.NewDateTimeFromTime(now),
		}

		dao2 := testDAO{
			Identifier: "identifier",
			Sortable:   2,
			CreatedAt:  primitive.NewDateTimeFromTime(now),
			UpdatedAt:  primitive.NewDateTimeFromTime(now),
		}

		seed.InsertMany(t, db, collection, dao, dao2)

		repo := newTestRepository(d)

		opt := NewSearchOptions(newSearchFilters(), NewSearchOrders()).WithSkip(1)
		models, err := repo.Search(context.Background(), opt)

		assert.NoError(t, err)
		assert.Len(t, models, 1)
		assert.Equal(t, dao2.Sortable, models[0].Sortable)
	})

	t.Run("success search with sort", func(t *testing.T) {
		d := getTestDriver(t)
		db := d.Client().Database(d.DbName())

		now := time.Now().UTC()

		dao := testDAO{
			Identifier: "identifier",
			Sortable:   1,
			CreatedAt:  primitive.NewDateTimeFromTime(now),
			UpdatedAt:  primitive.NewDateTimeFromTime(now),
		}

		dao2 := testDAO{
			Identifier: "identifier",
			Sortable:   2,
			CreatedAt:  primitive.NewDateTimeFromTime(now),
			UpdatedAt:  primitive.NewDateTimeFromTime(now),
		}

		seed.InsertMany(t, db, collection, dao, dao2)

		repo := newTestRepository(d)

		opt := NewSearchOptions(newSearchFilters(), NewSearchOrders().Add(sortableField, 1))
		models, err := repo.Search(context.Background(), opt)

		assert.NoError(t, err)
		assert.Len(t, models, 2)
		assert.Equal(t, dao.Sortable, models[0].Sortable)
		assert.Equal(t, dao2.Sortable, models[1].Sortable)
	})

	t.Run("success search with sort desc", func(t *testing.T) {
		d := getTestDriver(t)
		db := d.Client().Database(d.DbName())
		now := time.Now().UTC()

		dao := testDAO{
			Identifier: "identifier",
			Sortable:   1,
			CreatedAt:  primitive.NewDateTimeFromTime(now),
			UpdatedAt:  primitive.NewDateTimeFromTime(now),
		}

		dao2 := testDAO{
			Identifier: "identifier",
			Sortable:   2,
			CreatedAt:  primitive.NewDateTimeFromTime(now),
			UpdatedAt:  primitive.NewDateTimeFromTime(now),
		}

		seed.InsertMany(t, db, collection, dao, dao2)

		repo := newTestRepository(d)

		opt := NewSearchOptions(newSearchFilters(), NewSearchOrders().Add(sortableField, -1))
		models, err := repo.Search(context.Background(), opt)

		assert.NoError(t, err)
		assert.Len(t, models, 2)
		assert.Equal(t, dao2.Sortable, models[0].Sortable)
		assert.Equal(t, dao.Sortable, models[1].Sortable)
	})

	t.Run("success omitting deleted items", func(t *testing.T) {
		d := getTestDriver(t)
		db := d.Client().Database(d.DbName())
		now := time.Now().UTC()
		pnow := primitive.NewDateTimeFromTime(now)

		dao := testDAO{
			Identifier: "identifier",
			Sortable:   1,
			CreatedAt:  primitive.NewDateTimeFromTime(now),
			UpdatedAt:  primitive.NewDateTimeFromTime(now),
		}

		dao2 := testDAO{
			Identifier: "identifier",
			Sortable:   2,
			CreatedAt:  primitive.NewDateTimeFromTime(now),
			UpdatedAt:  primitive.NewDateTimeFromTime(now),
			DeletedAt:  pnow,
		}

		seed.InsertMany(t, db, collection, dao, dao2)

		repo := newTestRepository(d)

		opt := NewSearchOptions(newSearchFilters(), NewSearchOrders())
		models, err := repo.Search(context.Background(), opt)

		assert.NoError(t, err)
		assert.Len(t, models, 1)
		assert.Equal(t, dao.Sortable, models[0].Sortable)
	})

	t.Run("success search with projection", func(t *testing.T) {
		d := getTestDriver(t)
		db := d.Client().Database(d.DbName())
		now := time.Now().UTC()
		dao := testDAO{
			Identifier: "identifier",
			Sortable:   1,
			CreatedAt:  primitive.NewDateTimeFromTime(now),
			UpdatedAt:  primitive.NewDateTimeFromTime(now),
		}

		dao2 := testDAO{
			Identifier: "identifier",
			Sortable:   2,
			CreatedAt:  primitive.NewDateTimeFromTime(now),
			UpdatedAt:  primitive.NewDateTimeFromTime(now),
		}

		seed.InsertMany(t, db, collection, dao, dao2)

		repo := newTestRepository(d)

		opt := NewSearchOptions(newSearchFilters(), NewSearchOrders().Add(sortableField, 1)).
			WithProject(sortableField, FieldAdd)

		models, err := repo.Search(context.Background(), opt)

		assert.NoError(t, err)
		assert.Len(t, models, 2)
		assert.NotEqual(t, "", models[0].ID)
		assert.NotEqual(t, "", models[1].ID)
		assert.Equal(t, dao.Sortable, models[0].Sortable)
		assert.Equal(t, dao2.Sortable, models[1].Sortable)
		assert.Equal(t, "", models[0].Identifier)
		assert.Equal(t, "", models[1].Identifier)
		ptesting.AssertEmptyDate(t, models[0].CreatedAt)
		ptesting.AssertEmptyDate(t, models[1].CreatedAt)
		ptesting.AssertEmptyDate(t, models[0].UpdatedAt)
		ptesting.AssertEmptyDate(t, models[1].UpdatedAt)
		ptesting.AssertEmptyDate(t, models[0].DeletedAt)
		ptesting.AssertEmptyDate(t, models[1].DeletedAt)
	})

	t.Run("success search with inverse projection", func(t *testing.T) {
		d := getTestDriver(t)
		db := d.Client().Database(d.DbName())
		now := time.Now().UTC()
		dao := testDAO{
			Identifier: "identifier",
			Sortable:   1,
			CreatedAt:  primitive.NewDateTimeFromTime(now),
			UpdatedAt:  primitive.NewDateTimeFromTime(now),
		}

		dao2 := testDAO{
			Identifier: "identifier",
			Sortable:   2,
			CreatedAt:  primitive.NewDateTimeFromTime(now),
			UpdatedAt:  primitive.NewDateTimeFromTime(now),
		}

		seed.InsertMany(t, db, collection, dao, dao2)

		repo := newTestRepository(d)

		opt := NewSearchOptions(newSearchFilters(), NewSearchOrders().Add(sortableField, OrderAsc)).
			WithProject(sortableField, FieldRemove)

		models, err := repo.Search(context.Background(), opt)

		assert.NoError(t, err)
		assert.Len(t, models, 2)
		assert.NotEqual(t, "", models[0].ID)
		assert.NotEqual(t, "", models[1].ID)
		assert.Equal(t, 0, models[0].Sortable)
		assert.Equal(t, 0, models[1].Sortable)
		assert.Equal(t, "identifier", models[0].Identifier)
		assert.Equal(t, "identifier", models[1].Identifier)
		ptesting.AssertDate(t, now, models[0].CreatedAt)
		ptesting.AssertDate(t, now, models[1].CreatedAt)
		ptesting.AssertDate(t, now, models[0].UpdatedAt)
		ptesting.AssertDate(t, now, models[1].UpdatedAt)
		ptesting.AssertEmptyDate(t, models[0].DeletedAt)
		ptesting.AssertEmptyDate(t, models[1].DeletedAt)
	})

	t.Run("success search with multi field projection", func(t *testing.T) {
		d := getTestDriver(t)
		db := d.Client().Database(d.DbName())
		oid := primitive.NewObjectID()
		oid2 := primitive.NewObjectID()
		now := time.Now().UTC()

		dao := testDAO{
			ID:         oid,
			Identifier: "identifier",
			Sortable:   1,
			CreatedAt:  primitive.NewDateTimeFromTime(now),
			UpdatedAt:  primitive.NewDateTimeFromTime(now),
		}

		dao2 := testDAO{
			ID:         oid2,
			Identifier: "identifier",
			Sortable:   2,
			CreatedAt:  primitive.NewDateTimeFromTime(now),
			UpdatedAt:  primitive.NewDateTimeFromTime(now),
		}

		seed.InsertMany(t, db, collection, dao, dao2)

		repo := newTestRepository(d)

		opt := NewSearchOptions(newSearchFilters(), NewSearchOrders().Add(sortableField, 1)).
			WithProjectFields(map[string]int{
				idField:         FieldRemove,
				identifierField: FieldAdd,
				sortableField:   FieldAdd,
			})

		models, err := repo.Search(context.Background(), opt)

		assert.NoError(t, err)
		assert.Len(t, models, 2)

		assert.Equal(t, "", models[0].ID)
		assert.Equal(t, 1, models[0].Sortable)
		assert.Equal(t, "identifier", models[0].Identifier)
		ptesting.AssertEmptyDate(t, models[0].CreatedAt)
		ptesting.AssertEmptyDate(t, models[0].UpdatedAt)
		ptesting.AssertEmptyDate(t, models[0].DeletedAt)

		assert.Equal(t, "", models[1].ID)
		assert.Equal(t, 2, models[1].Sortable)
		assert.Equal(t, "identifier", models[1].Identifier)
		ptesting.AssertEmptyDate(t, models[1].CreatedAt)
		ptesting.AssertEmptyDate(t, models[1].UpdatedAt)
		ptesting.AssertEmptyDate(t, models[1].DeletedAt)
	})
}
