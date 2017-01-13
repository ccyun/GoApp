package hbase

import (
	"log"
	"strconv"
	"testing"

	"git.apache.org/thrift.git/lib/go/thrift"

	"time"
)

//REVERSED =>true
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
	//hostPort = "192.168.40.12:9090"
	hostPort = "192.168.197.128:9090"
	if trans, err = thrift.NewTSocket(hostPort); err != nil {
		t.Error("initHbase NewTSocket error:", err)
	}

	trans.SetTimeout(60 * time.Second)

	client = NewTHBaseServiceClientFactory(trans, thrift.NewTBinaryProtocolFactoryDefault())
	if err = trans.Open(); err != nil {
		t.Error("initHbase NewTSocket open error:", err)
	}

	// log.Println("1")
	// var tputs []*TPut
	// for i := 1000000001; i < 1000160000; i++ {
	// 	k := function.ReverseString(strconv.Itoa(i))

	// 	tput := NewTPut()
	// 	tput.Row = []byte(k)
	// 	tColumnValue := NewTColumnValue()
	// 	tColumnValue.Family = []byte("info")
	// 	tColumnValue.Qualifier = []byte("title")
	// 	tColumnValue.Value = []byte(k)
	// 	tput.ColumnValues = append(tput.ColumnValues, tColumnValue)

	// 	tputs = append(tputs, tput)
	// }
	// log.Println("2")
	// if err = client.PutMultiple([]byte("news"), tputs); err != nil {
	// 	t.Error("Put error:", err)
	// }

	isinit = true
}

func TestHbasePut(t *testing.T) {
	initHbase(t)
	defer trans.Close()
	for i := 1; i < 100; i++ {
		tput := NewTPut()
		tput.Row = []byte("000001")
		tColumnValue := NewTColumnValue()
		tColumnValue.Family = []byte("data")
		tColumnValue.Qualifier = []byte(strconv.Itoa(i))
		tColumnValue.Value = []byte("5")
		tput.ColumnValues = append(tput.ColumnValues, tColumnValue)
		if err = client.Put([]byte("news"), tput); err != nil {
			t.Error("Put error:", err)
		}
	}

}

func TestHbaseGet(t *testing.T) {
	initHbase(t)
	defer trans.Close()
	tget := NewTGet()
	tget.Row = []byte("000001")

	time := int64(1)
	tget.Timestamp = &time
	tcol := NewTColumn()
	tcol.Family = []byte("data")

	tget.Columns = append(tget.Columns, tcol)
	aa, err2 := client.Get([]byte("news"), tget)
	for _, v := range aa.GetColumnValues() {
		log.Println(string(v.GetValue()))
	}

	log.Println(err2)
}

func TestHbaseDel(t *testing.T) {
	initHbase(t)
	defer trans.Close()
	tdelete := NewTDelete()
	tcol := NewTColumn()
	tcol.Family = []byte("data")
	tcol.Qualifier = []byte("feed")
	time := int64(3)
	tcol.Timestamp = &time
	tdelete.Columns = append(tdelete.Columns, tcol)
	tdelete.Row = []byte("000001")

	client.DeleteSingle([]byte("news"), tdelete)

}

func TestHbaseScan(t *testing.T) {
	initHbase(t)
	defer trans.Close()

	tscan := NewTScan()

	Reversed := true
	tscan.Reversed = &Reversed

	tscan.FilterString = []byte("ColumnPaginationFilter(1, 0)")

	sacnid, _ := client.OpenScanner([]byte("news"), tscan)
	sss, _ := client.GetScannerRows(sacnid, 10)

	for _, v := range sss {
		for _, va := range v.GetColumnValues() {
			log.Println(va)
		}

	}
}
