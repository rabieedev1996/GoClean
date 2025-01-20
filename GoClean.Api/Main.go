package main

import (
	_ "GoClean/GoClean.Api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

var Debug string

// default swagger address : http://localhost:8080/swagger/index.html
// swagger commands for generation docs
// cd .\GoClean.Api\
// "swag init --parseDependency"
func main() {
	Debug = os.Getenv("IS_DEBUG")
	GetProjectConfigs()
	MongoConn = *NewMongoDatabaseConn()
	Neo4JConn = *NewNeo4JConn()
	GORMConn = *NewSQLGormConnecting()
	go mongoMonitorConnection()
	go Neo4jMonitorConnection()
	ginRouter := RegisterGin()
	ginRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	ginRouter.Run(":8080")
}
