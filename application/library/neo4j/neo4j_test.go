package neo4j

import (
	"log"
	"testing"

	neo4j "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

var (
	isInit bool
	pool   neo4j.DriverPool
	err    error
)

func initNeo4j(t *testing.T) {
	if isInit == false {
		if pool, err = neo4j.NewDriverPool("bolt://neo4j:admin@node_b:7687", 10); err != nil {
			t.Error(err)
		}
	}
	isInit = true
}

func TestNeo4j(t *testing.T) {
	initNeo4j(t)
	client, err := pool.OpenPool()
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
	tx, _ := client.Begin()

	data, data2, data3, data4 := client.QueryNeoAll(`MATCH (n1)-[r]->(n2) RETURN r`, nil)
	tx.Rollback()
	log.Println(data)
	log.Println(data2)
	log.Println(data3)
	log.Println(data4)

}
