package hbase

import "testing"

var isInit bool

func initHbase() {
	if isInit == false {
		InitHbase("192.168.197.128", "9090", 10)
	}
	isInit = true
}

func TestHbasePut(t *testing.T) {
	initHbase()
	client, _ := OpenClient()
	defer CloseClient(client)
	TPut := &TPut{
		Row: []byte("ddddd"),
		ColumnValues: []*TColumnValue{
			&TColumnValue{
				Family:    []byte("data"),
				Qualifier: []byte("feed"),
				Value:     []byte("ddddd"),
			},
		},
	}

	client.Put([]byte("bbs_feed"), TPut)
}
