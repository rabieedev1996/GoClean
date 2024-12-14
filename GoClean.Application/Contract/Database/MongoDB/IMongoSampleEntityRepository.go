package MongoDB

import (
	"GoClean/GoClean.Domain/Entities/MongoDB"
)

type IMongoSampleEntityRepository interface {
	IMongoBaseRepository[MongoDB.MongoSampleEntity]
}
