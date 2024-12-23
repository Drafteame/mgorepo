package mgorepo

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Drafteame/mgorepo/clock"
	"github.com/Drafteame/mgorepo/internal/seed"
)

func TestRepository_Get(t *testing.T) {
	t.Parallel()

	t.Run("get error not found", func(t *testing.T) {
		t.Parallel()

		d := getTestDriver(t)
		db := d.Client().Database(d.DbName())

		oid := primitive.NewObjectID()
		c := clock.NewTest(time.Now()).ForceUTC()

		dao := testDAO{
			ID:         oid,
			Identifier: "identifier",
			CreatedAt:  primitive.NewDateTimeFromTime(c.Now()),
			UpdatedAt:  primitive.NewDateTimeFromTime(c.Now()),
		}

		seed.InsertOne(t, db, collection, dao)

		repo := newTestRepository(d).SetClock(c)

		model, err := repo.Get(context.Background(), primitive.NewObjectID().Hex())

		assert.Empty(t, model)
		assert.Error(t, err)
		assert.True(t, errors.Is(err, ErrNotFound))
	})

	t.Run("get error not found on deleted doc", func(t *testing.T) {
		t.Parallel()

		d := getTestDriver(t)
		db := d.Client().Database(d.DbName())

		oid := primitive.NewObjectID()
		c := clock.NewTest(time.Now()).ForceUTC()

		dao := testDAO{
			ID:         oid,
			Identifier: "identifier",
			CreatedAt:  primitive.NewDateTimeFromTime(c.Now()),
			UpdatedAt:  primitive.NewDateTimeFromTime(c.Now()),
			DeletedAt:  primitive.NewDateTimeFromTime(c.Now()),
		}

		seed.InsertOne(t, db, collection, dao)

		repo := newTestRepository(d).SetClock(c)

		model, err := repo.Get(context.Background(), oid.Hex())

		assert.Empty(t, model)
		assert.Error(t, err)
		assert.True(t, errors.Is(err, ErrNotFound))
	})

	t.Run("get success", func(t *testing.T) {
		t.Parallel()

		d := getTestDriver(t)
		db := d.Client().Database(d.DbName())

		oid := primitive.NewObjectID()
		c := clock.NewTest(time.Now()).ForceUTC()

		dao := testDAO{
			ID:         oid,
			Identifier: "identifier",
			CreatedAt:  primitive.NewDateTimeFromTime(c.Now()),
			UpdatedAt:  primitive.NewDateTimeFromTime(c.Now()),
		}

		seed.InsertOne(t, db, collection, dao)

		repo := newTestRepository(d).SetClock(c)
		model, err := repo.Get(context.Background(), oid.Hex())

		assert.NotNil(t, model)
		assert.NoError(t, err)
		assert.IsType(t, testModel{}, model)
	})
}
