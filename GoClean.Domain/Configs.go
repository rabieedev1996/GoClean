package GoClean_Domain

type Configs struct {
	MongoConfig Configs_Mongo
	Neo4JConfig Configs_Neo4J
	TokeKey     string
}
type Configs_Mongo struct {
	Connection   string
	DatabaseName string
}
type Configs_Neo4J struct {
	URI      string
	Username string
	Password string
}
