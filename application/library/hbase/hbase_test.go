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
				Value:     []byte("ddddd32"),
			},
		},
	}
	if err := client.Put([]byte("news"), TPut); err != nil {
		t.Error(err)
	}
}

func TestHbaseDel(t *testing.T) {
	initHbase()
	client, _ := OpenClient()
	defer CloseClient(client)
	tdel := &TDelete{
		Row: []byte("ddddd"),
		// Columns: []*TColumn{
		// 	&TColumn{
		// 		Family:    []byte("data"),
		// 		Qualifier: []byte("feed"),
		// 		//	Timestamp:
		// 	},
		// },
		DeleteType: TDeleteType_DELETE_COLUMN,
	}
	if err := client.DeleteSingle([]byte("news"), tdel); err != nil {
		t.Error(err)
	}
}
