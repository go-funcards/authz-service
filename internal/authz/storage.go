package authz

import "context"

type DefinitionStorage interface {
	SaveMany(ctx context.Context, models []Definition) error
	DeleteMany(ctx context.Context, id ...string) error
	Find(ctx context.Context) ([]Definition, error)
}

type RuleStorage interface {
	SaveMany(ctx context.Context, models []Rule) error
	DeleteMany(ctx context.Context, id ...string) error
	Find(ctx context.Context) ([]Rule, error)
}

type SubjectStorage interface {
	Save(ctx context.Context, model Subject) error
	Delete(ctx context.Context, sub string) error
	DeleteByRefID(ctx context.Context, ref string) error
	FindOne(ctx context.Context, sub string) (Subject, error)
}
