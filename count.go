package mgorepo

import (
	"context"
	"fmt"
)

// Count returns the number of documents in the collection that match the given search options.
func (r Repository[M, D, SF, SORD, SO, UF]) Count(ctx context.Context, opts SF) (int64, error) {
	filters, err := r.BuildSearchFilters(opts)
	if err != nil {
		return 0, err
	}

	ctx, seg := r.tracer.BeginSubSegment(ctx, fmt.Sprintf("MongoDB.Count.%s", r.collectionName))
	defer seg.Close(nil)

	_ = seg.AddMetadata("filters", filters)

	r.logDebugf(actionCount, "filters: %+v", filters)

	return r.Collection().CountDocuments(ctx, filters)
}
