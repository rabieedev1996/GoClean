package MongoDB

type MongoSampleEntity struct {
	*MongoBaseEntity `bson:",inline"`
}

func NewMongoSampleEntity() MongoSampleEntity {
	return MongoSampleEntity{
		MongoBaseEntity: NewMongoBaseEntity(),
	}
}
