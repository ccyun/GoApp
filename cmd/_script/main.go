package main

import (
	_ "bbs_server/application"
	"bbs_server/application/library/conf"
	"fmt"
	"log"

	"github.com/astaxie/beego/orm"
)

func main() {
	d := orm.NewOrm()
	dbShardDsns := conf.String("db_shard_dsns")
	log.Println(dbShardDsns)
	tableNum, _ := conf.Int("db_shard_table_mun")
	log.Println(tableNum)
	if err := orm.RegisterDataBase("msg0", "mysql", dbShardDsns, 2, 2); err != nil {
		log.Println(err)
		return
	}
	if err := d.Using("msg0"); err != nil {
		log.Println(err)
		return
	}
	// sql1 := `INSERT INTO %s(site_id,board_id,discuss_id,bbs_id,feed_type,feed_id,user_id,user_org_id,task_status,is_read,created_at) (select site_id,board_id,discuss_id,bbs_id,feed_type,1711 feed_id,user_id,user_org_id,task_status,8 is_read,1500861532308 created_at from %s where bbs_id=299 and feed_id=3 and feed_type='bbs')` //聆听优秀的声音——【武明-越努力越幸运】
	// sql2 := `INSERT INTO %s(site_id,board_id,discuss_id,bbs_id,feed_type,feed_id,user_id,user_org_id,task_status,is_read,created_at) (select site_id,board_id,discuss_id,bbs_id,feed_type,1567 feed_id,user_id,user_org_id,task_status,8 is_read,1500024950815 created_at from %s where bbs_id=301 and feed_id=3 and feed_type='bbs')` //聆听优秀的声音——【王鹏-今天你的客户是什么颜色】

	sql1 := "INSERT INTO %s(site_id,board_id,discuss_id,bbs_id,feed_type,feed_id,user_id,user_org_id,task_status,is_read,created_at) (select site_id,board_id,discuss_id,10005548 bbs_id,feed_type,6080 feed_id,user_id,user_org_id,task_status,1 is_read,1500861532308 created_at from %s where bbs_id=10005066 and feed_type='bbs')" //聆听优秀的声音——【武明-越努力越幸运】
	sql2 := "INSERT INTO %s(site_id,board_id,discuss_id,bbs_id,feed_type,feed_id,user_id,user_org_id,task_status,is_read,created_at) (select site_id,board_id,discuss_id,10005372 bbs_id,feed_type,5904 feed_id,user_id,user_org_id,task_status,1 is_read,1500024950815 created_at from %s where bbs_id=10005066 and feed_type='bbs')" //聆听优秀的声音——【王鹏-今天你的客户是什么颜色】10005066

	for i := 0; i < tableNum; i++ {
		tableName := fmt.Sprintf("bbs_msg_%.4d", i)
		d.Raw(fmt.Sprintf("delete from %s where bbs_id=10005548", tableName)).Exec()
		d.Raw(fmt.Sprintf("delete from %s where bbs_id=10005372", tableName)).Exec()

		sql11 := fmt.Sprintf(sql1, tableName, tableName)
		log.Println(sql11)
		if _, err := d.Raw(sql11).Exec(); err != nil {
			log.Println(err)
			return
		}
		sql22 := fmt.Sprintf(sql2, tableName, tableName)
		log.Println(sql22)
		if _, err := d.Raw(sql22).Exec(); err != nil {
			log.Println(err)
			return
		}
	}
}
