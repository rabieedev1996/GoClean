package main

import (
	GoClean_Domain "GoClean/GoClean.Domain"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

var projectConfig *GoClean_Domain.Configs

func GetProjectConfigs() *GoClean_Domain.Configs {
	if projectConfig != nil {
		return projectConfig
	}
	isDebug, _ := strconv.ParseBool(Debug)
	if isDebug {
		filePath := `C:\Users\Rabiee\Secret\RasaAPIConfigs.json`
		content, _ := os.ReadFile(filePath)
		projectConfig = &GoClean_Domain.Configs{}
		err := json.Unmarshal(content, projectConfig)
		if err != nil {
			fmt.Print(err.Error())
		}

	} else {
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
			SqlConfig:         GoClean_Domain.Configs_Sql{Connection: os.Getenv("SQL_CONN")},
			TokeKey:           os.Getenv("TOKEN_KEY"),
			SMSConfigs:        GoClean_Domain.Configs_SMS{SMSIRToken: os.Getenv("SMSIR_CONFIG")},
			FileServiceConfig: GoClean_Domain.Configs_FileService{},
		}
	}
	return projectConfig
}
