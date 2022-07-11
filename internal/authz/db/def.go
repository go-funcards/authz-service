package db

import (
	"context"
	"fmt"
	"github.com/go-funcards/authz-service/internal/authz"
	"github.com/go-funcards/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"time"
)

var _ authz.DefinitionStorage = (*defStorage)(nil)

const timeout = 5 * time.Second

type defStorage struct {
	c mongodb.Collection[authz.Definition]
}

func NewDefStorage(ctx context.Context, db *mongo.Database, logger *zap.Logger) (*defStorage, error) {
	s := &defStorage{c: mongodb.Collection[authz.Definition]{
		Inner: db.Collection("authz_definitions"),
		Log:   logger,
	}}
	if err := s.indexes(ctx); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *defStorage) indexes(ctx context.Context) error {
	_, err := s.c.Inner.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{"sec", 1}, {"key", 1}, {"value", 1}},
		Options: options.Index().SetUnique(true),
	})
	return err
}

func (s *defStorage) SaveMany(ctx context.Context, models []authz.Definition) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var write []mongo.WriteModel

	for _, model := range models {
		data, err := s.c.ToM(model)
		if err != nil {
			return err
		}

		delete(data, "_id")

		write = append(write, mongo.
			NewUpdateOneModel().
			SetUpsert(true).
			SetFilter(bson.M{"_id": model.DefID}).
			SetUpdate(bson.M{"$set": data}),
		)
	}

	s.c.Log.Debug("bulk update")

	result, err := s.c.Inner.BulkWrite(ctx, write, options.BulkWrite())
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("bulk update: %s", mongodb.ErrMsgQuery), err)
	}

	s.c.Log.Info("documents updated", zap.Any("result", result))

	return nil
}

func (s *defStorage) DeleteMany(ctx context.Context, id ...string) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	s.c.Log.Debug("documents delete")
	result, err := s.c.Inner.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": id}})
	if err != nil {
		return fmt.Errorf(mongodb.ErrMsgQuery, err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf(mongodb.ErrMsgQuery, mongo.ErrNoDocuments)
	}
	s.c.Log.Debug("documents deleted", zap.Int64("deleted", result.DeletedCount))

	return nil
}

func (s *defStorage) Find(ctx context.Context) ([]authz.Definition, error) {
	return s.c.Find(ctx, bson.M{})
}
