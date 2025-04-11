package mgorepo

import (
	"context"
	"time"

	"github.com/Drafteame/mgorepo/tracer"
	"go.mongodb.org/mongo-driver/mongo"
)

type Driver interface {
	Client() *mongo.Client
	DbName() string
}

type Clock interface {
	Now() time.Time
}

type Logger interface {
	Debugf(action, message string, args ...any)
	Errorf(err error, action, message string, args ...any)
}

type Model any

type Dao any

type DaoFiller[M Model] interface {
	FromModel(M) error
	ToModel() M
}

type SearchFilters any

type UpdateFields any

type SearchOrderer interface {
	ToMap() map[string]int
}

type SearchOptioner[SF SearchFilters, O SearchOrderer] interface {
	Filters() SF
	Orders() O
	Limit() int64
	Skip() int64
	Projection() map[string]int
}

type Tracer interface {
	BeginSubSegment(ctx context.Context, name string) (context.Context, tracer.Segment)
}
