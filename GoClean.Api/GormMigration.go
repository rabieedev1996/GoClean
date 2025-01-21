package main

import "GoClean/GoClean.Domain/Entities/Sql"

func main1() {
	GetProjectConfigs()
	GORMConn := *NewPGSqlGormConnecting()
	GORMConn.DB.Migrator().CreateTable(&Sql.SqlSampleEntity{})
}
