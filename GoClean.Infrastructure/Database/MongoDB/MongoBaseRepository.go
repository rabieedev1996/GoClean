package MongoDB

import (
	Entities "GoClean/GoClean.Domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
)

type MongoBaseRepository[entity any] struct {
	Collection *mongo.Collection
	Context    context.Context
}

func NewMongoBaseRepository[T any](configs Entities.Configs_Mongo) *MongoBaseRepository[T] {
	client, err := mongo.NewClient(options.Client().ApplyURI(configs.Connection))
	if err != nil {
		panic(err)
	}
	ctx := context.TODO()
	err = client.Connect(ctx)

	database := client.Database(configs.DatabaseName)
	collection := database.Collection(Entities.GetAllias[T]())

	return &MongoBaseRepository[T]{
		Context:    ctx,
		Collection: collection,
	}

}

func (r MongoBaseRepository[entity]) Insert(document *entity) (*mongo.InsertOneResult, error) {
	val := reflect.ValueOf(document).Elem()
	idField := val.FieldByName("ID")

	if idField.IsValid() && idField.CanSet() && idField.Type() == reflect.TypeOf("") {
		idField.Set(reflect.ValueOf(primitive.NewObjectID().Hex()))
	}

	result, err := r.Collection.InsertOne(r.Context, document)
	if id, ok := result.InsertedID.(primitive.ObjectID); ok {
		// استفاده از reflection برای اختصاص دادن ID به فیلد entity
		val := reflect.ValueOf(document).Elem() // گرفتن مقدار اصلی از پونتر
		idField := val.FieldByName("ID")        // پیدا کردن فیلد ID

		// بررسی اینکه آیا فیلد ID وجود دارد و از نوع primitive.ObjectID است
		if idField.IsValid() && idField.CanSet() && idField.Type() == reflect.TypeOf("") {
			idField.Set(reflect.ValueOf(id.Hex())) // اختصاص دادن ID به فیلد
		}
	}
	return result, err
}

func (r MongoBaseRepository[entity]) Update(id string, document *entity) (*mongo.UpdateResult, error) {

	//objId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": id}

	// تبدیل entity به bson.M
	update := bson.M{}
	updateData, err := bson.Marshal(document)
	if err != nil {
		return nil, err
	}

	err = bson.Unmarshal(updateData, update)
	if err != nil {
		return nil, err
	}
	delete(update, "_id")

	result, err1 := r.Collection.UpdateOne(r.Context, filter, bson.M{"$set": update})
	return result, err1
}

func (r MongoBaseRepository[entity]) Delete(id string) (*mongo.DeleteResult, error) {
	//objId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": id}
	return r.Collection.DeleteOne(r.Context, filter)
}

func (r MongoBaseRepository[entity]) GetByID(id string) (*entity, error) {
	//objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": id}
	var result entity
	err := r.Collection.FindOne(r.Context, filter).Decode(&result)
	return &result, err
}
func (r MongoBaseRepository[entity]) GetCollection() mongo.Collection {
	return *r.Collection
}
func (r MongoBaseRepository[entity]) GetContext() context.Context {
	return r.Context
}
