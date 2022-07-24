package db

import (
	"context"
	"fmt"
	"github.com/go-funcards/authz-service/internal/authz"
	"github.com/go-funcards/mongodb"
	"github.com/go-funcards/slice"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ authz.SubjectStorage = (*subStorage)(nil)

const subCollection = "authz_subjects"

type subStorage struct {
	c   *mongo.Collection
	log zerolog.Logger
}

func NewSubStorage(ctx context.Context, db *mongo.Database, log zerolog.Logger) *subStorage {
	s := &subStorage{
		c:   db.Collection(subCollection),
		log: log.With().Str("storage", "mongodb").Str("collection", subCollection).Logger(),
	}
	s.indexes(ctx)
	return s
}

func (s *subStorage) indexes(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	name, err := s.c.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{"refs.ref_id", 1}},
	})
	if err != nil {
		s.log.Fatal().Err(err).Msg("index not created")
	}

	s.log.Info().Str("index.name", name).Msg("index created")
}

func (s *subStorage) Save(ctx context.Context, model authz.Subject) error {
	var write []mongo.WriteModel
	data, err := mongodb.ToBson(model)
	if err != nil {
		return err
	}

	delete(data, "_id")
	delete(data, "refs")

	if deleteRefs := slice.Map(model.Refs, func(item authz.Ref) string {
		return item.RefID
	}); len(deleteRefs) > 0 {
		s.log.Info().
			Str("sub_id", model.SubID).
			Strs("refs", deleteRefs).
			Msg("delete refs")

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

	s.log.Info().Str("sub_id", model.SubID).Msg("sub save")

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	result, err := s.c.BulkWrite(ctx, write)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("sub save: %s", mongodb.ErrMsgQuery), err)
	}

	s.log.Info().Str("sub_id", model.SubID).Interface("result", result).Msg("sub saved")

	return nil
}

func (s *subStorage) Delete(ctx context.Context, sub string) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	s.log.Info().Str("sub_id", sub).Msg("sub delete")
	result, err := s.c.DeleteOne(ctx, bson.M{"_id": sub})
	if err != nil {
		return fmt.Errorf(mongodb.ErrMsgQuery, err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf(mongodb.ErrMsgQuery, mongo.ErrNoDocuments)
	}
	s.log.Info().Str("sub_id", sub).Msg("sub deleted")

	return nil
}

func (s *subStorage) DeleteByRefID(ctx context.Context, ref string) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	_, err := s.c.UpdateMany(
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
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	result := s.c.FindOne(ctx, bson.M{"_id": sub})

	return mongodb.DecodeOne[authz.Subject](result)
}
