package adapter

import (
	"encoding/json"
	"log"
	"testing"
)

type aaa struct {
	Name string `json:"name"`
	Age  string `json:"-"`
}

func Test(t *testing.T) {
	s := `{"name":"fdsfdsf","age":"dddd"}`
	var aa aaa

	err := json.Unmarshal([]byte(s), &aa)

	log.Println(err)
	log.Println(aa.Name)
}
