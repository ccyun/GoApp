package hbase

import (
	"log"
	"strconv"
	"testing"

	"time"

	"github.com/ccyun/GoApp/application/library/hbase/thrift"
)

var (
	hostPort string
	trans    *thrift.TSocket
	client   *THBaseServiceClient
	err      error
	isinit   bool
)

func initHbase(t *testing.T) {
	if isinit {
		return
	}
	hostPort = "192.168.40.12:9090"
	if trans, err = thrift.NewTSocket(hostPort); err != nil {
		t.Error("initHbase NewTSocket error:", err)
	}

	trans.SetTimeout(60 * time.Second)

	client := NewTHBaseServiceClientFactory(trans, thrift.NewTBinaryProtocolFactoryDefault())
	if err = trans.Open(); err != nil {
		t.Error("initHbase NewTSocket open error:", err)
	}
	log.Println("1")
	var tputs []*TPut
	for i := 1000000001; i < 1000160000; i++ {
		k := strconv.Itoa(i)

		tput := NewTPut()
		tput.Row = []byte(k)
		tColumnValue := NewTColumnValue()
		tColumnValue.Family = []byte("info")
		tColumnValue.Qualifier = []byte("title")
		tColumnValue.Value = []byte(k)
		tput.ColumnValues = append(tput.ColumnValues, tColumnValue)
		tputs = append(tputs, tput)
	}
	log.Println("2")
	if err = client.PutMultiple([]byte("news"), tputs); err != nil {
		t.Error("Put error:", err)
	}
	defer func() {
		if err = trans.Close(); err != nil {
			t.Error("initHbase NewTSocket close error:", err)
		}

	}()
	isinit = true
}

func TestHbase(t *testing.T) {
	initHbase(t)

}
