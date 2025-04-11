package tracer

import (
	"context"
)

type Tracer struct {
	BeginSubSegmentFn func(ctx context.Context, name string) (context.Context, Segment)
}

func New() Tracer {
	return Tracer{}
}

func (xt Tracer) BeginSubSegment(ctx context.Context, name string) (context.Context, Segment) {
	if xt.BeginSubSegmentFn != nil {
		return xt.BeginSubSegmentFn(ctx, name)
	}

	return ctx, &segment{}
}
