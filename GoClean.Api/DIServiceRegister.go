package main

import (
	"GoClean/GoClean.Api/Middlewares"
	"GoClean/GoClean.Application/Contract/Database/MongoDB"
	Neo4J3 "GoClean/GoClean.Application/Contract/Database/Neo4J"
	Sql3 "GoClean/GoClean.Application/Contract/Database/Sql"
	"GoClean/GoClean.Application/Contract/Services"
	"GoClean/GoClean.Application/Features/Api/Sample"
	"GoClean/GoClean.Application/Model"
	GoClean_Domain "GoClean/GoClean.Domain"
	MongoDB2 "GoClean/GoClean.Domain/Entities/MongoDB"
	Neo4J2 "GoClean/GoClean.Domain/Entities/Neo4J"
	Sql2 "GoClean/GoClean.Domain/Entities/Sql"
	MongoDB3 "GoClean/GoClean.Infrastructure/Database/MongoDB"
	"GoClean/GoClean.Infrastructure/Database/Neo4J"
	"GoClean/GoClean.Infrastructure/Database/Sql"
	Services2 "GoClean/GoClean.Infrastructure/Services"
	"go.uber.org/dig"
)

func GetDIContainer() *dig.Container {
	container := dig.New()
	container.Provide(func() *Model.TokenUser {
		return &Model.TokenUser{}
	})
	DBRegisteration(projectConfig, container)
	ServiceRegisteration(projectConfig, container)
	HandlerRegisteration(projectConfig, container)
	return container
}
func DBRegisteration(projectConfigs *GoClean_Domain.Configs, container *dig.Container) {
	container.Provide(func() *MongoDB3.MongoBaseRepository[MongoDB2.MongoSampleEntity] {
		service := MongoDB3.NewMongoBaseRepository[MongoDB2.MongoSampleEntity](projectConfigs.MongoConfig)
		return service
	})
	container.Provide(func(mongoBaseRepository *MongoDB3.MongoBaseRepository[MongoDB2.MongoSampleEntity]) *MongoDB.IMongoSampleEntityRepository {
		var service MongoDB.IMongoSampleEntityRepository = MongoDB3.MongoSampleEntityRepository{MongoBaseRepository: mongoBaseRepository}
		return &service
	})

	container.Provide(func() *Neo4J.Neo4JBaseRepository[Neo4J2.Neo4JSampleEntity] {
		service := Neo4J.NewNeo4JBaseRepository[Neo4J2.Neo4JSampleEntity](projectConfigs.Neo4JConfig)
		return service
	})
	container.Provide(func(neoBaseRepository *Neo4J.Neo4JBaseRepository[Neo4J2.Neo4JSampleEntity]) *Neo4J3.INeo4JSampleEntityRepository {
		var service Neo4J3.INeo4JSampleEntityRepository = Neo4J.Neo4JSampleEntityRepository{Neo4JBaseRepository: neoBaseRepository}
		return &service
	})

	container.Provide(func() *Sql.SqlBaseRepository[Sql2.SqlSampleEntity] {
		service := Sql.NewSqlBaseRepository[Sql2.SqlSampleEntity](projectConfigs.SqlConfig)
		return service
	})
	container.Provide(func(sqlBaseRepository *Sql.SqlBaseRepository[Sql2.SqlSampleEntity]) *Sql3.ISqlSampleEntityRepository {
		var service Sql3.ISqlSampleEntityRepository = Sql.SqlSampleEntityRepository{
			SqlBaseRepository: sqlBaseRepository,
		}
		return &service
	})
}
func ServiceRegisteration(projectConfigs *GoClean_Domain.Configs, container *dig.Container) {
	container.Provide(func() *Services.IFileService {
		var service Services.IFileService = Services2.NewFileService(projectConfigs.FileServiceConfig)
		return &service
	})
	container.Provide(func() *Services.ISMSService {
		var service Services.ISMSService = Services2.NewSMSService(projectConfigs.SMSConfigs)
		return &service
	})
	container.Provide(func() *Services.IMessageService {
		var service Services.IMessageService = Services2.MessageService{}
		return &service
	})
	container.Provide(func() *Middlewares.JwtHelper {
		return &Middlewares.JwtHelper{}
	})
}

func HandlerRegisteration(projectConfigs *GoClean_Domain.Configs, container *dig.Container) {
	container.Provide(func(messageService *Services.IMessageService) *Sample.SampleBusinessServiceHandler {
		service := Sample.NewSampleBusinessServiceHandler(messageService)
		return service
	})

}
