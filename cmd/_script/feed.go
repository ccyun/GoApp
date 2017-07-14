package main

import (
	"bbs_server/application/model"
	"bbs_server/application/module/feed"
	"encoding/json"
)

//handleFeed 处理关系
func (T *task) handleFeed() error {
	feedData := model.FeedData{
		Title:          T.bbsInfo.Title,
		Description:    T.bbsInfo.Description,
		CreatedAt:      T.bbsInfo.PublishAt,
		UserID:         T.bbsInfo.UserID,
		Type:           T.bbsInfo.Type,
		Category:       T.category,
		CommentEnabled: T.bbsInfo.CommentEnabled,
		IsBrowser:      T.bbsInfo.IsBrowser,
		IsAuth:         T.bbsInfo.IsAuth,
	}
	if T.category == "bbs" {
		feedData.Thumb = T.bbsInfo.Thumb
		feedData.Link = T.bbsInfo.Link
	} else if T.category == "task" {
		feedData.EndTime = T.taskInfo.EndTime
		feedData.AllowExpired = T.taskInfo.AllowExpired
		feedData.Status = feed.BbsTaskStatus
	}

	feedList := []model.Feed{}
	if _, err := o.Raw("select id,bbs_id,feed_type,data,created_at from bbs_feed where bbs_id=?", T.bbsInfo.ID).QueryRows(&feedList); err != nil {
		return err
	}
	for _, v := range feedList {
		switch v.FeedType {
		case "task":
			feedData.Status = feed.BbsTaskStatus
		case "taskReply":
			feedData.Status = feed.BbsTaskReplyStatus
			feedData.CreatedAt = v.CreatedAt
		case "taskClose":
			feedData.Status = feed.BbsTaskCloseStatus
			feedData.CreatedAt = v.CreatedAt
		case "taskAudit":
			fD := feedDataer{}
			if err := json.Unmarshal([]byte(v.Data), &fD); err != nil {
				return err
			}
			feedData.Status = fD.Status
			feedData.CreatedAt = v.CreatedAt
		}
		s, _ := json.Marshal(feedData)
		if _, err := o.Raw("UPDATE bbs_feed SET data = ? where id=?", string(s), v.ID).Exec(); err != nil {
			return err
		}
		if userIDs, ok := T.publishScopeUserIDs[v.ID]; ok {
			msgData := model.Msg{
				SiteID:    T.bbsInfo.SiteID,
				BoardID:   T.bbsInfo.BoardID,
				DiscussID: T.bbsInfo.DiscussID,
				BbsID:     T.bbsInfo.ID,
				FeedID:    v.ID,
				FeedType:  v.FeedType,
				CreatedAt: v.CreatedAt,
			}
			if err := createRelation(msgData, userIDs, T.taskStatus); err != nil {
				return err
			}
			if v.FeedType == "bbs" || v.FeedType == "task" {
				if _, err := o.Raw("UPDATE bbs_bbs SET msg_count = ? where id=?", len(T.publishScopeUserIDs[v.ID]), T.bbsInfo.ID).Exec(); err != nil {
					return err
				}
				if v.FeedType == "task" {
					//chuli//图片类的子任务更新到子任务扩展表和审核表

				}
			}
		}
	}
	return nil
}
