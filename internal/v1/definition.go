package v1

import (
	"context"
	"github.com/go-funcards/authz-service/internal/authz"
	"github.com/go-funcards/authz-service/proto/v1"
	"github.com/go-funcards/slice"
	"google.golang.org/protobuf/types/known/emptypb"
)

type defServer struct {
	v1.UnimplementedDefinitionServer
	storage authz.DefinitionStorage
}

func NewDefServer(storage authz.DefinitionStorage) *defServer {
	return &defServer{storage: storage}
}

func (s *defServer) SaveDefs(ctx context.Context, in *v1.SaveDefsRequest) (*emptypb.Empty, error) {
	if err := s.storage.SaveMany(ctx, authz.SaveDefs(in)); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *defServer) DeleteDefs(ctx context.Context, in *v1.DeleteDefsRequest) (*emptypb.Empty, error) {
	if err := s.storage.DeleteMany(ctx, in.GetDefIds()...); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *defServer) GetDefs(ctx context.Context, _ *emptypb.Empty) (*v1.DefsResponse, error) {
	data, err := s.storage.Find(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.DefsResponse{
		Defs: slice.Map(data, func(item authz.Definition) *v1.DefsResponse_Def {
			return item.ToProto()
		}),
	}, nil
}
