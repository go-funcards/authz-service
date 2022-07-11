package db

import (
	"context"
	"fmt"
	"github.com/go-funcards/authz-service/internal/authz"
	"github.com/go-funcards/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"go.uber.org/zap"
)

var _ authz.RuleStorage = (*ruleStorage)(nil)

type ruleStorage struct {
	c mongodb.Collection[authz.Rule]
}

func NewRuleStorage(ctx context.Context, db *mongo.Database, logger *zap.Logger) (*ruleStorage, error) {
	s := &ruleStorage{c: mongodb.Collection[authz.Rule]{
		Inner: db.Collection("authz_rules"),
		Log:   logger,
	}}
	if err := s.indexes(ctx); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *ruleStorage) indexes(ctx context.Context) error {
	indexes := []string{"type", "v0", "v1", "v2", "v3", "v4", "v5"}
	keysDoc := bsonx.Doc{}

	for _, k := range indexes {
		keysDoc = keysDoc.Append(k, bsonx.Int32(1))
	}

	_, err := s.c.Inner.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    keysDoc,
		Options: options.Index().SetUnique(true),
	})
	return err
}

func (s *ruleStorage) SaveMany(ctx context.Context, models []authz.Rule) error {
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
			SetFilter(bson.M{"_id": model.RuleID}).
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

func (s *ruleStorage) DeleteMany(ctx context.Context, id ...string) error {
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

func (s *ruleStorage) Find(ctx context.Context) ([]authz.Rule, error) {
	return s.c.Find(ctx, bson.M{})
}
