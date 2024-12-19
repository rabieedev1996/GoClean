package MongoDB

import "GoClean/GoClean.Domain/Entities/MongoDB"

type MongoSampleEntityRepository struct {
	*MongoBaseRepository[MongoDB.MongoSampleEntity]
}
