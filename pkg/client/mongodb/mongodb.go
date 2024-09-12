package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewClient(ctx context.Context, host, port, username, password, database, authDB string) (*mongo.Database, error) {
	var uri string

	if username != "" && password != "" {
		uri = "mongodb://" + username + ":" + password + "@" + host + ":" + port + "/?authSource=" + authDB
	} else {
		uri = "mongodb://" + host + ":" + port
	}

	fmt.Println(uri)

	//credentials := options.Credential{
	//	AuthSource: authDB,
	//	Username:   username,
	//	Password:   password,
	//}

	client, err := mongo.Connect(options.Client().ApplyURI(uri))

	//defer func() {
	//	if err = client.Disconnect(ctx); err != nil {
	//		panic(err)
	//	}
	//}()

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("ping failed: %w", err)
	}

	return client.Database(database), nil
}
