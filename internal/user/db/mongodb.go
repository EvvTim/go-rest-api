package db

import (
	"context"
	"errors"
	"fmt"
	"go-rest-api/internal/apperror"
	"go-rest-api/internal/user"
	"go-rest-api/pkg/logging"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func (d db) Create(ctx context.Context, user *user.User) (string, error) {
	d.logger.Debug("create user")

	if user == nil {
		return "", fmt.Errorf("user object is nil")
	}

	if user.ID == "" {
		objectID := bson.NewObjectID()
		user.ID = objectID.Hex()
	}

	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}

	insertedID, ok := result.InsertedID.(string)
	if !ok {
		return "", fmt.Errorf("unexpected type for InsertedID: %T", result.InsertedID)
	}

	return insertedID, nil
}
func (d db) GetByUUID(ctx context.Context, uuid string) (*user.User, error) {
	d.logger.Debug("get user by uuid")

	filter := bson.M{"_id": uuid}

	var u user.User
	err := d.collection.FindOne(ctx, filter).Decode(&u)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperror.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get user by uuid: %w", err)
	}

	return &u, nil
}
func (d db) GetList(ctx context.Context) ([]*user.User, error) {
	d.logger.Debug("get user list")

	opts := options.Find().SetBatchSize(100)

	cursor, err := d.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to get user list due to err: %w", err)
	}
	defer cursor.Close(ctx)

	var users []*user.User

	for cursor.Next(ctx) {
		var u user.User
		if err = cursor.Decode(&u); err != nil {
			return nil, fmt.Errorf("failed to decode user due to err: %w", err)
		}
		users = append(users, &u)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return users, nil
}

func (d db) Update(ctx context.Context, user *user.User) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("context error: %w", err)
	}

	d.logger.Debug("update user")

	filter := bson.M{"_id": user.ID}

	result, err := d.collection.ReplaceOne(ctx, filter, user)

	if err != nil {
		return fmt.Errorf("failed to update user due to err: %w", err)
	}

	if result.MatchedCount == 0 {
		return apperror.ErrNotFound
	}

	return nil
}

func (d db) Delete(ctx context.Context, uuid string) error {
	d.logger.Debug("delete user")

	filter := bson.M{"_id": uuid}

	fmt.Println("filter", filter)

	result, err := d.collection.DeleteOne(ctx, filter)

	if err != nil {
		return fmt.Errorf("failed to delete user due to err: %w", err)
	}

	if result.DeletedCount == 0 {
		return apperror.ErrNotFound
	}

	return nil
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) user.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}
