package casbin

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/go-funcards/authz-service/internal/authz"
)

type Config struct {
	ModelPath  string `yaml:"model_path" env:"MODEL_PATH" env-default:"model.conf"`
	PolicyPath string `yaml:"policy_path" env:"POLICY_PATH" env-default:"policy.csv"`
}

type Factory func(filter Filter) (*casbin.Enforcer, error)

func (cfg Config) EnforcerFactory(def authz.DefinitionStorage, rule authz.RuleStorage, sub authz.SubjectStorage) Factory {
	modelPath := cfg.ModelPath
	policyPath := cfg.PolicyPath

	return func(filter Filter) (*casbin.Enforcer, error) {
		model, err := NewModel(context.TODO(), modelPath, def)
		if err != nil {
			return nil, err
		}

		e, err := casbin.NewEnforcer(model, NewFilteredAdapter(policyPath, rule, sub))
		if err != nil {
			return nil, err
		}

		e.EnableAutoSave(false)

		if err = e.LoadFilteredPolicy(filter); err != nil {
			return nil, err
		}

		return e, nil
	}
}
