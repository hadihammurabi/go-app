package index

import (
	"context"
)

type Index struct {
	*UnimplementedIndexServer
}

func New() *Index {
	return &Index{}
}

func (a *Index) Index(ctx context.Context, none *None) (*IndexIndexResp, error) {
	return &IndexIndexResp{
		Message: "welcome to the app via GRPC",
	}, nil
}
