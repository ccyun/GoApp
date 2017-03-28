package feed

import (
	"fmt"
	"strings"
	"time"

	"bbs_server/application/module/pic"
)

//GetBbsView 图文广播模板
func GetBbsView(data *Bbs) string {
	description := data.BoardName
	if data.Description != "" {
		description = data.Description
	}
	return `<div style="font-family:PingFangSC-Medium,Microsoft YaHei,Arial,serif;font-size:14px;padding:10px;box-sizing:border-box;text-align:left;" class="bbs"><div style="padding-bottom:5px;"><span style="display:inline-block;overflow:hidden;text-overflow:ellipsis;white-space:nowrap;max-width:100%;">` + data.Title + `</span></div><div style="text-align:center;border-radius:3px;overflow:hidden;height:0;padding-bottom:56.112%"><img src="` + data.Thumb + `" style="width:100%"></div><div style="overflow:hidden;-webkit-line-clamp:2;padding:0 10px;line-height:1.5;text-align:left;height:36px;font-size:12px;color:#666;word-wrap:break-word;text-overflow:ellipsis;margin-top:10px;display:-webkit-box;-webkit-box-orient:vertical;">` + description + `</div></div>`
}

//GetTaskView 广播任务模板
func GetTaskView(data *Task) string {
	endTimeStr := "<!--{{-->不限时<!--}}-->"
	if data.EndTime != 0 {
		endTimeStr = `<!--{{-->任务到期时间<!--}}-->：<span style="color:rgb(249,104,104)">` + time.Unix(int64(data.EndTime/1000), 0).Format("01/02 15:04") + `</span>`
	}
	return `<div style="font-family:PingFangSC-Medium,Microsoft YaHei,Arial,serif;font-size:12px;padding:5px" class="task"><div style="padding:2px 5px;text-align:center;border-radius:3px;display:inline-block;color:#FFF;float:right;"><img src="` + pic.GetFeedIcons("task") + `" style="width:25px;height:25px"></div><div><div style="font-size:16px;color:rgb(59,79,97);line-height:1.8;overflow:hidden;text-overflow:ellipsis;white-space:nowrap;word-wrap:break-word;">` + data.Title + `</div><div style="color:rgb(153,153,153)">` + endTimeStr + `</div></div></div>`
}

//GetTaskListView 广播任务列表模板
func GetTaskListView(data *Task, keyword string) string {
	title := data.Title
	if keyword != "" {
		title = strings.Replace(title, keyword, `<span style="color:rgb(113,174,226)">`+keyword+`</span>`, -1)
	}
	endTimeStr := "<!--{{-->不限时<!--}}-->"
	if data.EndTime != 0 {
		endTimeStr = `<!--{{-->任务到期时间<!--}}-->：<span style="color:rgb(249,104,104)">` + time.Unix(int64(data.EndTime/1000), 0).Format("01/02 15:04") + `</span>`
	}
	taskStatusStr := "<div style=\"padding:1px 3px;text-align:center;min-width:44px;border-radius:2px;display:inline-block;color:#FFF;float:right;margin-top:12px;font-size:14px;background-color:rgb(%s);\"><!--{{-->%s<!--}}--></div>"
	if data.Type == "preview" {
		taskStatusStr = fmt.Sprintf(taskStatusStr, "113,174,226", "预览")
	} else {
		switch data.Status {
		case 3:
			taskStatusStr = fmt.Sprintf(taskStatusStr, "249,104,104", "逾期")
		case 2:
			taskStatusStr = fmt.Sprintf(taskStatusStr, "249,104,104", "未通过")
		case 1:
			taskStatusStr = fmt.Sprintf(taskStatusStr, "173,173,173", "已完成")
		case -1:
			taskStatusStr = fmt.Sprintf(taskStatusStr, "255,160,39", "待处理")
		case 0:
			taskStatusStr = fmt.Sprintf(taskStatusStr, "113,174,226", "待审核")
		case 4:
			taskStatusStr = fmt.Sprintf(taskStatusStr, "173,173,173", "已取消")
		}
	}
	return `<div style="font-family:PingFangSC-Medium,Microsoft YaHei,Arial,serif;font-size:12px;padding:7px 15px" class="' . $feed_info['feed_type'] . '">` + taskStatusStr + `<div><div style="font-size:16px;color:rgb(59,79,97);line-height:26px;overflow:hidden;text-overflow:ellipsis;white-space:nowrap;word-wrap:break-word;margin-right:6px">` + title + `</div><div style="color:rgb(153,153,153)">` + endTimeStr + `</div></div></div>`
}

//GetTaskAuditView 广播任务审核模板
func GetTaskAuditView(data *Task) string {
	message := "审核通过"
	if data.Status == 2 {
		message = "未通过审核"
	}
	return msgTaskView("taskAudit", data.Title, message)
}

//GetTaskCloseView 广播任务关闭模板
func GetTaskCloseView(data *Task) string {
	return msgTaskView("taskClose", data.Title, "已取消")
}

//GetTaskReplyView 广播任务反馈提醒模板
func GetTaskReplyView(data *Task) string {
	return msgTaskView("taskReply", data.Title, "即将逾期，请尽快处理")
}

//_taskView
func msgTaskView(feedType, title, message string) string {
	return `<div style="font-family:PingFangSC-Medium,Microsoft YaHei,Arial,serif;font-size:16px;padding:5px;color:rgb(59,79,97);" class = "` + feedType + `"><!--{{-->你的任务<!--}}--> <span style="color:rgb(113,174,226)">` + title + `</span> <!--{{-->` + message + `<!--}}--></div>`
}

//GetFromView 广播任务模板
func GetFromView(data *Form) string {
	return ""
}
