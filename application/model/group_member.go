package model

import (
	"github.com/astaxie/beego/orm"
	"github.com/ccyun/GoApp/application/library/conf"
)

//GroupMember 任务表结构
type GroupMember struct {
	isInit    bool      `orm:"-",json:"-"`
	o         orm.Ormer `orm:"-",json:"-"`
	GroupID   uint64    `orm:"column(group_id)"`
	UserID    uint64    `orm:"column(profile_id)"`
	JoinState uint64    `orm:"column(join_state)"`
}

//InitDB 初始化db
func (M *GroupMember) InitDB() error {
	if M.isInit == true {
		return nil
	}
	var dbConfig map[string]string
	if err := conf.JSON("db_dsn_list", &dbConfig); err != nil {
		return err
	}
	pool, _ := conf.Int("db_pool")
	if groupMember, ok := dbConfig["group_member"]; ok {
		if err := orm.RegisterDataBase("group_member", "mysql", groupMember, pool, pool); err != nil {
			return err
		}
	}
	M.o = orm.NewOrm()
	M.o.Using("group_member")
	M.isInit = true
	return nil
}

//GetGroupMember 查询讨论组有效成员
func (M *GroupMember) GetGroupMember(groupID uint64) ([]uint64, error) {
	if err := M.InitDB(); err != nil {
		return nil, err
	}
	var (
		memberData []GroupMember
		userIDs    []uint64
	)
	_, err := M.o.Raw("select profile_id from group_member_0000 where group_id=? and join_state<>1 union select profile_id from group_member_0001 where group_id=? and join_state<>1 union select profile_id from group_member_0002 where group_id=? and join_state<>1 union select profile_id from group_member_0003 where group_id=? and join_state<>1", groupID, groupID, groupID, groupID).QueryRows(&memberData)
	if err != nil {
		return nil, err
	}
	for _, v := range memberData {
		userIDs = append(userIDs, v.UserID)
	}
	return userIDs, nil
}
