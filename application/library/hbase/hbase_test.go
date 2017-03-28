package hbase

import (
	"encoding/json"
	"log"
	"strconv"
	"testing"

	"github.com/astaxie/beego/config"
	"bbs_server/application/function"
)

var (
	userID, boardID, feedID, bbsID, feedType string
	feedIDInt, bbsIDInt                      int64
)

//Conf 配置
var Conf config.Configer

func initHbase() {
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
				Host string `json:"host"`
				Port string `json:"port"`
				Pool int    `json:"pool"`
			}
		)
		if err = json.Unmarshal([]byte(Conf.String("hbase")), &config); err != nil {
			return err
		}
		return Init(config.Host, config.Port, config.Pool)
	})

	userID = "63706854"
	boardID = "50000124"
	feedID = "1955"
	feedType = "bbs"
	bbsID = "50001544"
	feedIDInt, _ = strconv.ParseInt(feedID, 10, 0)
	bbsIDInt, _ = strconv.ParseInt(bbsID, 10, 0)

}

func TestHbasePut(t *testing.T) {
	initHbase()
	client, _ := OpenClient()
	defer CloseClient(client)

	TPuts := []*TPut{
		&TPut{
			Row: []byte(userID + "_home"),
			ColumnValues: []*TColumnValue{
				&TColumnValue{
					Family:    []byte("cf"),
					Qualifier: []byte(boardID),
					Value:     []byte(bbsID),
					Timestamp: &feedIDInt,
				},
			},
		},
		&TPut{
			Row: []byte(userID + "_" + feedType),
			ColumnValues: []*TColumnValue{
				&TColumnValue{
					Family:    []byte("cf"),
					Qualifier: []byte(boardID),
					Value:     []byte(bbsID),
					Timestamp: &feedIDInt,
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
		Row: []byte(userID + "_" + feedType),
		Columns: []*TColumn{
			&TColumn{
				Family:    []byte("cf"),
				Qualifier: []byte(boardID),
				Timestamp: &feedIDInt,
			},
		},
		DeleteType: TDeleteType_DELETE_COLUMN,
	}
	log.Println(feedIDInt)
	if err := client.DeleteSingle([]byte("bbs_feed"), tdel); err != nil {
		t.Error(err)
	}
}

func TestHbaseGet(t *testing.T) {
	initHbase()
	client, _ := OpenClient()
	defer CloseClient(client)
	var maxV int32 = 105
	//minStamp := int64()
	log.Println(function.MakeRowkey(int64(63706854)) + "_" + feedType)
	tget := &TGet{
		//	rowkey := function.MakeRowkey(int64(u))

		Row: []byte(function.MakeRowkey(int64(63706854)) + "_" + feedType),
		Columns: []*TColumn{
			&TColumn{
				Family:    []byte("cf"),
				Qualifier: []byte(boardID),
			},
		},
		MaxVersions: &maxV,
		TimeRange: &TTimeRange{
			//MinStamp: int64(1922),
			MaxStamp: int64(1919),
		},
	}

	result, _ := client.Get([]byte("bbs_feed"), tget)

	for _, v := range result.GetColumnValues() {
		log.Println(v.GetTimestamp())
		log.Println(string(v.GetValue()))
	}
}
