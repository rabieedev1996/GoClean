package Neo4J

import (
	Rasa_Common "GoClean/GoClean.Common"
	Entities "GoClean/GoClean.Domain"
	"GoClean/GoClean.Domain/Entities/Neo4J"
	"context"
	"encoding/json"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Neo4JBaseRepository[T any] struct {
	Driver neo4j.DriverWithContext
	Label  string
}

func NewNeo4JBaseRepository[T any](configs Entities.Configs_Neo4J) *Neo4JBaseRepository[T] {
	uri := configs.URI
	username := configs.Username
	password := configs.Password
	driver, _ := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	label := Entities.GetAllias[T]()
	return &Neo4JBaseRepository[T]{
		Driver: driver,
		Label:  label,
	}
}

func (r *Neo4JBaseRepository[T]) Create(entity *T) (int64, error) {
	jsonStr, _ := json.Marshal(*entity)
	props1 := make(map[string]any)

	err := json.Unmarshal([]byte(jsonStr), &props1)
	props := map[string]any{
		"props": &props1,
	}
	result, err := neo4j.ExecuteQuery(context.Background(), r.Driver,
		`CREATE (n:`+r.Label+` $props) RETURN id(n) AS recordId`,
		props, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		panic(err)
	}
	rId, _ := result.Records[0].Get("recordId")
	return rId.(int64), nil
}
func (r *Neo4JBaseRepository[T]) GetById(id int64) (*T, error) {
	res, err := neo4j.ExecuteQuery(context.Background(), r.Driver,
		`MATCH (n:`+r.Label+`) WHERE id(n) = $id RETURN properties(n)`,
		map[string]interface{}{
			"id": id,
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		panic(err)
	}

	return Neo4J.NeoRecordToModel[T](res.Records[0]), err
}

func (r *Neo4JBaseRepository[T]) Update(entity *T, Id int64) error {
	session := r.Driver.NewSession(context.Background(), neo4j.SessionConfig{})
	defer session.Close(context.Background())
	jsonStr, _ := json.Marshal(*entity)
	props1 := make(map[string]any)
	json.Unmarshal([]byte(jsonStr), &props1)
	_, err := session.ExecuteWrite(context.Background(), func(tx neo4j.ManagedTransaction) (any, error) {
		query := `MATCH (n:` + r.Label + `) where id(n)=$id SET n += $props RETURN n`
		_, err := tx.Run(context.Background(), query, map[string]interface{}{
			"id":    Id, // فرض کنید ID در ساختار T وجود دارد
			"props": props1,
		})
		return nil, err
	})
	return err
}

func (r *Neo4JBaseRepository[T]) Delete(id int64) error {
	session := r.Driver.NewSession(context.Background(), neo4j.SessionConfig{})
	defer session.Close(context.Background())

	_, err := session.ExecuteWrite(context.Background(), func(tx neo4j.ManagedTransaction) (any, error) {
		query := `MATCH (n:` + r.Label + `) WHERE id(n) = $id DELETE n`
		_, err := tx.Run(context.Background(), query, map[string]interface{}{
			"id": id,
		})
		return nil, err
	})
	return err
}

func (r *Neo4JBaseRepository[T]) GetAll() ([]*T, error) {
	var result Rasa_Common.Slice[*T]
	res, err := neo4j.ExecuteQuery(context.Background(), r.Driver,
		`MATCH (n:`+r.Label+`) RETURN properties(n)`,
		nil,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		panic(err)
	}

	for _, record := range res.Records {
		result.Add(Neo4J.NeoRecordToModel[T](record))
	}

	return result, err
}

func (r *Neo4JBaseRepository[T]) GetDriver() neo4j.DriverWithContext {
	return r.Driver
}
