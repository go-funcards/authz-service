package db

import (
	"context"
	"fmt"
	"github.com/go-funcards/authz-service/internal/authz"
	"github.com/go-funcards/mongodb"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var _ authz.RuleStorage = (*ruleStorage)(nil)

const ruleCollection = "authz_rules"

type ruleStorage struct {
	c   *mongo.Collection
	log zerolog.Logger
}

func NewRuleStorage(ctx context.Context, db *mongo.Database, log zerolog.Logger) *ruleStorage {
	s := &ruleStorage{
		c:   db.Collection(ruleCollection),
		log: log.With().Str("storage", "mongodb").Str("collection", ruleCollection).Logger(),
	}
	s.indexes(ctx)
	return s
}

func (s *ruleStorage) indexes(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	indexes := []string{"type", "v0", "v1", "v2", "v3", "v4", "v5"}
	keysDoc := bsonx.Doc{}

	for _, k := range indexes {
		keysDoc = keysDoc.Append(k, bsonx.Int32(1))
	}

	name, err := s.c.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    keysDoc,
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		s.log.Fatal().Err(err).Msg("index not created")
	}

	s.log.Info().Str("index.name", name).Msg("index created")
}

func (s *ruleStorage) SaveMany(ctx context.Context, models []authz.Rule) error {
	var write []mongo.WriteModel
	for _, model := range models {
		data, err := mongodb.ToBson(model)
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

	s.log.Info().Msg("rules save")

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	result, err := s.c.BulkWrite(ctx, write)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("rules save: %s", mongodb.ErrMsgQuery), err)
	}

	s.log.Info().Interface("result", result).Msg("rules saved")

	return nil
}

func (s *ruleStorage) DeleteMany(ctx context.Context, id ...string) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	s.log.Info().Strs("rule_ids", id).Msg("rules delete")
	result, err := s.c.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": id}})
	if err != nil {
		return fmt.Errorf(mongodb.ErrMsgQuery, err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf(mongodb.ErrMsgQuery, mongo.ErrNoDocuments)
	}
	s.log.Info().Strs("rule_ids", id).Msg("rules deleted")

	return nil
}

func (s *ruleStorage) Find(ctx context.Context) ([]authz.Rule, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	cur, err := s.c.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf(mongodb.ErrMsgQuery, err)
	}
	return mongodb.DecodeAll[authz.Rule](ctx, cur)
}
