package tracer

type Segment interface {
	Close(error)
	AddMetadata(key string, value any) error
	AddError(error) error
}

type segment struct{}

var _ Segment = (*segment)(nil)

func (s segment) Close(_ error) {}

func (s segment) AddMetadata(_ string, _ any) error {
	return nil
}

func (s segment) AddError(_ error) error {
	return nil
}
