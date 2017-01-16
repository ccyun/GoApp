package hbase

import (
	"log"
	"strconv"
	"testing"
)

var (
	isInit                             bool
	user_id, board_id, feed_id, bbs_id string
	feed_id_int, bbs_id_int            int64
)

func initHbase() {
	if isInit == false {
		InitHbase("192.168.197.128", "9090", 10)
		user_id = "63669051"
		board_id = "50000116"
		feed_id = "1921"
		bbs_id = "50001588"
		feed_id_int, _ = strconv.ParseInt(feed_id, 10, 0)
		bbs_id_int, _ = strconv.ParseInt(bbs_id, 10, 0)
	}

	isInit = true
}

func TestHbasePut(t *testing.T) {
	initHbase()
	client, _ := OpenClient()
	defer CloseClient(client)

	TPuts := []*TPut{
		&TPut{
			Row: []byte(user_id + "_home"),
			ColumnValues: []*TColumnValue{
				&TColumnValue{
					Family:    []byte("cf"),
					Qualifier: []byte(board_id),
					Value:     []byte(bbs_id),
					Timestamp: &feed_id_int,
				},
			},
		},
		&TPut{
			Row: []byte(user_id + "_list"),
			ColumnValues: []*TColumnValue{
				&TColumnValue{
					Family:    []byte("cf"),
					Qualifier: []byte(board_id),
					Value:     []byte(bbs_id),
					Timestamp: &feed_id_int,
				},
			},
		},
	}

	if err := client.PutMultiple([]byte("bbs_feed"), TPuts); err != nil {
		t.Error(err)
	}
}

func TestHbaseDel(t *testing.T) {
	initHbase()
	client, _ := OpenClient()
	defer CloseClient(client)
	tdel := &TDelete{
		Row: []byte(user_id + "_list"),
		Columns: []*TColumn{
			&TColumn{
				Family:    []byte("cf"),
				Qualifier: []byte(board_id),
				Timestamp: &bbs_id_int,
			},
		},
		DeleteType: TDeleteType_DELETE_COLUMN,
	}
	if err := client.DeleteSingle([]byte("bbs_feed"), tdel); err != nil {
		t.Error(err)
	}
}

func TestHbaseGet(t *testing.T) {
	initHbase()
	client, _ := OpenClient()
	defer CloseClient(client)
	var maxV int32 = 5
	//minStamp := int64()
	tget := &TGet{
		Row: []byte(user_id + "_list"),
		Columns: []*TColumn{
			&TColumn{
				Family:    []byte("cf"),
				Qualifier: []byte(board_id),
			},
		},
		MaxVersions: &maxV,
		TimeRange: &TTimeRange{
			MinStamp: int64(1922),
			MaxStamp: int64(192100000000000000),
		},
	}

	result, _ := client.Get([]byte("bbs_feed"), tget)
	for _, v := range result.GetColumnValues() {
		log.Println(v.GetTimestamp())
		log.Println(string(v.GetValue()))
	}
}
