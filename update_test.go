package mgorepo

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Drafteame/mgorepo/clock"
	"github.com/Drafteame/mgorepo/internal/seed"
	ptesting "github.com/Drafteame/mgorepo/internal/testing"
)

func TestRepository_Update(t *testing.T) {
	t.Run("success update", func(t *testing.T) {
		d := getTestDriver(t)
		db := d.Client().Database(d.DbName())

		oid := primitive.NewObjectID()
		c := clock.NewTest(time.Now()).ForceUTC()

		dao := testDAO{
			ID:         oid,
			Identifier: "test",
		}

		seed.InsertOne(t, db, collection, dao)

		opts := newUpdateFields().WithIdentifier("test2")

		repo := newTestRepository(d).SetClock(c)
		affected, err := repo.Update(context.Background(), oid.Hex(), opts)

		assert.NoError(t, err)
		assert.Equal(t, int64(1), affected)

		updated, err := repo.Get(context.Background(), oid.Hex())
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "test2", updated.Identifier)
		ptesting.AssertDate(t, c.Now(), updated.UpdatedAt)
	})

	t.Run("update error no fields to update", func(t *testing.T) {
		d := getTestDriver(t)
		db := d.Client().Database(d.DbName())

		oid := primitive.NewObjectID()
		c := clock.NewTest(time.Now()).ForceUTC()

		dao := testDAO{
			ID:         oid,
			Identifier: "test",
		}

		seed.InsertOne(t, db, collection, dao)

		opts := newUpdateFields()

		repo := newTestRepository(d).SetClock(c)
		_, err := repo.Update(context.Background(), oid.Hex(), opts)

		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrEmptyUpdate)
	})

	t.Run("update with no timestamps", func(t *testing.T) {
		d := getTestDriver(t)
		db := d.Client().Database(d.DbName())

		oid := primitive.NewObjectID()
		c := clock.NewTest(time.Now()).ForceUTC()

		dao := testDAO{
			ID:         oid,
			Identifier: "test",
		}

		seed.InsertOne(t, db, collection, dao)

		opts := newUpdateFields().WithIdentifier("test2")

		repo := newTestRepository(d).SetClock(c).WithTimestamps(false)
		affected, err := repo.Update(context.Background(), oid.Hex(), opts)

		assert.NoError(t, err)
		assert.Equal(t, int64(1), affected)

		updated, err := repo.Get(context.Background(), oid.Hex())
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "test2", updated.Identifier)
		ptesting.AssertEmptyDate(t, updated.UpdatedAt)
	})
}
