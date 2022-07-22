package v1

import (
	"context"
	"encoding/json"
	"github.com/go-funcards/authz-service/internal/casbin"
	"github.com/go-funcards/authz-service/proto/v1"
	"github.com/go-funcards/slice"
)

const guest = "ROLE_GUEST"

type checkerServer struct {
	v1.UnimplementedAuthorizationCheckerServer
	factory casbin.Factory
}

func NewCheckerServer(factory casbin.Factory) *checkerServer {
	return &checkerServer{factory: factory}
}

func (s *checkerServer) IsGranted(_ context.Context, in *v1.IsGrantedRequest) (*v1.Granted, error) {
	params := slice.Map(in.GetParams(), func(param string) any {
		return param
	})

	if len(params) == 0 {
		params = append(params, guest)
	} else if len(in.Params[0]) == 0 {
		params[0] = guest
	}

	var po paramObject
	if len(params) > 1 {
		unmarshal(in.Params[1], &po)
		params[1] = po.toObject()
	}

	e, err := s.factory(newFilter(params[0].(string), po.Ref))
	if err == nil {
		if ok, err := e.Enforce(params...); err == nil && ok {
			return &v1.Granted{Yes: true}, nil
		}
	}

	return &v1.Granted{Yes: false}, nil
}

type object struct {
	Name  string
	Owner string
}

type paramObject struct {
	Name  string `json:"name"`
	Owner string `json:"owner,omitempty"`
	Ref   string `json:"ref,omitempty"`
}

func (po paramObject) toObject() object {
	return object{Name: po.Name, Owner: po.Owner}
}

var _ casbin.Filter = (*filter)(nil)

type filter struct {
	sub string
	ref string
}

func newFilter(sub, ref string) *filter {
	return &filter{sub: sub, ref: ref}
}

func (f filter) Sub() string {
	return f.sub
}

func (f filter) Ref() string {
	return f.ref
}

func unmarshal(param string, po *paramObject) {
	if err := json.Unmarshal([]byte(param), po); err != nil {
		po.Name = param
	}
}
