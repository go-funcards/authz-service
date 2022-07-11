package casbin

import (
	"context"
	"errors"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/casbin/casbin/v2/persist/file-adapter"
	"github.com/casbin/json-adapter/v2"
	"github.com/go-funcards/authz-service/internal/authz"
	"github.com/google/uuid"
	"os"
	"strings"
)

type Filter interface {
	Sub() string
	Ref() string
}

type adapter struct {
	filtered    bool
	path        string
	ruleStorage authz.RuleStorage
	subStorage  authz.SubjectStorage
}

// NewAdapter is the constructor for Adapter.
func NewAdapter(path string, ruleStorage authz.RuleStorage, subStorage authz.SubjectStorage) persist.Adapter {
	return &adapter{
		filtered:    false,
		path:        path,
		ruleStorage: ruleStorage,
		subStorage:  subStorage,
	}
}

// NewFilteredAdapter is the constructor for FilteredAdapter.
// Casbin will not automatically call LoadPolicy() for a filtered adapter.
func NewFilteredAdapter(path string, ruleStorage authz.RuleStorage, subStorage authz.SubjectStorage) persist.FilteredAdapter {
	a := NewAdapter(path, ruleStorage, subStorage)
	a.(*adapter).filtered = true
	return a.(*adapter)
}

// LoadPolicy loads all policy rules from the storage.
func (a *adapter) LoadPolicy(model model.Model) error {
	return a.LoadFilteredPolicy(model, nil)
}

// LoadFilteredPolicy loads only policy rules that match the filter.
func (a *adapter) LoadFilteredPolicy(model model.Model, filter any) error {
	data, err := a.ruleStorage.Find(context.TODO())
	if err != nil {
		return err
	}

	if len(data) == 0 {
		if err = a.loadDefaultPolicy(model); err != nil {
			return err
		}
		rules := make([]authz.Rule, 0)
		for k, j := range model["p"] {
			for _, v := range j.Policy {
				rules = append(rules, toRule(k, v))
			}
		}
		for k, j := range model["g"] {
			for _, v := range j.Policy {
				rules = append(rules, toRule(k, v))
			}
		}
		if err = a.ruleStorage.SaveMany(context.TODO(), rules); err != nil {
			return err
		}
	} else {
		for _, line := range data {
			if len(line.Type) > 0 {
				loadPolicyLine(line, model)
			}
		}
	}

	switch f := filter.(type) {
	case Filter:
		a.filtered = true
		sub, err := a.subStorage.FindOne(context.TODO(), f.Sub())
		if err != nil {
			return err
		}
		for _, line := range sub.Rules("g", f.Ref()) {
			loadPolicyLine(line, model)
		}
	default:
		a.filtered = false
	}
	return nil
}

func (a *adapter) loadDefaultPolicy(model model.Model) error {
	fa := a.loadFromPath()
	return fa.LoadPolicy(model)
}

func (a *adapter) loadFromPath() persist.Adapter {
	if _, err := os.Stat(a.path); err == nil {
		return fileadapter.NewAdapter(a.path)
	}

	data := []byte(a.path)
	return jsonadapter.NewAdapter(&data)
}

// SavePolicy saves all policy rules to the storage.
func (a *adapter) SavePolicy(model.Model) error {
	return errors.New("not implemented")
}

// AddPolicy adds a policy rule to the storage.
// This is part of the Auto-Save feature.
func (a *adapter) AddPolicy(string, string, []string) error {
	return errors.New("not implemented")
}

// RemovePolicy removes a policy rule from the storage.
// This is part of the Auto-Save feature.
func (a *adapter) RemovePolicy(string, string, []string) error {
	return errors.New("not implemented")
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
// This is part of the Auto-Save feature.
func (a *adapter) RemoveFilteredPolicy(string, string, int, ...string) error {
	return errors.New("not implemented")
}

// IsFiltered returns true if the loaded policy has been filtered.
func (a *adapter) IsFiltered() bool {
	return a.filtered
}

func toRule(t string, data []string) authz.Rule {
	r := authz.Rule{RuleID: uuid.NewString(), Type: t}
	for i, v := range data {
		switch i {
		case 0:
			r.V0 = v
		case 1:
			r.V1 = v
		case 2:
			r.V2 = v
		case 3:
			r.V3 = v
		case 4:
			r.V4 = v
		case 5:
			r.V5 = v
		}
	}
	return r
}

func loadPolicyLine(line authz.Rule, model model.Model) {
	var p = []string{
		line.Type,
		line.V0,
		line.V1,
		line.V2,
		line.V3,
		line.V4,
		line.V5,
	}

	var lineText string
	if line.V5 != "" {
		lineText = strings.Join(p, ", ")
	} else if line.V4 != "" {
		lineText = strings.Join(p[:6], ", ")
	} else if line.V3 != "" {
		lineText = strings.Join(p[:5], ", ")
	} else if line.V2 != "" {
		lineText = strings.Join(p[:4], ", ")
	} else if line.V1 != "" {
		lineText = strings.Join(p[:3], ", ")
	} else if line.V0 != "" {
		lineText = strings.Join(p[:2], ", ")
	}

	persist.LoadPolicyLine(lineText, model)
}
