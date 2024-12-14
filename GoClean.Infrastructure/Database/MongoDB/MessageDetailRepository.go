package MongoDB

import (
	"GoClean/GoClean.Application/Contract/Database/MongoDB"
	MongoEntities "GoClean/GoClean.Domain/Entities/MongoDB"
)

type MessageDetailRepository struct {
	MongoDB.IMongoBaseRepository[MongoEntities.MongoSampleEntity]
}
