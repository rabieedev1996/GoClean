package Neo4J

import (
	"encoding/json"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func NeoRecordToModel[T any](record *neo4j.Record) *T {
	var item T
	resultJson, _ := json.Marshal(record.Values[0])
	json.Unmarshal(resultJson, &item)
	return &item
}
func NeoValueToModel[T any](value any) *T {
	var item T
	resultJson, _ := json.Marshal(value)
	json.Unmarshal(resultJson, &item)
	return &item
}
