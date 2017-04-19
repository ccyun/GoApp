package adapter

import (
	"fmt"
	"strconv"

	"bbs_server/application/function"
	"bbs_server/application/model"
)

//TaskUnreplyCount 广播任务提醒反馈
type TaskUnreplyCount struct {
	base
}

func init() {
	Register("taskUnreplyCount", func() Tasker {
		return new(TaskUnreplyCount)
	})
}

//NewTask 新任务对象
func (T *TaskUnreplyCount) NewTask(task model.Queue) error {
	T.base.NewTask(task)
	bbsID, err := strconv.Atoi(T.action)
	if err != nil {
		return fmt.Errorf("NewTask strconv.Atoi error,taskID:%d,action:%s", T.taskID, T.action)
	}
	T.bbsID = uint64(bbsID)
	if err := T.getBbsInfo(); err != nil {
		return err
	}
	return nil
}

//GetPublishScopeUsers 分析发布范围
func (T *TaskUnreplyCount) GetPublishScopeUsers() error {
	userIDs, err := new(model.BbsTaskReply).GetReplyUserIDs(T.bbsID)
	if err != nil {
		return err
	}
	T.userIDs = function.SliceDiff(T.bbsInfo.PublishScopeUserIDsArr, userIDs).Uint64()
	return nil
}

//CreateUnread 处理接收者关系
func (T *TaskUnreplyCount) CreateUnread() error {
	if len(T.userIDs) == 0 {
		return nil
	}
	_, err := T.o.QueryTable(new(model.BbsTaskUnreplyCount)).Filter("SiteID", T.siteID).Filter("BoardID", T.boardID).Filter("BbsID", T.bbsID).Filter("UserID__in", T.userIDs).Delete()
	return err
}
