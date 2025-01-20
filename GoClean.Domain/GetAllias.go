package GoClean_Domain

import "reflect"

func GetAllias[T any]() string {
	var t T
	typeName := reflect.TypeOf(t).Name()
	switch typeName {
	case "MongoSampleEntity":
		return "mongo_sample_entity"
	case "Neo4JSampleEntity":
		return "neo_sample_entity"
	}
	return ""
}
