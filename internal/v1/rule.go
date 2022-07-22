package v1

import (
	"context"
	"github.com/go-funcards/authz-service/internal/authz"
	"github.com/go-funcards/authz-service/proto/v1"
	"github.com/go-funcards/slice"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ruleServer struct {
	v1.UnimplementedRuleServer
	storage authz.RuleStorage
}

func NewRuleServer(storage authz.RuleStorage) *ruleServer {
	return &ruleServer{storage: storage}
}

func (s *ruleServer) SaveRules(ctx context.Context, in *v1.SaveRulesRequest) (*emptypb.Empty, error) {
	if err := s.storage.SaveMany(ctx, authz.SaveRules(in)); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *ruleServer) DeleteRules(ctx context.Context, in *v1.DeleteRulesRequest) (*emptypb.Empty, error) {
	if err := s.storage.DeleteMany(ctx, in.GetRuleIds()...); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *ruleServer) GetRules(ctx context.Context, _ *emptypb.Empty) (*v1.RulesResponse, error) {
	data, err := s.storage.Find(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.RulesResponse{
		Rules: slice.Map(data, func(item authz.Rule) *v1.RulesResponse_Rule {
			return item.ToProto()
		}),
	}, nil
}
