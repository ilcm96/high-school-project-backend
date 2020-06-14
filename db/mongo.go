package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Client() *mongo.Client {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://root:ilcmgpro903@db.sub02111041190.generalvcn.oraclevcn.com:27017")
	client, _ := mongo.Connect(ctx, clientOptions)
	return client
}
