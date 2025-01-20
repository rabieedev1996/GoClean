package main

import (
	GoClean_Domain "GoClean/GoClean.Domain"
	context2 "context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"time"
)

var MongoConn GoClean_Domain.MongoConn
var Neo4JConn GoClean_Domain.Neo4JConn
var GORMConn GoClean_Domain.GormConn

func NewMongoDatabaseConn() *GoClean_Domain.MongoConn {
	context := context2.TODO()
	clientOption := options.Client().ApplyURI(projectConfig.MongoConfig.Connection)
	client, _ := mongo.Connect(context, clientOption)
	database := client.Database(projectConfig.MongoConfig.DatabaseName)
	return &GoClean_Domain.MongoConn{
		DbContext: context,
		Client:    client,
		Database:  database,
	}
}
func mongoMonitorConnection() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		err := MongoConn.Client.Ping(context.TODO(), nil)
		if err != nil {
			fmt.Println("Connection lost. Attempting to reconnect...")
			// تلاش برای بازیابی اتصال
			mongoReconnectMongo()
		}
	}
}
func mongoReconnectMongo() {
	for {
		clientOption := options.Client().ApplyURI(projectConfig.MongoConfig.Connection)
		context := context2.TODO()
		client, err := mongo.Connect(context, clientOption)
		MongoConn.DbContext = context
		MongoConn.Client = client
		database := client.Database(projectConfig.MongoConfig.DatabaseName)
		MongoConn.Database = database
		if err != nil {
			fmt.Println("Reconnection failed. Retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
		} else {
			fmt.Println("Reconnected to MongoDB.")
			break
		}
	}
}

func NewNeo4JConn() *GoClean_Domain.Neo4JConn {
	uri := projectConfig.Neo4JConfig.URI
	username := projectConfig.Neo4JConfig.Username
	password := projectConfig.Neo4JConfig.Password
	driver, _ := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	return &GoClean_Domain.Neo4JConn{
		Driver: driver,
	}
}

func neo4jPing() error {
	session := Neo4JConn.Driver.NewSession(context.Background(), neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(context.Background())

	_, err := session.Run(context.Background(), "RETURN 1", nil)
	return err
}

func Neo4jMonitorConnection() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		err := neo4jPing()
		if err != nil {
			fmt.Println("Neo4j connection lost. Attempting to reconnect...")
			neo4jReconnect()
		}
	}
}

func neo4jReconnect() {
	for {
		err := neo4jPing()
		if err != nil {
			fmt.Println("Reconnection failed. Retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
		} else {
			fmt.Println("Reconnected to Neo4j.")
			break
		}
	}
}

func NewSQLGormConnecting() *GoClean_Domain.GormConn {
	db, _ := gorm.Open(sqlserver.Open(projectConfig.SqlConfig.Connection), &gorm.Config{})
	sqlDb, _ := db.DB()
	sqlDb.SetMaxOpenConns(200)
	//sqlDb.SetMaxIdleConns(5)
	sqlDb.SetConnMaxLifetime(time.Hour)
	conn := GoClean_Domain.GormConn{
		DB: db,
	}

	return &conn
}
func NewPGSqlGormConnecting() *GoClean_Domain.GormConn {
	db, _ := gorm.Open(postgres.Open(projectConfig.SqlConfig.Connection), &gorm.Config{})
	sqlDb, _ := db.DB()
	sqlDb.SetMaxOpenConns(200)
	//sqlDb.SetMaxIdleConns(5)
	sqlDb.SetConnMaxLifetime(time.Hour)
	conn := GoClean_Domain.GormConn{
		DB: db,
	}
	return &conn
}
