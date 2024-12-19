package Neo4J

import "GoClean/GoClean.Domain/Entities/Neo4J"

type INeo4JSampleEntityRepository interface {
	INeo4JBaseRepository[Neo4J.Neo4JSampleEntity]
}
