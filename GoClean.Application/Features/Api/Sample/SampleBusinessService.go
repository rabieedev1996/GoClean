package Sample

import (
	"GoClean/GoClean.Application/Contract/Database/MongoDB"
	"GoClean/GoClean.Application/Contract/Database/Neo4J"
	"GoClean/GoClean.Application/Contract/Database/Sql"
	"GoClean/GoClean.Application/Contract/Services"
	MongoDB2 "GoClean/GoClean.Domain/Entities/MongoDB"
	Neo4J2 "GoClean/GoClean.Domain/Entities/Neo4J"
)

type SampleBusinessServiceHandler struct {
	MessageService              *Services.IMessageService
	MongoSampleEntityRepository *MongoDB.IMongoSampleEntityRepository
	SqlSampleEntityRepository   *Sql.ISqlSampleEntityRepository
	Neo4JSampleEntityRepository *Neo4J.INeo4JSampleEntityRepository
}

func NewSampleBusinessServiceHandler(messageService *Services.IMessageService, mongoSampleEntityRepository *MongoDB.IMongoSampleEntityRepository, sqlSampleEntityRepository *Sql.ISqlSampleEntityRepository, neo4JSampleEntityRepository *Neo4J.INeo4JSampleEntityRepository) *SampleBusinessServiceHandler {
	return &SampleBusinessServiceHandler{MessageService: messageService, MongoSampleEntityRepository: mongoSampleEntityRepository, SqlSampleEntityRepository: sqlSampleEntityRepository, Neo4JSampleEntityRepository: neo4JSampleEntityRepository}
}

type SampleBusinessServiceInput struct {
	Field string
}
type SampleBusinessServiceOutput struct {
}

func (receiver SampleBusinessServiceHandler) Handle(input *SampleBusinessServiceInput) *SampleBusinessServiceOutput {

	////Error Generation
	//responseGenerator := Commons.NewResponseGenerator(receiver.MessageService)
	//responseObject := responseGenerator.GetResponseObjectResult(Commons.RESULT_CODE_VALIDATION_ERROR, nil)
	//panic(&responseObject)

	mongoSampleEntity := MongoDB2.NewMongoSampleEntity()
	(*receiver.MongoSampleEntityRepository).Insert(&mongoSampleEntity)

	neoSampleEntity := Neo4J2.NewNeo4JSampleEntity()
	(*receiver.Neo4JSampleEntityRepository).Create(&neoSampleEntity)

	return &SampleBusinessServiceOutput{}
}
