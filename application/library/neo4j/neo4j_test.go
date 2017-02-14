package neo4j

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/astaxie/beego/config"
)

//Conf 配置
var Conf config.Configer

func initNeo4j() {
	func(funcs ...func() error) {
		for _, f := range funcs {
			if err := f(); err != nil {
				panic(err)
			}
		}
	}(func() error {
		conf, err := config.NewConfig("ini", "../../../cmd/TaskScript/conf.ini")
		if err != nil {
			return err
		}
		Conf = conf
		return nil
	}, func() error {
		var (
			err    error
			config struct {
				Host     string `json:"host"`
				Port     string `json:"port"`
				UserName string `json:"username"`
				Password string `json:"password"`
				Pool     int    `json:"pool"`
			}
		)
		if err = json.Unmarshal([]byte(Conf.String("neo4j")), &config); err != nil {
			return err
		}
		return Init(config.Host, config.Port, config.UserName, config.Password, config.Pool)
	})
}

func TestNeo4j(t *testing.T) {
	initNeo4j()
	client, err := OpenClient()
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
