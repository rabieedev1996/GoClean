package GoClean_Domain

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoConn struct {
	DbContext context.Context
	Client    *mongo.Client
	Database  *mongo.Database
}
