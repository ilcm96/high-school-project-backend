package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Client() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://root:ilcmgpro903@db.sub02111041190.generalvcn.oraclevcn.com:27017")
	client, _ := mongo.Connect(context.TODO(), clientOptions)
	return client
}
