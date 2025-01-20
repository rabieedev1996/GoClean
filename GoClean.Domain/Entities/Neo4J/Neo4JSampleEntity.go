package Neo4J

type Neo4JSampleEntity struct {
	*Neo4JBaseEntity
}

func NewNeo4JSampleEntity() Neo4JSampleEntity {
	return Neo4JSampleEntity{
		Neo4JBaseEntity: NewNeo4JBaseEntity(),
	}
}
