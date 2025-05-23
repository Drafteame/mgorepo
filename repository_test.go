package mgorepo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {
	t.Parallel()

	d := getTestDriver(t)

	repo := NewRepository[
		testModel,
		testDAO,
		searchFilters,
		SearchOrders,
		SearchOptions[searchFilters, SearchOrders],
		updateFields,
	](
		d,
		collection,
		getFilterBuilders(),
		getUpdateBuilders(),
	)

	assert.NotNil(t, repo.clock, "clock should not be nil")
	assert.NotEmpty(t, repo.db, "db should not be empty")
	assert.NotEmpty(t, repo.collectionName, "collectionName should not be empty")
	assert.NotEmpty(t, repo.filterBuilders, "filterBuilders should not be empty")
	assert.NotEmpty(t, repo.updateBuilders, "updateBuilders should not be empty")

	repo = repo.SetCreatedAtField("created")
	assert.Equal(t, "created", repo.createdAtField, "createdAtField should be equal")
}

func TestRepository_SetDefaultSearchLimit(t *testing.T) {
	t.Parallel()

	d := getTestDriver(t)

	repo := newTestRepository(d)

	assert.Equal(t, DefaultSearchLimit, repo.searchLimit, "defaultSearchLimit should be equal")

	repo = repo.SetDefaultSearchLimit(100)
	assert.Equal(t, 100, repo.searchLimit, "defaultSearchLimit should be equal")

	repo = repo.SetDefaultSearchLimit(0)
	assert.Equal(t, DefaultSearchLimit, repo.searchLimit, "defaultSearchLimit should be equal")
}

func TestRepository_SetLogger(t *testing.T) {
	t.Parallel()

	d := getTestDriver(t)

	repo := newTestRepository(d)

	assert.NotNil(t, repo.Logger(), "logger should not be nil")

	repo = repo.SetLogger(nil)
	assert.Nil(t, repo.Logger(), "logger should be nil")
}

func TestRepository_SetUpdatedAtField(t *testing.T) {
	t.Parallel()

	d := getTestDriver(t)

	repo := newTestRepository(d)

	assert.Equal(t, DefaultUpdatedAtField, repo.updatedAtField, "updatedAtField should be equal")

	repo = repo.SetUpdatedAtField("updated")
	assert.Equal(t, "updated", repo.updatedAtField, "updatedAtField should be equal")

	repo = repo.SetUpdatedAtField("")
	assert.Equal(t, DefaultUpdatedAtField, repo.updatedAtField, "updatedAtField should be equal")
}

func TestRepository_SetCreatedAtField(t *testing.T) {
	t.Parallel()

	d := getTestDriver(t)

	repo := newTestRepository(d)

	assert.Equal(t, DefaultCreatedAtField, repo.createdAtField, "createdAtField should be equal")

	repo = repo.SetCreatedAtField("created")
	assert.Equal(t, "created", repo.createdAtField, "createdAtField should be equal")

	repo = repo.SetCreatedAtField("")
	assert.Equal(t, DefaultCreatedAtField, repo.createdAtField, "createdAtField should be equal")
}

func TestRepository_SetDeletedAtField(t *testing.T) {
	t.Parallel()

	d := getTestDriver(t)

	repo := newTestRepository(d)

	assert.Equal(t, DefaultDeletedAtField, repo.deletedAtField, "deletedAtField should be equal")

	repo = repo.SetDeletedAtField("deleted")
	assert.Equal(t, "deleted", repo.deletedAtField, "deletedAtField should be equal")

	repo = repo.SetDeletedAtField("")
	assert.Equal(t, DefaultDeletedAtField, repo.deletedAtField, "deletedAtField should be equal")
}

func TestRepository_WithTimestamps(t *testing.T) {
	t.Parallel()

	d := getTestDriver(t)

	repo := newTestRepository(d)

	assert.True(t, repo.withTimestamps, "timestamps should be true")

	repo = repo.WithTimestamps(false)
	assert.False(t, repo.withTimestamps, "timestamps should be false")
}

func TestRepository_Db(t *testing.T) {
	t.Parallel()

	d := getTestDriver(t)

	repo := newTestRepository(d)

	assert.NotNil(t, repo.db, "db should not be nil")
	assert.Equal(t, repo.db, repo.Db(), "db should be equal")
}

func TestRepository_Clock(t *testing.T) {
	t.Parallel()

	d := getTestDriver(t)

	repo := newTestRepository(d)

	assert.NotNil(t, repo.clock, "clock should not be nil")
	assert.Equal(t, repo.clock, repo.Clock(), "clock should be equal")

	repo = repo.SetClock(nil)
	assert.Nil(t, repo.clock, "clock should be nil")
}
func TestRepository_Tracer(t *testing.T) {
	t.Parallel()

	d := getTestDriver(t)

	repo := newTestRepository(d)

	assert.NotNil(t, repo.tracer, "tracer should not be nil")
	assert.Equal(t, repo.tracer, repo.Tracer(), "tracer should be equal")

	repo = repo.SetTracer(nil)
	assert.Nil(t, repo.tracer, "tracer should be nil")
}

func TestRepository_CollectionName(t *testing.T) {
	t.Parallel()

	d := getTestDriver(t)

	repo := newTestRepository(d)

	assert.NotEmpty(t, repo.collectionName, "collectionName should not be empty")
	assert.Equal(t, repo.collectionName, repo.CollectionName(), "collectionName should be equal")
}
