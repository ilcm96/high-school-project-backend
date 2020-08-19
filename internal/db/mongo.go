package db

import (
	"context"
	"runtime"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func mongodbURI() string {
	os := runtime.GOOS
	if os == "windows" {
		return "mongodb://0.0.0.0:27017"
	}
	return "mongodb://root:root@db.sub02111041190.generalvcn.oraclevcn.com:27017"
}

func Client() *mongo.Client {
	clientOptions := options.Client().ApplyURI(mongodbURI())
	client, _ := mongo.Connect(context.TODO(), clientOptions)
	return client
}
