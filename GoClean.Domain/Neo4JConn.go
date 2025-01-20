package GoClean_Domain

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Neo4JConn struct {
	Driver neo4j.DriverWithContext
}
