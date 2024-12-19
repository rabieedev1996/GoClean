package Neo4J

import "GoClean/GoClean.Domain/Entities/Neo4J"

type Neo4JSampleEntityRepository struct {
	*Neo4JBaseRepository[Neo4J.Neo4JSampleEntity]
}
