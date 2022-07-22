package db

import (
	"context"
	"fmt"
	"github.com/go-funcards/authz-service/internal/authz"
	"github.com/go-funcards/mongodb"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var _ authz.DefinitionStorage = (*defStorage)(nil)

const (
	timeout       = 5 * time.Second
	defCollection = "authz_definitions"
)

type defStorage struct {
	c   *mongo.Collection
	log logrus.FieldLogger
}

func NewDefStorage(ctx context.Context, db *mongo.Database, log logrus.FieldLogger) *defStorage {
	s := &defStorage{
		c:   db.Collection(defCollection),
		log: log,
	}
	s.indexes(ctx)
	return s
}

func (s *defStorage) indexes(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	name, err := s.c.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{"sec", 1}, {"key", 1}, {"value", 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"collection": defCollection,
			"error":      err,
		}).Fatal("index not created")
	}

	s.log.WithFields(logrus.Fields{
		"collection": defCollection,
		"name":       name,
	}).Info("index created")
}

func (s *defStorage) SaveMany(ctx context.Context, models []authz.Definition) error {
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
			SetFilter(bson.M{"_id": model.DefID}).
			SetUpdate(bson.M{"$set": data}),
		)
	}

	s.log.Info("defs save")

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	result, err := s.c.BulkWrite(ctx, write)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("defs save: %s", mongodb.ErrMsgQuery), err)
	}

	s.log.WithFields(logrus.Fields{"result": result}).Info("defs saved")

	return nil
}

func (s *defStorage) DeleteMany(ctx context.Context, id ...string) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	s.log.WithField("def_ids", id).Debug("defs delete")
	result, err := s.c.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": id}})
	if err != nil {
		return fmt.Errorf(mongodb.ErrMsgQuery, err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf(mongodb.ErrMsgQuery, mongo.ErrNoDocuments)
	}
	s.log.WithField("def_ids", id).Debug("defs deleted")

	return nil
}

func (s *defStorage) Find(ctx context.Context) ([]authz.Definition, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	cur, err := s.c.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf(mongodb.ErrMsgQuery, err)
	}
	return mongodb.DecodeAll[authz.Definition](ctx, cur)
}
