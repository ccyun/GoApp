package feed

//GetBbsView 反馈图文广播模板
func GetBbsView(data Customizer) string {
	description := data.BoardName
	if data.Description != "" {
		description = data.Description
	}
	return `<div style="font-family:PingFangSC-Medium,Microsoft YaHei,Arial,serif;font-size:14px;padding:10px;box-sizing:border-box;text-align:left;" class="bbs"><div style="padding-bottom:5px;"><span style="display:inline-block;overflow:hidden;text-overflow:ellipsis;white-space:nowrap;max-width:100%;">` + data.Title + `</span></div><div style="text-align:center;border-radius:3px;overflow:hidden;height:0;padding-bottom:56.112%"><img src="` + data.Thumb + `" style="width:100%"></div><div style="overflow:hidden;-webkit-line-clamp:2;padding:0 10px;line-height:1.5;text-align:left;height:36px;font-size:12px;color:#666;word-wrap:break-word;text-overflow:ellipsis;margin-top:10px;display:-webkit-box;-webkit-box-orient:vertical;">` + description + `</div></div>`
}
