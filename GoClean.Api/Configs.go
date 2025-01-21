package main

import (
	GoClean_Domain "GoClean/GoClean.Domain"
	"os"
	"strconv"
)

var projectConfig *GoClean_Domain.Configs

func getConfigFilePath() string {
	const (
		unixFilePath    = "/home/shawn/.config/RasaAPIConfigsStaging.json"
		windowsFilePath = `C:\Users\Mohammad\Secret\RasaAPIConfigsPublic.json`
	)

	envUnix := os.Getenv("IS_UNIX")
	isUnix, _ := strconv.ParseBool(envUnix)
	if isUnix {
		return unixFilePath
	}
	return windowsFilePath
}

func GetProjectConfigs() *GoClean_Domain.Configs {
	projectConfig = &GoClean_Domain.Configs{
		MongoConfig: GoClean_Domain.Configs_Mongo{
			Connection:   os.Getenv("MONGO_CONN"),
			DatabaseName: os.Getenv("MONGO_DB_NAME"),
		},
		Neo4JConfig: GoClean_Domain.Configs_Neo4J{
			URI:      os.Getenv("NEO4J_URI"),
			Username: os.Getenv("NEO4J_USERNAME"),
			Password: os.Getenv("NEO4J_PASSWORD"),
		},
		SqlConfig: GoClean_Domain.Configs_Sql{
			Connection: os.Getenv("PSQL_CONNECTION"),
		},
		TokeKey:   os.Getenv("TOKEN_KEY"),
		SMSConfig: GoClean_Domain.Configs_SMS{SMSIRToken: os.Getenv("SMSIR_CONFIG")},
	}
	return projectConfig
}
