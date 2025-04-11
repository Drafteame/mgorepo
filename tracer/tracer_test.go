package tracer

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockSegment struct {
	mock.Mock
}

func (m *MockSegment) Close(err error) {
	m.Called(err)
}

func (m *MockSegment) AddMetadata(key string, value any) error {
	args := m.Called(key, value)
	return args.Error(0)
}

func (m *MockSegment) AddError(err error) error {
	args := m.Called(err)
	return args.Error(0)
}

func TestNew(t *testing.T) {
	tracer := New()
	assert.NotNil(t, tracer)
}

func TestTracer_BeginSubSegment_Default(t *testing.T) {
	tracer := New()
	ctx := context.Background()

	newCtx, seg := tracer.BeginSubSegment(ctx, "test-segment")

	assert.Equal(t, ctx, newCtx, "Context should be returned unchanged by default")
	assert.NotNil(t, seg, "Default segment should not be nil")
	assert.IsType(t, &segment{}, seg, "Default segment should be of type *segment")

	assert.NotPanics(t, func() { seg.Close(nil) })
	assert.Nil(t, seg.AddMetadata("key", "value"))
	assert.Nil(t, seg.AddError(errors.New("test error")))
}

func TestTracer_BeginSubSegment_WithFn(t *testing.T) {
	mockSeg := new(MockSegment)
	tracer := New()
	ctx := context.Background()
	ctxKey := "key"
	expectedCtx := context.WithValue(ctx, ctxKey, "value")

	tracer.BeginSubSegmentFn = func(fnCtx context.Context, fnName string) (context.Context, Segment) {
		assert.Equal(t, ctx, fnCtx)
		assert.Equal(t, "test-segment-fn", fnName)
		return expectedCtx, mockSeg
	}

	newCtx, seg := tracer.BeginSubSegment(ctx, "test-segment-fn")

	assert.Equal(t, expectedCtx, newCtx, "Context returned by injected function should be returned")
	assert.Equal(t, mockSeg, seg, "Segment returned by injected function should be returned")
}
