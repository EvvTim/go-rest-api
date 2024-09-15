package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewClient(ctx context.Context, host, port, username, password, database, authDB string) (*mongo.Database, error) {
	var uri string
	var client *mongo.Client
	var err error

	if username != "" && password != "" {
		uri = "mongodb://" + username + ":" + password + "@" + host + ":" + port

		credentials := options.Credential{
			Username: username,
			Password: password,
		}

		clientOpts := options.Client().ApplyURI(uri).SetAuth(credentials)

		client, err = mongo.Connect(clientOpts)

	} else {
		uri = "mongodb://" + host + ":" + port

		clientOpts := options.Client().ApplyURI(uri)
		client, err = mongo.Connect(clientOpts)

	}

	fmt.Println(uri)

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("ping failed: %w", err)
	}

	result, err := client.ListDatabaseNames(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Databases:")
	for _, db := range result {
		fmt.Println(db)
	}

	return client.Database(database), nil
}
