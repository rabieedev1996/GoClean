package Neo4J

import "github.com/neo4j/neo4j-go-driver/v5/neo4j"

type INeo4JBaseRepository[TEntity any] interface {
	Create(entity *TEntity) (int64, error)
	GetById(id int64) (*TEntity, error)
	Update(entity *TEntity, id int64) error
	Delete(id int64) error
	GetAll() ([]*TEntity, error)
	GetDriver() neo4j.DriverWithContext
}
