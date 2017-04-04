package form2json

import (
	"log"
	"testing"
)

type dss struct {
	User struct {
		Name string `json:"name"`
		Sex  bool   `json:"sex"`
	} `json:"user"`
}

func TestUnmarshal(t *testing.T) {
	s := "user%5Bname%5D=Bob+Smith&user%5Bage%5D=047&user%5Bsex%5D=TRUE&user%5Bdob%5D=5%2F12%2F1956&pastimes%5B0%5D=golf&pastimes%5B1%5D=opera&pastimes%5B2%5D=poker&pastimes%5B3%5D=rap&children%5Bbobby%5D%5Bage%5D=012&children%5Bbobby%5D%5Bsex%5D=M&children%5Bsally%5D%5Bage%5D=8&children%5Bsally%5D%5Bsex%5D=F&flags_0=CEO"
	var ssss dss
	ss, _ := Unmarshal(s, &ssss)

	log.Println(ss)
	log.Println(ssss)
}
