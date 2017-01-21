package adapter

import (
	"strconv"

	"encoding/json"

	"fmt"

	"github.com/ccyun/GoApp/application/library/httpcurl"
	"github.com/ccyun/GoApp/application/model"
)

//Bbs 图文广播
type Bbs struct {
	base
}

func init() {
	Register("bbs", new(Bbs))
}

//getBbsTaskInfo 读取广播任务信息
func (B *base) getBbsTaskInfo() error {
	var err error
	model := new(model.BbsTask)
	B.bbsTaskInfo, err = model.GetOne(B.bbsID)

	return err
}

//NewTask 新任务对象
func (B *Bbs) NewTask(task model.Queue) error {
	B.base.NewTask(task)

	var action map[string]string
	if err := json.Unmarshal([]byte(B.action), &action); err != nil {
		return fmt.Errorf("NewTask action Unmarshal error,taskID:%d,action:%s", B.taskID, B.action)
	}
	bbsID, err := strconv.Atoi(action["bbs_id"])
	if err != nil {
		return fmt.Errorf("NewTask strconv.Atoi error,taskID:%d,action:%s", B.taskID, B.action)
	}
	B.bbsID = uint64(bbsID)
	B.attachmentsBase64 = action["attachments_base64"]
	if err := B.getBbsInfo(); err != nil {
		return err
	}
	if err := B.getBoardInfo(); err != nil {
		return err
	}
	///////判断广播类型
	switch B.category {
	case "task":
		if err := B.getBbsTaskInfo(); err != nil {
			return err
		}
	}

	return nil
}

//GetPublishScopeUsers 分析发布范围
func (B *Bbs) GetPublishScopeUsers() error {
	ums := new(httpcurl.UMS)

	userIDs, err := ums.GetAllUserIDsByOrgIDs(B.customerCode, B.bbsInfo.PublishScope.GroupIDs)
	if err != nil {
		return err
	}
	B.PublishScope = make(map[string][]uint64)
	B.PublishScope["group_ids"] = B.bbsInfo.PublishScope.GroupIDs
	B.PublishScope["user_ids"] = B.bbsInfo.PublishScope.UserIDs
	B.userIDs = append(B.bbsInfo.PublishScope.UserIDs, userIDs[0:]...)
	return nil
}

//CreateFeed 创建Feed
func (B *Bbs) CreateFeed() error {

	return nil
}

//CreateRelation 创建接收者关系
func (B *Bbs) CreateRelation() error {
	return nil
}

//CreateUnread 创建未读计数
func (B *Bbs) CreateUnread() error {
	return nil
}

//UpdateStatus 更新状态
func (B *Bbs) UpdateStatus() error {
	return nil
}

//SendMsg 发送消息
func (B *Bbs) SendMsg() error {
	return nil
}
