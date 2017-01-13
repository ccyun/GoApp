package hbase

import (
	"log"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/astaxie/beego/logs"
)

var (
	hostPort string
	trans    *thrift.TSocket
	client   *THBaseServiceClient
	err      error
)

func initHbase() {
	hostPort = "192.168.197.128:9090"
	if trans, err = thrift.NewTSocket(hostPort); err != nil {
		logs.Error("thrift NewTSocket error:", err)
	}
	trans.SetTimeout(60 * time.Second)
	client = NewTHBaseServiceClientFactory(trans, thrift.NewTBinaryProtocolFactoryDefault())
	if err = trans.Open(); err != nil {
		logs.Error("trans Open error:", err)
	}
}

//DDL 表结构
type DDL struct {
	RowKey    []byte
	Family    []byte
	Qualifier []byte
	Value     []byte
}

//Puts 批量写入数据
func Puts(tableName string, datas []DDL) error {
	initHbase()
	defer trans.Close()
	var tputs []*TPut
	for _, d := range datas {
		tput := NewTPut()
		tput.Row = d.RowKey
		tColumnValue := NewTColumnValue()
		tColumnValue.Family = d.Family
		tColumnValue.Qualifier = d.Qualifier
		tColumnValue.Value = d.Value
		tput.ColumnValues = append(tput.ColumnValues, tColumnValue)
		tputs = append(tputs, tput)
	}
	if err := client.PutMultiple([]byte(tableName), tputs); err != nil {
		logs.Error("client PutMultiple error:", err)
		return err
	}
	return nil
}

//GetLastOne 读取单条数据
func GetLastOne(tableName, rowKey, family, qualifier string) error {
	initHbase()
	defer trans.Close()
	tscan := NewTScan()
	Reversed := true
	tscan.Reversed = &Reversed
	tcolumn := NewTColumn()
	tcolumn.Family = []byte(family)
	tcolumn.Qualifier = []byte(qualifier)
	tscan.Columns = append(tscan.Columns, tcolumn)
	tscan.StartRow = []byte("080001:LastFeed:100000000000112")
	tscan.FilterString = []byte("PrefixFilter('" + rowKey + "')")
	openScannerID, err := client.OpenScanner([]byte(tableName), tscan)

	if err != nil {
		logs.Error("client OpenScanner error:", err)
		return err
	}
	defer client.CloseScanner(openScannerID)

	datas, err := client.GetScannerRows(openScannerID, 1)
	if err != nil {
		logs.Error("client OpenScanner error:", err)
		return err
	}
	for _, v := range datas {
		log.Println(string(v.GetRow()))
		for _, vv := range v.GetColumnValues() {
			log.Println(string(vv.GetValue()))
		}
	}
	return nil
}
