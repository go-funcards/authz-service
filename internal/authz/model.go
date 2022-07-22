package authz

import (
	"github.com/go-funcards/authz-service/proto/v1"
	"github.com/go-funcards/slice"
)

type Definition struct {
	DefID string `json:"def_id" bson:"_id,omitempty"`
	Sec   string `json:"sec" bson:"sec,omitempty"`
	Key   string `json:"key" bson:"key,omitempty"`
	Value string `json:"value" bson:"value,omitempty"`
}

func (d Definition) ToProto() *v1.DefsResponse_Def {
	return &v1.DefsResponse_Def{
		DefId: d.DefID,
		Sec:   d.Sec,
		Key:   d.Key,
		Value: d.Value,
	}
}

func SaveDefs(in *v1.SaveDefsRequest) []Definition {
	return slice.Map(in.GetDefs(), func(item *v1.SaveDefsRequest_Def) Definition {
		return Definition{
			DefID: item.GetDefId(),
			Sec:   item.GetSec(),
			Key:   item.GetKey(),
			Value: item.GetValue(),
		}
	})
}

type Rule struct {
	RuleID string `json:"rule_id" bson:"_id,omitempty"`
	Type   string `json:"type" bson:"type,omitempty"`
	V0     string `json:"v0" bson:"v0,omitempty"`
	V1     string `json:"v1" bson:"v1,omitempty"`
	V2     string `json:"v2" bson:"v2,omitempty"`
	V3     string `json:"v3" bson:"v3,omitempty"`
	V4     string `json:"v4" bson:"v4,omitempty"`
	V5     string `json:"v5" bson:"v5,omitempty"`
}

func (r Rule) ToProto() *v1.RulesResponse_Rule {
	return &v1.RulesResponse_Rule{
		RuleId: r.RuleID,
		Type:   r.Type,
		V0:     r.V0,
		V1:     r.V1,
		V2:     r.V2,
		V3:     r.V3,
		V4:     r.V4,
		V5:     r.V5,
	}
}

func SaveRules(in *v1.SaveRulesRequest) []Rule {
	return slice.Map(in.GetRules(), func(item *v1.SaveRulesRequest_Rule) Rule {
		return Rule{
			RuleID: item.GetRuleId(),
			Type:   item.GetType(),
			V0:     item.GetV0(),
			V1:     item.GetV1(),
			V2:     item.GetV2(),
			V3:     item.GetV3(),
			V4:     item.GetV4(),
			V5:     item.GetV5(),
		}
	})
}

type Ref struct {
	RefID  string   `json:"ref_id" bson:"ref_id,omitempty"`
	Roles  []string `json:"roles" bson:"roles,omitempty"`
	Delete bool     `json:"-" bson:"-"`
}

func (r Ref) ToProto() *v1.SubResponse_Ref {
	return &v1.SubResponse_Ref{
		RefId: r.RefID,
		Roles: r.Roles,
	}
}

type Subject struct {
	SubID string   `json:"sub_id" bson:"_id,omitempty"`
	Roles []string `json:"roles" bson:"roles,omitempty"`
	Refs  []Ref    `json:"refs" bson:"refs,omitempty"`
}

func (s Subject) ToResponse() *v1.SubResponse {
	return &v1.SubResponse{
		SubId: s.SubID,
		Roles: s.Roles,
		Refs: slice.Map(s.Refs, func(r Ref) *v1.SubResponse_Ref {
			return r.ToProto()
		}),
	}
}

func (s Subject) Rules(t string, refID string) []Rule {
	roles := slice.Copy(s.Roles)

	for _, r := range s.Refs {
		if r.RefID == refID {
			roles = append(roles, r.Roles...)
			break
		}
	}

	rules := make([]Rule, 0, len(roles))

	for _, role := range roles {
		rules = append(rules, Rule{
			Type: t,
			V0:   s.SubID,
			V1:   role,
		})
	}

	return rules
}

func SaveSub(in *v1.SaveSubRequest) Subject {
	return Subject{
		SubID: in.GetSubId(),
		Roles: in.GetRoles(),
		Refs: slice.Map(in.GetRefs(), func(item *v1.SaveSubRequest_Ref) Ref {
			return Ref{
				RefID:  item.GetRefId(),
				Roles:  item.GetRoles(),
				Delete: item.GetDelete(),
			}
		}),
	}
}
