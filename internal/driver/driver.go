package driver

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func Connect()(*mongo.Database, error) {
	uri := "mongodb://root:password@mongodb"
	ctx, _ := context.WithTimeout(context.Background(),
		30 * time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	database := client.Database("FeedBackTestDatabase") // db name
	return database, err
}

func close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc){

	// CancelFunc to cancel to context
	defer cancel()

	defer func(){
		if err := client.Disconnect(ctx); err != nil{
			panic(err)
		}
	}()
}