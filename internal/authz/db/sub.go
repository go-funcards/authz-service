package db

import (
	"context"
	"fmt"
	"github.com/go-funcards/authz-service/internal/authz"
	"github.com/go-funcards/mongodb"
	"github.com/go-funcards/slice"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var _ authz.SubjectStorage = (*subStorage)(nil)

type subStorage struct {
	c mongodb.Collection[authz.Subject]
}

func NewSubStorage(ctx context.Context, db *mongo.Database, logger *zap.Logger) (*subStorage, error) {
	s := &subStorage{c: mongodb.Collection[authz.Subject]{
		Inner: db.Collection("authz_subjects"),
		Log:   logger,
	}}
	if err := s.indexes(ctx); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *subStorage) indexes(ctx context.Context) error {
	_, err := s.c.Inner.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{"refs.ref_id", 1}},
		Options: options.Index(),
	})
	return err
}

func (s *subStorage) Save(ctx context.Context, model authz.Subject) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var write []mongo.WriteModel
	data, err := s.c.ToM(model)
	if err != nil {
		return err
	}

	delete(data, "_id")
	delete(data, "refs")

	if deleteRefs := slice.Map(model.Refs, func(item authz.Ref) string {
		return item.RefID
	}); len(deleteRefs) > 0 {
		s.c.Log.Info("delete refs", zap.String("sub_id", model.SubID), zap.Strings("refs", deleteRefs))

		write = append(write, mongo.
			NewUpdateOneModel().
			SetFilter(bson.M{"_id": model.SubID}).
			SetUpdate(bson.M{
				"$pull": bson.M{
					"refs": bson.M{
						"ref_id": bson.M{
							"$in": deleteRefs,
						},
					},
				},
			}),
		)
	}

	addRefs := slice.Filter(model.Refs, func(item authz.Ref) bool {
		return !item.Delete
	})

	write = append(write, mongo.
		NewUpdateOneModel().
		SetUpsert(true).
		SetFilter(bson.M{"_id": model.SubID}).
		SetUpdate(bson.M{
			"$set": data,
			"$addToSet": bson.M{
				"refs": bson.M{"$each": addRefs},
			},
		}),
	)

	s.c.Log.Debug("bulk update")

	result, err := s.c.Inner.BulkWrite(ctx, write, options.BulkWrite())
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("bulk update: %s", mongodb.ErrMsgQuery), err)
	}

	s.c.Log.Info("document updated", zap.String("sub_id", model.SubID), zap.Any("result", result))

	return nil
}

func (s *subStorage) Delete(ctx context.Context, sub string) error {
	return s.c.DeleteOne(ctx, bson.M{"_id": sub})
}

func (s *subStorage) DeleteByRefID(ctx context.Context, ref string) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	_, err := s.c.Inner.UpdateMany(
		ctx,
		bson.M{"refs.ref_id": ref},
		bson.M{
			"$pull": bson.M{
				"refs": bson.M{
					"ref_id": ref,
				},
			},
		},
	)
	return err
}

func (s *subStorage) FindOne(ctx context.Context, sub string) (authz.Subject, error) {
	return s.c.FindOne(ctx, bson.M{"_id": sub})
}
