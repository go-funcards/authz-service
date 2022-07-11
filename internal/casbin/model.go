package casbin

import (
	"context"
	"github.com/casbin/casbin/v2/model"
	"github.com/go-funcards/authz-service/internal/authz"
	"github.com/google/uuid"
	"os"
)

type factory struct {
	path    string
	storage authz.DefinitionStorage
}

func (f factory) fromFile() (model.Model, error) {
	if _, err := os.Stat(f.path); err == nil {
		return model.NewModelFromFile(f.path)
	}
	return model.NewModelFromString(f.path)
}

func (f factory) fromStorage(ctx context.Context) (m model.Model, err error) {
	data, err := f.storage.Find(ctx)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		m, err = f.fromFile()
		if err != nil {
			return
		}

		for k, v := range m {
			if k == "logger" {
				continue
			}
			for i, j := range v {
				data = append(data, authz.Definition{
					DefID: uuid.NewString(),
					Sec:   k,
					Key:   i,
					Value: j.Value,
				})
			}
		}

		err = f.storage.SaveMany(ctx, data)
		return
	}

	m = model.NewModel()
	for _, item := range data {
		if len(item.Sec) > 0 && len(item.Key) > 0 && len(item.Value) > 0 {
			m.AddDef(item.Sec, item.Key, item.Value)
		}
	}
	return
}

func NewModel(ctx context.Context, path string, storage authz.DefinitionStorage) (model.Model, error) {
	return factory{path: path, storage: storage}.fromStorage(ctx)
}
