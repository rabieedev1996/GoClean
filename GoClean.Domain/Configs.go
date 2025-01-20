package GoClean_Domain

type Configs struct {
	MongoConfig       Configs_Mongo
	Neo4JConfig       Configs_Neo4J
	SqlConfig         Configs_Sql
	FileServiceConfig Configs_FileService
	SMSConfigs        Configs_SMS
	TokeKey           string
	SMSConfig         Configs_SMS
}
type Configs_Mongo struct {
	Connection   string
	DatabaseName string
}
type Configs_Sql struct {
	Connection string
}
type Configs_Neo4J struct {
	URI      string
	Username string
	Password string
}
type Configs_FileService struct {
	URI      string
	Username string
	Password string
}

type Configs_SMS struct {
	SMSIRToken string
}
