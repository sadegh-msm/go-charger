package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Close this function will close connection to database if user cancel it or when program ends.
func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

// Connect this function will connect to database by url passed in main function
func Connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	return client, ctx, cancel, err
}

func InsertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {
	// select database and collection ith Client.Database method and Database.Collection method
	collection := client.Database(dataBase).Collection(col)
	result, err := collection.InsertOne(ctx, doc)

	return result, err
}

func InsertMany(client *mongo.Client, ctx context.Context, dataBase, col string, docs []interface{}) (*mongo.InsertManyResult, error) {
	// select database and collection ith Client.Database method and Database.Collection method
	collection := client.Database(dataBase).Collection(col)
	result, err := collection.InsertMany(ctx, docs)

	return result, err
}

func UpdateOne(client *mongo.Client, ctx context.Context, dataBase, col string, filter, update interface{}) (result *mongo.UpdateResult, err error) {
	// select the database and the collection
	collection := client.Database(dataBase).Collection(col)

	// A single document that match with the filter will get updated. update contains the filed which should get updated.
	result, err = collection.UpdateOne(ctx, filter, update)

	return
}

// Query this function helps to write query to find the wanted result
func Query(client *mongo.Client, ctx context.Context, dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error) {
	// select database and collection.
	collection := client.Database(dataBase).Collection(col)

	// collection has a method Find, that returns a mongo.cursor based on query and field.
	result, err = collection.Find(ctx, query, options.Find().SetProjection(field))

	return
}
