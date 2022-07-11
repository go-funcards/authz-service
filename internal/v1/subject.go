package v1

import (
	"context"
	"github.com/go-funcards/authz-service/internal/authz"
	"github.com/go-funcards/authz-service/proto/v1"
	"github.com/go-funcards/slice"
	"google.golang.org/protobuf/types/known/emptypb"
)

type subService struct {
	v1.UnimplementedSubjectServer
	storage authz.SubjectStorage
}

func NewSubService(storage authz.SubjectStorage) *subService {
	return &subService{storage: storage}
}

func (s *subService) SaveSub(ctx context.Context, in *v1.SaveSubRequest) (*emptypb.Empty, error) {
	if err := s.storage.Save(ctx, authz.SaveSub(in)); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *subService) DeleteSub(ctx context.Context, in *v1.DeleteSubRequest) (*emptypb.Empty, error) {
	if err := s.storage.Delete(ctx, in.GetSubId()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *subService) DeleteRef(ctx context.Context, in *v1.DeleteRefRequest) (*emptypb.Empty, error) {
	if err := s.storage.DeleteByRefID(ctx, in.GetRefId()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *subService) GetSub(ctx context.Context, in *v1.SubRequest) (*v1.SubResponse, error) {
	item, err := s.storage.FindOne(ctx, in.GetSubId())
	if err != nil {
		return nil, err
	}
	if len(in.GetRefId()) > 0 {
		item.Refs = slice.Filter(item.Refs, func(ref authz.Ref) bool {
			return ref.RefID == in.GetRefId()
		})
	}
	return item.ToResponse(), nil
}
