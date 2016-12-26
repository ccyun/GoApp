package log

import (
	"encoding/json"
	"log"
	"testing"
)

func A(arg ...interface{}) {

	val, err := json.Marshal(arg)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(val))
}

func TestSyslog(t *testing.T) {
	A(1, "1", true, []string{"a", "b"})
}
