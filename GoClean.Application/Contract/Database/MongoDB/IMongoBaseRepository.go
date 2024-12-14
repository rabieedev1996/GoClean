package MongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type IMongoBaseRepository[entity any] interface {
	Insert(document *entity) (*mongo.InsertOneResult, error)
	Update(id string, entity *entity) (*mongo.UpdateResult, error)
	Delete(id string) (*mongo.DeleteResult, error)
	GetByID(id string) (*entity, error)
	GetCollection() mongo.Collection
	GetContext() context.Context
}
