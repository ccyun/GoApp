package pic

import "fmt"

//BoardAvatar 默认头像
type BoardAvatar struct {
	OnCloud   []string
	TestCloud []string
}

//BbsThumb 默认封面
type BbsThumb struct {
	OnCloud   map[string][]string
	TestCloud map[string][]string
}

//FeedIcons Feed图标
type FeedIcons struct {
	OnCloud   map[string]string
	TestCloud map[string]string
}

var (
	//BoardAvatarData 默认头像
	BoardAvatarData BoardAvatar
	//BbsThumbData 默认封面
	BbsThumbData BbsThumb
	//FeedIconsData Feed图标
	FeedIconsData FeedIcons
	//ServerName 运行环境
	ServerName string
)

func init() {
	BoardAvatarData.TestCloud = []string{
		"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_pu5jorqTlpLTlg48gMF8yNDBweC5wbmdeXl50YW5naGRmc15eXmFlMmViN2IxMWIxYWM0NWU2MmQ4YmYzMThkYTFjMzkzXl5edGFuZ2hkZnNeXl4xMjI5NA$&u=62051318",
		"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_pu5jorqTlpLTlg48gMV8yNDBweC5wbmdeXl50YW5naGRmc15eXjBmZmZjMGY2Zjg4NWE2YTFhN2RlZGU0YzM3MmE2NTJjXl5edGFuZ2hkZnNeXl4xMzQzNA$&u=62051318",
		"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_pu5jorqTlpLTlg48gMl8yNDBweC5wbmdeXl50YW5naGRmc15eXmY5OTNmNWUwMDZiOTgzNmE0MjlmODMwNWU3ZmQzNjU0Xl5edGFuZ2hkZnNeXl45NTg4$&u=62051318",
		"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_pu5jorqTlpLTlg48gM18yNDBweC5wbmdeXl50YW5naGRmc15eXjQ0ODIyNDA2M2U5ZGIzYjc1MGIzNWYwMzI0ZmQ1NzFhXl5edGFuZ2hkZnNeXl4xMjAxOQ$&u=62051318",
		"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_pu5jorqTlpLTlg48gNF8yNDBweC5wbmdeXl50YW5naGRmc15eXjlmY2I2ZTJjM2NlMWE2MmVjMzAyNWNmMWE0YzUxYzY0Xl5edGFuZ2hkZnNeXl45MTgy$&u=62051318",
		"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_pu5jorqTlpLTlg48gNV8yNDBweC5wbmdeXl50YW5naGRmc15eXjQ4MDM1OWU2YzU4NDI3ZGU3M2FlYmIwOTE5NTdjN2YyXl5edGFuZ2hkZnNeXl4xMjk0OA$&u=62051318",
		"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_pu5jorqTlpLTlg48gNl8yNDBweC5wbmdeXl50YW5naGRmc15eXjM2MTRlYWQ4M2MzOTdmYjU4NGFmNDEyMjNiMzYzMmQ1Xl5edGFuZ2hkZnNeXl45NDc5$&u=62051318",
		"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_pu5jorqTlpLTlg48gN18yNDBweC5wbmdeXl50YW5naGRmc15eXjhmMmMyNTVmMmRkZWNhYWE0ODc1N2U4MjVmMThlMDdjXl5edGFuZ2hkZnNeXl4xMjE4Nw$&u=62051318",
		"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_pu5jorqTlpLTlg48gOF8yNDBweC5wbmdeXl50YW5naGRmc15eXjI2ZWJmZjYxMDkwNGIwODE0OGM5NTI5ODRiYjFjOTUwXl5edGFuZ2hkZnNeXl4xMTYwOA$&u=62051318",
		"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_pu5jorqTlpLTlg48gOV8yNDBweC5wbmdeXl50YW5naGRmc15eXjg5NzYxOTcwODIyNzIyMjc4NmEyODY0OGFjYjhhZmRhXl5edGFuZ2hkZnNeXl4xMTMxNg$&u=62051318",
		"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_pu5jorqTlpLTlg48gMTBfMjQwcHgucG5nXl5edGFuZ2hkZnNeXl5mYTU1MjFlOGUxZjkzM2Y3Mjg4MDJiZTc4NTY5Mzc3YV5eXnRhbmdoZGZzXl5eMTMyOTE$&u=62051318",
		"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_pu5jorqTlpLTlg48gMTFfMjQwcHgucG5nXl5edGFuZ2hkZnNeXl4wZmJjMDY4MDcyZDk1M2RmNmRmN2QyZDdhNmFlMTYwZl5eXnRhbmdoZGZzXl5eOTE1OA$&u=62051318",
		"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_pu5jorqTlpLTlg48gMTJfMjQwcHgucG5nXl5edGFuZ2hkZnNeXl42NzM5NTExNWFjMmFmM2NmMzk2MjYwNDNlMGI1Y2ZlOF5eXnRhbmdoZGZzXl5eMTAwMDE$&u=62051318",
		"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_pu5jorqTlpLTlg48gMTNfMjQwcHgucG5nXl5edGFuZ2hkZnNeXl44YzU5YTlmOTJmNmZkNWVhMTAzZGZkNzJjOGIyNzE5ZF5eXnRhbmdoZGZzXl5eMTIzOTY$&u=62051318",
		"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_pu5jorqTlpLTlg48gMTRfMjQwcHgucG5nXl5edGFuZ2hkZnNeXl44MWYyOTk1YzBiMTM3NmM1MWY5OTc3N2YzZDFhZDEzN15eXnRhbmdoZGZzXl5eMTEzODQ$&u=62051318",
	}
	BoardAvatarData.OnCloud = []string{
		"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-m7mOiupOWktOWDjyAwXzI0MHB4LnBuZ15eXnRhbmdoZGZzXl5eYWUyZWI3YjExYjFhYzQ1ZTYyZDhiZjMxOGRhMWMzOTNeXl50YW5naGRmc15eXjEyMjk0$&u=1980293",
		"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-m7mOiupOWktOWDjyAxXzI0MHB4LnBuZ15eXnRhbmdoZGZzXl5eMGZmZmMwZjZmODg1YTZhMWE3ZGVkZTRjMzcyYTY1MmNeXl50YW5naGRmc15eXjEzNDM0$&u=1980293",
		"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-m7mOiupOWktOWDjyAyXzI0MHB4LnBuZ15eXnRhbmdoZGZzXl5eZjk5M2Y1ZTAwNmI5ODM2YTQyOWY4MzA1ZTdmZDM2NTReXl50YW5naGRmc15eXjk1ODg$&u=1980293",
		"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-m7mOiupOWktOWDjyAzXzI0MHB4LnBuZ15eXnRhbmdoZGZzXl5eNDQ4MjI0MDYzZTlkYjNiNzUwYjM1ZjAzMjRmZDU3MWFeXl50YW5naGRmc15eXjEyMDE5$&u=1980293",
		"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-m7mOiupOWktOWDjyA0XzI0MHB4LnBuZ15eXnRhbmdoZGZzXl5eOWZjYjZlMmMzY2UxYTYyZWMzMDI1Y2YxYTRjNTFjNjReXl50YW5naGRmc15eXjkxODI$&u=1980293",
		"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-m7mOiupOWktOWDjyA1XzI0MHB4LnBuZ15eXnRhbmdoZGZzXl5eNDgwMzU5ZTZjNTg0MjdkZTczYWViYjA5MTk1N2M3ZjJeXl50YW5naGRmc15eXjEyOTQ4$&u=1980293",
		"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-m7mOiupOWktOWDjyA2XzI0MHB4LnBuZ15eXnRhbmdoZGZzXl5eMzYxNGVhZDgzYzM5N2ZiNTg0YWY0MTIyM2IzNjMyZDVeXl50YW5naGRmc15eXjk0Nzk$&u=1980293",
		"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-m7mOiupOWktOWDjyA3XzI0MHB4LnBuZ15eXnRhbmdoZGZzXl5eOGYyYzI1NWYyZGRlY2FhYTQ4NzU3ZTgyNWYxOGUwN2NeXl50YW5naGRmc15eXjEyMTg3$&u=1980293",
		"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-m7mOiupOWktOWDjyA4XzI0MHB4LnBuZ15eXnRhbmdoZGZzXl5eMjZlYmZmNjEwOTA0YjA4MTQ4Yzk1Mjk4NGJiMWM5NTBeXl50YW5naGRmc15eXjExNjA4$&u=1980293",
		"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-m7mOiupOWktOWDjyA5XzI0MHB4LnBuZ15eXnRhbmdoZGZzXl5eODk3NjE5NzA4MjI3MjIyNzg2YTI4NjQ4YWNiOGFmZGFeXl50YW5naGRmc15eXjExMzE2$&u=1980293",
		"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-m7mOiupOWktOWDjyAxMF8yNDBweC5wbmdeXl50YW5naGRmc15eXmZhNTUyMWU4ZTFmOTMzZjcyODgwMmJlNzg1NjkzNzdhXl5edGFuZ2hkZnNeXl4xMzI5MQ$&u=1980293",
		"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-m7mOiupOWktOWDjyAxMV8yNDBweC5wbmdeXl50YW5naGRmc15eXjBmYmMwNjgwNzJkOTUzZGY2ZGY3ZDJkN2E2YWUxNjBmXl5edGFuZ2hkZnNeXl45MTU4$&u=1980293",
		"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-m7mOiupOWktOWDjyAxMl8yNDBweC5wbmdeXl50YW5naGRmc15eXjY3Mzk1MTE1YWMyYWYzY2YzOTYyNjA0M2UwYjVjZmU4Xl5edGFuZ2hkZnNeXl4xMDAwMQ$&u=1980293",
		"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-m7mOiupOWktOWDjyAxM18yNDBweC5wbmdeXl50YW5naGRmc15eXjhjNTlhOWY5MmY2ZmQ1ZWExMDNkZmQ3MmM4YjI3MTlkXl5edGFuZ2hkZnNeXl4xMjM5Ng$&u=1980293",
		"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-m7mOiupOWktOWDjyAxNF8yNDBweC5wbmdeXl50YW5naGRmc15eXjgxZjI5OTVjMGIxMzc2YzUxZjk5Nzc3ZjNkMWFkMTM3Xl5edGFuZ2hkZnNeXl4xMTM4NA$&u=1980293",
	}

	BbsThumbData = BbsThumb{
		TestCloud: map[string][]string{
			"推荐": []string{
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_mjqjojZAgMC5wbmdeXl50YW5naGRmc15eXmRiNGY5YmQ4NjVhNmZhNDZhMjI0ODRkYWM3MmM2MjRjXl5edGFuZ2hkZnNeXl43MTExNDg$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_mjqjojZAgMS5wbmdeXl50YW5naGRmc15eXmNmYTlhODk2YjY0MDc3OGQ2ZDBjMWUxYjM4YzM5ZjFjXl5edGFuZ2hkZnNeXl45MjI3OTY$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_mjqjojZAgMi5wbmdeXl50YW5naGRmc15eXmIwNGQyMTIyYjVjOTFiMzkzZTQzNTcwZjUxNDU4ZjhlXl5edGFuZ2hkZnNeXl43ODY0MTA$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_mjqjojZAgMy5wbmdeXl50YW5naGRmc15eXjE2OGEwZjllMjM4Yjk4ODgxYmVhZjFiYmQ4YWMxMGJmXl5edGFuZ2hkZnNeXl44NDAzNjA$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_mjqjojZAgNC5wbmdeXl50YW5naGRmc15eXmIzMTE3YTUwMTIxMGZiOTY3YzIyMGFhZGMyMjFjYjBmXl5edGFuZ2hkZnNeXl4zNjI4ODM$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_mjqjojZAgNS5wbmdeXl50YW5naGRmc15eXjNlNzA2ZDI5NzQ0MTAxOTBjZDJkNzlkNTVlNTEyMWYxXl5edGFuZ2hkZnNeXl42ODczMjg$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_mjqjojZAgNi5wbmdeXl50YW5naGRmc15eXmIwMjlhMGYxNjJjMDc3NmVhNGUzNmRjODFhMjk0ZGQ4Xl5edGFuZ2hkZnNeXl44ODQ1Mzg$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_mjqjojZAgNy5wbmdeXl50YW5naGRmc15eXjQ4NDIxYWE3NTFkM2JkMzFkNDc3MzMxYjMyZmIyOTVjXl5edGFuZ2hkZnNeXl42OTE1OTk$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_mjqjojZAgOC5wbmdeXl50YW5naGRmc15eXjY4YjkyOTJkZmE3ZDFkYWQ0ODhlMzJmNTkyNjMxMzg1Xl5edGFuZ2hkZnNeXl41OTM4OTk$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_mjqjojZAgOS5wbmdeXl50YW5naGRmc15eXjliMmQwZTcwYTA2NjUyZGIwMjFjMjM0YTc0ZjM2ZTMxXl5edGFuZ2hkZnNeXl44NzUzODM$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_mjqjojZAgMTAucG5nXl5edGFuZ2hkZnNeXl41YWUyM2VlZTExMGI1YzViYzE1NTU4OTY1NjcxNjNjY15eXnRhbmdoZGZzXl5eNjk0Njgz$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_mjqjojZAgMTEucG5nXl5edGFuZ2hkZnNeXl44MmY2OTZjNTRlM2YzZmFkNTBmM2EwOWEzMTQ0OTVlOV5eXnRhbmdoZGZzXl5eMzEwODcw$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_mjqjojZAgMTIucG5nXl5edGFuZ2hkZnNeXl5jNTY5NDQzYmViNmExNzFlZDRlYjI4ZDM0NWNiMDZmN15eXnRhbmdoZGZzXl5eMzQ5NzA0$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_mjqjojZAgMTMucG5nXl5edGFuZ2hkZnNeXl45OTk4ODEwYjkyMzU3NGI4YmU2ZmZjOGZiOGMwNDZiOV5eXnRhbmdoZGZzXl5eNjA1ODA0$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_mjqjojZAgMTQucG5nXl5edGFuZ2hkZnNeXl43ZTE0YzMyYWE1MzZhZmRiNGZkMGJmMTQ2NTA1MTlmNF5eXnRhbmdoZGZzXl5eNzUzNzk1$&u=62051318",
			},
			"商务": []string{
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_llYbliqEgMC5wbmdeXl50YW5naGRmc15eXmQ1N2IyMmNmZmNkMjMzYjAzMjA2YWMxMDE3YWE0NDBlXl5edGFuZ2hkZnNeXl42OTUzMzg$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_llYbliqEgMS5wbmdeXl50YW5naGRmc15eXmQzZmEyYTU4YWJjNGRjN2Y4ZDAzNTc2NTFhOWM1MzQ5Xl5edGFuZ2hkZnNeXl43NDc1NzE$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_llYbliqEgMi5wbmdeXl50YW5naGRmc15eXjEwYjhjZjNkNzU1N2QyMzZmZDA3ZTliMDIzZWU4YTQxXl5edGFuZ2hkZnNeXl42ODI5MTk$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_llYbliqEgMy5wbmdeXl50YW5naGRmc15eXmFiYjM1OWZhOWI1M2U0MTJiMmQ0NDdjY2MyMDc2OWNhXl5edGFuZ2hkZnNeXl40MDk0MzA$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_llYbliqEgNC5wbmdeXl50YW5naGRmc15eXmFiNDVmZTNiOTk0YjI1NjZkYzdlMzgyMjFkNDE5N2RjXl5edGFuZ2hkZnNeXl43NDI0MDc$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_llYbliqEgNS5wbmdeXl50YW5naGRmc15eXmU0N2VlNDNmMGUxZThjNTk4MmM3NGU2NDNjOGFlOTZlXl5edGFuZ2hkZnNeXl42NjQxNzQ$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_llYbliqEgNi5wbmdeXl50YW5naGRmc15eXjUyYjNiOTE0MTc0NWIxMDBmNmM5NTNlYTIyMzI2NmRkXl5edGFuZ2hkZnNeXl42NDIzOTE$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_llYbliqEgNy5wbmdeXl50YW5naGRmc15eXjZmYTc3OTMyNzc0NzI0NDJkMzU5MWZjZDNhYTJmNjhkXl5edGFuZ2hkZnNeXl4yNjkxOTM$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_llYbliqEgOC5wbmdeXl50YW5naGRmc15eXmI1ZDBiYjM0NWYwMGY3YTAwY2I5NjNmMWU0YjRiN2ExXl5edGFuZ2hkZnNeXl4xODAwMDQ$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_llYbliqEgOS5wbmdeXl50YW5naGRmc15eXmE3OGE4Njk5YTgwMzAyYzk1YmQ1OWM3Y2RjODhiYTExXl5edGFuZ2hkZnNeXl45NDM0OQ$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_llYbliqEgMTAucG5nXl5edGFuZ2hkZnNeXl5mYTZlZWM1ZmVkMWQ2YjA0ZjY2ZDhlM2ExM2UxODc4ZV5eXnRhbmdoZGZzXl5eMTMzOTM1$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_llYbliqEgMTEucG5nXl5edGFuZ2hkZnNeXl5iNjQ4OGY5OTg5YzlhZmEwOGVjMjY2OWE4YzI4NWI2MF5eXnRhbmdoZGZzXl5eODg1NzQ$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_llYbliqEgMTIucG5nXl5edGFuZ2hkZnNeXl41YmY0Y2ExZWRkYTU5ZjY3NGQyZDY0MjNkMDI5ZjAzYV5eXnRhbmdoZGZzXl5eNzk3NTg0$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_llYbliqEgMTMucG5nXl5edGFuZ2hkZnNeXl4zMjlmYmE4ZjhlN2Y4ZTk4Y2UzZmQ5OTM3MTBkYjJkY15eXnRhbmdoZGZzXl5eNzI3MTU4$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_llYbliqEgMTQucG5nXl5edGFuZ2hkZnNeXl4xZDA0ZmI0YTI5MjM2MTNiN2U2ZmYwM2I0MjQ3MTczYV5eXnRhbmdoZGZzXl5eNzAzODI1$&u=62051318",
			},
			"办公": []string{
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lip7lhawgMC5wbmdeXl50YW5naGRmc15eXjE5NDVlYjUwM2I2MmU5ZDcxOTRjZTlmNTExNzIzZDI2Xl5edGFuZ2hkZnNeXl45MjgwOTA$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lip7lhawgMS5wbmdeXl50YW5naGRmc15eXmZmY2E4NzJlNzZkNDhlYTNjMTk5ZTQ2ZDQ0MWM2ODJkXl5edGFuZ2hkZnNeXl42MzIxMDA$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lip7lhawgMi5wbmdeXl50YW5naGRmc15eXjc0ZWNhMmMyYzU4MDc0NThmY2ZmODk4ODk3MTgyY2U2Xl5edGFuZ2hkZnNeXl42MTU4Njk$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lip7lhawgMy5wbmdeXl50YW5naGRmc15eXmQ0NmJmMWIwOTBlNTdhYWNjMGRlYWYyZDE5OTdmNTk3Xl5edGFuZ2hkZnNeXl43Mjk3MTY$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lip7lhawgNC5wbmdeXl50YW5naGRmc15eXjYxMzBjMzA0NjNlYjUxNTA3Mjc4NGQwN2RlNDljMGIxXl5edGFuZ2hkZnNeXl43Nzc3Mjk$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lip7lhawgNS5wbmdeXl50YW5naGRmc15eXmIyZjc5MWVhZmUzYzUyYTE2NDRiYmRlZmU2ZTkzMmYxXl5edGFuZ2hkZnNeXl4zNzUxMzA$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lip7lhawgNi5wbmdeXl50YW5naGRmc15eXmMzYTJhMjI2Y2FkMWZmZDkzYjAwYWI5MGMyYTdkOGY2Xl5edGFuZ2hkZnNeXl42NjM4NjU$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lip7lhawgNy5wbmdeXl50YW5naGRmc15eXjljMzA1YTNiZjg3Mjg5ZDZhYjdjMzc5OGViNWI2Zjg0Xl5edGFuZ2hkZnNeXl43NDYxNjA$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lip7lhawgOC5wbmdeXl50YW5naGRmc15eXjRjYzUwNDc4MjkwMjJiYjVhYmM4NGQxMjlmNmJlNTlmXl5edGFuZ2hkZnNeXl43ODQ5Njk$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lip7lhawgOS5wbmdeXl50YW5naGRmc15eXjM3ODg3MWIwMmJmMmNhYTJmOTNkYTM1Mjg0ZDFmMzM3Xl5edGFuZ2hkZnNeXl44Njk3OTQ$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lip7lhawgMTAucG5nXl5edGFuZ2hkZnNeXl41YTJlNWEyMzdjZjc1OGJkNjA5MmViYTNjNDk5NWZiM15eXnRhbmdoZGZzXl5eMzI0MDk3$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lip7lhawgMTEucG5nXl5edGFuZ2hkZnNeXl43NjhmOGIxZjJjZTZhMjMyNzU3MjA1NjAxMjllY2EyOF5eXnRhbmdoZGZzXl5eNzIyODE4$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lip7lhawgMTIucG5nXl5edGFuZ2hkZnNeXl5iNmI3ZDJlNjk5NTE3MDQ1NDdhZDE3MTViMmYzYWY5OV5eXnRhbmdoZGZzXl5eNjA0MDIz$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lip7lhawgMTMucG5nXl5edGFuZ2hkZnNeXl5lMDNmNjhiYzMxYWQ0ZjUwZTllYWU5NzIzNDJhYTYzNl5eXnRhbmdoZGZzXl5eODYwNjYy$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lip7lhawgMTQucG5nXl5edGFuZ2hkZnNeXl5iNWI2OGE3YTQwZDhkYmFjNWJkOTVkZDVkNmY1ZmJiZV5eXnRhbmdoZGZzXl5eMzQ3NDQ3$&u=62051318",
			},
			"节日": []string{
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_oioLml6UgMC5wbmdeXl50YW5naGRmc15eXjA0OGI2MDNiY2IxNTlhMjgyZjY3ZjRmNzU5OTFkZDdlXl5edGFuZ2hkZnNeXl45MTMwNzc$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_oioLml6UgMS5wbmdeXl50YW5naGRmc15eXmIwNGY0Mjc5NTkxZGUwNDg2ZWJiMzdhN2Q2ZDlhMWZhXl5edGFuZ2hkZnNeXl42NjgzNDM$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_oioLml6UgMi5wbmdeXl50YW5naGRmc15eXjg4ZjM0YjRmMmNiM2JlYTFlNjNjODZmYWNjZWVmZmMyXl5edGFuZ2hkZnNeXl42OTk3OTg$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_oioLml6UgMy5wbmdeXl50YW5naGRmc15eXmM3YzIzZmEwY2NhNTM3ZWI0ZjA4NTZmNjdhNWYwM2MxXl5edGFuZ2hkZnNeXl43NjM5MzU$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_oioLml6UgNC5wbmdeXl50YW5naGRmc15eXjYzNWIyNzg0OWEyZjM2YmUyZjI5YmM0YzM0MWUxNWYyXl5edGFuZ2hkZnNeXl44NDAwMTk$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_oioLml6UgNS5wbmdeXl50YW5naGRmc15eXjVlZDMwYmI4NWRmODcyOGM0MjZkY2ExYjAzOTNjZDZmXl5edGFuZ2hkZnNeXl42MDIxNjU$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_oioLml6UgNi5wbmdeXl50YW5naGRmc15eXjU1NWY3NWM4MTdkZWM5ZDIyYzliMGRhODYxZGMzZWYxXl5edGFuZ2hkZnNeXl41Mzg3NjA$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_oioLml6UgNy5wbmdeXl50YW5naGRmc15eXjAxM2U2NGZjZjYxMWM4ZGQwZjNiN2E4MDZmMDM4YzI1Xl5edGFuZ2hkZnNeXl42NDE3ODM$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_oioLml6UgOC5wbmdeXl50YW5naGRmc15eXmI0YmEyMTNkZjVjMDc2MDZhZTk0YjdiOGJhZjJhMzA1Xl5edGFuZ2hkZnNeXl41Njk1NTg$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_oioLml6UgOS5wbmdeXl50YW5naGRmc15eXjc3NzZhMjVkNWZmNTdkYTlkNmQyZjJjYjk3MWU4Y2IzXl5edGFuZ2hkZnNeXl44MTAzOTY$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_oioLml6UgMTAucG5nXl5edGFuZ2hkZnNeXl4wYmM0M2VkZjJiY2YzZmI4NmRjMzJkYzkwZDljNjgwNF5eXnRhbmdoZGZzXl5eNzQ3OTgy$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_oioLml6UgMTEucG5nXl5edGFuZ2hkZnNeXl4zMDI1YzljYmQxY2Y2MzlhODNkMDE1MTlhNDBhOTZjZV5eXnRhbmdoZGZzXl5eNzY4ODU3$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_oioLml6UgMTIucG5nXl5edGFuZ2hkZnNeXl45MzcwMDA5YzFhODc3N2JkMjdjNDJjMzJiMTAxOTRhMV5eXnRhbmdoZGZzXl5eNTQwNDE2$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_oioLml6UgMTMucG5nXl5edGFuZ2hkZnNeXl4xZjcyODRkMDc0MTU2ODg2ZGZkNGRmNGU3YzgyODRmZF5eXnRhbmdoZGZzXl5eNTA4NzUy$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_oioLml6UgMTQucG5nXl5edGFuZ2hkZnNeXl41Yzc4M2ZjN2Q1MmU3NmFlNGI1OWM3M2QyZDYxMDlmZF5eXnRhbmdoZGZzXl5eODgxOTY5$&u=62051318",
			},
			"运动": []string{
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ov5DliqggMC5wbmdeXl50YW5naGRmc15eXjFiMTliOWE5YjNkM2NiMTc3YWEwZjAyYmIzMzRjN2I0Xl5edGFuZ2hkZnNeXl43NTU1OTA$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ov5DliqggMS5wbmdeXl50YW5naGRmc15eXjcxMGQ0YmRiNGJjY2M1YjE4ODY4NGYwZjU4Y2EwOTZhXl5edGFuZ2hkZnNeXl44NzI2NDc$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ov5DliqggMi5wbmdeXl50YW5naGRmc15eXjgwZmNkNzc0ZTY0N2Y2NzU5Yzk0MzAxMjVhYjA4NDJjXl5edGFuZ2hkZnNeXl44NzMzMDE$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ov5DliqggMy5wbmdeXl50YW5naGRmc15eXjYwZTVkMDRiZmY3YjQ0ZWZiY2QxZWI3NGVkZWI5Njk4Xl5edGFuZ2hkZnNeXl43MDY5ODI$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ov5DliqggNC5wbmdeXl50YW5naGRmc15eXmQ4YzUwNDliYTM1Yjk0YjA5OTE0YWY0YzQwZTg0YjMyXl5edGFuZ2hkZnNeXl41Njk0ODc$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ov5DliqggNS5wbmdeXl50YW5naGRmc15eXmZiMGY5MzYxMjk3NmY3MDVjNTNiMmQ1MzQxNmJiM2IzXl5edGFuZ2hkZnNeXl41Mjk4NzQ$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ov5DliqggNi5wbmdeXl50YW5naGRmc15eXjI3NzE5ZDE2MDBhNTRmODBkMmI5MWVjYzZjM2ZhOTI5Xl5edGFuZ2hkZnNeXl4xMDI2MzAx$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ov5DliqggNy5wbmdeXl50YW5naGRmc15eXjEwMzc4OWQ5MjMwYzdkODVkYzQ3MmE3MmJmOGVkMjczXl5edGFuZ2hkZnNeXl41NTA0MzQ$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ov5DliqggOC5wbmdeXl50YW5naGRmc15eXmU5MGM0ZWYyNTdlODc1ZjFmN2IyYmM0ZDQzZjNmMWIxXl5edGFuZ2hkZnNeXl43MDg2ODY$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ov5DliqggOS5wbmdeXl50YW5naGRmc15eXjM4Y2NlYmViZDk4ZTE0NDUxNDgyYzJjMjRhYTY3ZTE5Xl5edGFuZ2hkZnNeXl43MzI5MDA$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ov5DliqggMTAucG5nXl5edGFuZ2hkZnNeXl5iMTA4YmE5NGMwYTExOTc3MDQ0ZWY0NjY0NDIxNjNjM15eXnRhbmdoZGZzXl5eODkwMDE5$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ov5DliqggMTEucG5nXl5edGFuZ2hkZnNeXl5hZjI3MWNlZjcwZWQ3ZTVhMzNiYWJiNDQ4MTljZjMxMV5eXnRhbmdoZGZzXl5eMTEzODQ3OQ$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ov5DliqggMTIucG5nXl5edGFuZ2hkZnNeXl5jNTFjZmU5NTZjODE5NTNkNmQyYTM0MjRmNjE0ZTI5OV5eXnRhbmdoZGZzXl5eODM4NjQz$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ov5DliqggMTMucG5nXl5edGFuZ2hkZnNeXl5lOWZjODk5OTcwNzQzMDEzYmQ4YTk0NTE4NDcxMjgxMF5eXnRhbmdoZGZzXl5eMTAwMzc0NQ$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ov5DliqggMTQucG5nXl5edGFuZ2hkZnNeXl42OGM4MGUwYmE2NTI4NWI0YmE1OThmYTY4MWNmNjMzZV5eXnRhbmdoZGZzXl5eNzc4NTUy$&u=62051318",
			},
			"美食": []string{
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_nvo7po58gMC5wbmdeXl50YW5naGRmc15eXjhiYjdiYmVmY2I5OTYyN2QwNGJhZDVjNGY4Y2RhOTNmXl5edGFuZ2hkZnNeXl44MzEzNzE$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_nvo7po58gMS5wbmdeXl50YW5naGRmc15eXmQxODhlMDkxNDdhOTZjYWQ0MjJkOTZmNmIwNmYxOTRkXl5edGFuZ2hkZnNeXl43NDIyMzc$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_nvo7po58gMi5wbmdeXl50YW5naGRmc15eXjYxNzRlMjBmMzZlMDQ5NGU3NWVkYjRmZjIyODVmMDcyXl5edGFuZ2hkZnNeXl43ODE3ODA$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_nvo7po58gMy5wbmdeXl50YW5naGRmc15eXmNhODg2YjFiODE0Mjk5ZDY4ZjhiNmJhYTFkNDc2NTVmXl5edGFuZ2hkZnNeXl42MDU2MDA$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_nvo7po58gNC5wbmdeXl50YW5naGRmc15eXjg2MTBkN2I3ZDAzMjYyYTEwMTI1ZmI3ZjY4OTk2MTcxXl5edGFuZ2hkZnNeXl44MjkyNDg$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_nvo7po58gNS5wbmdeXl50YW5naGRmc15eXjA0YzI0ZTQ5MTIxYzRmZmMxMzIxMWNiNjA0ZmZhMjA1Xl5edGFuZ2hkZnNeXl42MDIyMzg$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_nvo7po58gNi5wbmdeXl50YW5naGRmc15eXjI4N2RkNzgyMTFhOWJkMmI0YjQ5MDc0ZjA3MDE5N2E1Xl5edGFuZ2hkZnNeXl42OTM0MDE$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_nvo7po58gNy5wbmdeXl50YW5naGRmc15eXmE5NTU2NzYzMzQ5MjA1NWZiZGZjYWQ4NDdlNDM1YzdiXl5edGFuZ2hkZnNeXl43ODQ1NDY$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_nvo7po58gOC5wbmdeXl50YW5naGRmc15eXmRlODQ0YjNkODM5NjkzN2EyM2FjNmM5MTlkZmJmMWJhXl5edGFuZ2hkZnNeXl4xMDE5MDE1$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_nvo7po58gOS5wbmdeXl50YW5naGRmc15eXjk3ZGRjY2Q2YzdjOTcxNTczYmFlNjNiZmI1NGYyYTA2Xl5edGFuZ2hkZnNeXl43NzMzNTU$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_nvo7po58gMTAucG5nXl5edGFuZ2hkZnNeXl41ZjA3OWY1MjllNzQ5ZmUzZGM5NGQyOGI1Y2MyNzY3YV5eXnRhbmdoZGZzXl5eNTk3MjM5$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_nvo7po58gMTEucG5nXl5edGFuZ2hkZnNeXl40YmVkODBkYWM0MTNjZmE4MTE0YjlkMzYyOWNlOWQ1OF5eXnRhbmdoZGZzXl5eOTM2Mzg5$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_nvo7po58gMTIucG5nXl5edGFuZ2hkZnNeXl5kNjFiMzUwY2QyMGFhNmQ1ODc0YjBhNjU2ZWRlN2E2MV5eXnRhbmdoZGZzXl5eODI2MzY2$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_nvo7po58gMTMucG5nXl5edGFuZ2hkZnNeXl43YWY3ZDdhZWM3YWVmZTMzYjkyMzBkN2YwNDE2MDI2YV5eXnRhbmdoZGZzXl5eMTA3ODY0NQ$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_nvo7po58gMTQucG5nXl5edGFuZ2hkZnNeXl42MWM5Mzg3ZGFhMTljZDg4NmFkNGRmMzdkNmI5YjU5OF5eXnRhbmdoZGZzXl5eNzg2OTQ4$&u=62051318",
			},
			"城市": []string{
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ln47luIIgMC5wbmdeXl50YW5naGRmc15eXjcxZTljZjE1N2NiMTQxY2QxZGI1YTM0OTQ2MjBhNDQzXl5edGFuZ2hkZnNeXl43MDEzNzE$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ln47luIIgMS5wbmdeXl50YW5naGRmc15eXmY4ZTQ1N2YyOTVjNmZkMWQxMzAxN2ZiZjM3NzU5NzE5Xl5edGFuZ2hkZnNeXl43OTY3MTg$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ln47luIIgMi5wbmdeXl50YW5naGRmc15eXjcyMjUyYTdlZWI4NmFjOWU5YjNjMzgwYTUyYWY3Y2JmXl5edGFuZ2hkZnNeXl44MzY5MjI$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ln47luIIgMy5wbmdeXl50YW5naGRmc15eXjM4NTJiNjcwNDBjZDk0ODY2YjcwNTE4YTdhOTBhZTNkXl5edGFuZ2hkZnNeXl44NzI1MTc$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ln47luIIgNC5wbmdeXl50YW5naGRmc15eXjE4NzNjNjA0MDRlYzAyZDc0ZmZkYmM1NDhlNzc4OTAzXl5edGFuZ2hkZnNeXl43MzM4MDI$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ln47luIIgNS5wbmdeXl50YW5naGRmc15eXjI3NzNkYTRhNDY5ZDUwNDY3ZjYxYmY1ZjgyNjVjOGI2Xl5edGFuZ2hkZnNeXl45NzIxNDE$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ln47luIIgNi5wbmdeXl50YW5naGRmc15eXjgxZDllMWMxMTUxOTk3MWJjMDRmNDMyM2MyODIwNWJhXl5edGFuZ2hkZnNeXl43OTQyMzQ$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ln47luIIgNy5wbmdeXl50YW5naGRmc15eXmRmMGQ1MWEzZmQwYjkzMTA0ZGU1Y2Y1OWMwZDQ3NjliXl5edGFuZ2hkZnNeXl43NzkzMjk$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ln47luIIgOC5wbmdeXl50YW5naGRmc15eXjQ0N2Y2NzE3NzI2MTNhZjQ3ZjFiODM4NGU5Y2E3N2VlXl5edGFuZ2hkZnNeXl44OTk3MTU$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ln47luIIgOS5wbmdeXl50YW5naGRmc15eXjM5YTg1ZjQ2NDkwNTdmNDA1MDVhOGVjNDY0NWQyY2U2Xl5edGFuZ2hkZnNeXl43MTk4NTM$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ln47luIIgMTAucG5nXl5edGFuZ2hkZnNeXl5kZDBmM2M5MDIyMDdiYmViNWJlYTQ3Y2JmNWY5ZmZmY15eXnRhbmdoZGZzXl5eODYxNzUw$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ln47luIIgMTEucG5nXl5edGFuZ2hkZnNeXl5jY2U1ZGNhMzM3NmZkZjY5MWFiNDUzODhhMmQxMWExMV5eXnRhbmdoZGZzXl5eODMwMjc3$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ln47luIIgMTIucG5nXl5edGFuZ2hkZnNeXl40ZDVjNTljZmVjMmNjZGYyMWZjZmY5NWFlYzA3ZTgwMF5eXnRhbmdoZGZzXl5eOTI0Mzcx$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ln47luIIgMTMucG5nXl5edGFuZ2hkZnNeXl43NTMxYzc2NjdmYzY0ZWE4NzRkNDNkZDBkODdmYjU1ZF5eXnRhbmdoZGZzXl5eOTYwNDQy$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_ln47luIIgMTQucG5nXl5edGFuZ2hkZnNeXl5hOWM5Zjk5ZDk2ODk3ZWYzODA4YTc4MTk1YWQ2YWMzY15eXnRhbmdoZGZzXl5eNzAxOTc4$&u=62051318",
			},
			"风景": []string{
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_po47mma8gMC5wbmdeXl50YW5naGRmc15eXjA3OGQwZGQzYjllYzZhZWE4ZmMyNjVjMDJlNzVmMWZmXl5edGFuZ2hkZnNeXl43NTk1NzU$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_po47mma8gMS5wbmdeXl50YW5naGRmc15eXmE0MGRjZjAxOTUwNDM2Mzk3YjQ5ODVhODc4MTVmM2YyXl5edGFuZ2hkZnNeXl44MTIxMDQ$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_po47mma8gMi5wbmdeXl50YW5naGRmc15eXmQyY2E1NzEyZDYzMzJkZWNkZTAyOTNkZjUzM2MxMDIyXl5edGFuZ2hkZnNeXl42ODM3ODY$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_po47mma8gMy5wbmdeXl50YW5naGRmc15eXjcyODg4ZWUxY2QzN2QyODgyZWI1ZTBlNDdlN2U4ZjBiXl5edGFuZ2hkZnNeXl45MTI1NTA$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_po47mma8gNC5wbmdeXl50YW5naGRmc15eXmQxMzQ0ZDBhODE0MTNmOWFjNmUxYWU4ODc1NjM3MzUxXl5edGFuZ2hkZnNeXl43OTIzMTE$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_po47mma8gNS5wbmdeXl50YW5naGRmc15eXjU4ZWQwMWIwOWYxYjQ2MGNkYzA1MzgwNzBmYmIwOGEzXl5edGFuZ2hkZnNeXl42MzQ3NzM$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_po47mma8gNi5wbmdeXl50YW5naGRmc15eXmRiMGExN2M2ODgxODA5OWE3ZmNlZGJkZjJlYWUyYjVjXl5edGFuZ2hkZnNeXl42OTMyMTA$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_po47mma8gNy5wbmdeXl50YW5naGRmc15eXjUxNmZkZDAwY2NlNzVmODgxZTYzMGE3NWJkZWJmM2Y5Xl5edGFuZ2hkZnNeXl43NzQ1MDk$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_po47mma8gOC5wbmdeXl50YW5naGRmc15eXmRkNzQ2YjRmOTRhNzM0MTQ1ZDUwM2ExYzg1MzU4NmUzXl5edGFuZ2hkZnNeXl42OTczNTE$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_po47mma8gOS5wbmdeXl50YW5naGRmc15eXjAxMDRjNWFlZjZkY2U2MmE2NGUxN2M1MWU3YmU3MTYzXl5edGFuZ2hkZnNeXl43MDEwNzY$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_po47mma8gMTAucG5nXl5edGFuZ2hkZnNeXl44NTY2OTg3MWJiYmY4ZWFhZTExZDA1OWMxNTcwNjg0M15eXnRhbmdoZGZzXl5eNjM2MzA1$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_po47mma8gMTEucG5nXl5edGFuZ2hkZnNeXl4zN2U0MTk1ZmQ2ZTFmNGEyNTMzYzUzZDQ1N2ZlMGFiZV5eXnRhbmdoZGZzXl5eNzMxNzEz$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_po47mma8gMTIucG5nXl5edGFuZ2hkZnNeXl5jOWU5MDkwZjZhZjNhYjE5ZDdkYzg3YWJjOWRhYjc3Zl5eXnRhbmdoZGZzXl5eMTA3MjQwNQ$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_po47mma8gMTMucG5nXl5edGFuZ2hkZnNeXl40ODQ2YTFlYjhiZGZiNWU3MjRkNTA3MGYyODQ1MzMxOV5eXnRhbmdoZGZzXl5eNTg5NTIw$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_po47mma8gMTQucG5nXl5edGFuZ2hkZnNeXl5jZWQxOWU1NGRkNGRhOGRlN2M2NjkwMzJjODQ4MDlhYl5eXnRhbmdoZGZzXl5eNTkwNTcy$&u=62051318",
			},
			"季节": []string{
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lraPoioIgMC5wbmdeXl50YW5naGRmc15eXjczN2M4NmJmMjA3YmYyZjdlODljNzkzYmVlNWE3ZDk3Xl5edGFuZ2hkZnNeXl43NDgwNTU$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lraPoioIgMS5wbmdeXl50YW5naGRmc15eXmMyZGZkOTNmYWE0OTYxYjRiMjhiNGFkMzdlYmI2ZDNmXl5edGFuZ2hkZnNeXl41MjY2NDQ$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lraPoioIgMi5wbmdeXl50YW5naGRmc15eXjNlYTI3MmE0MmE4ZDFiYTJlMjRmZmJlZThiYzI1NDBiXl5edGFuZ2hkZnNeXl42MzI1OTU$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lraPoioIgMy5wbmdeXl50YW5naGRmc15eXmZhZDYxNWY5NzUyZDRmODE3Zjc4ZDI2MTBhZTZmYWE3Xl5edGFuZ2hkZnNeXl45Njk4MTU$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lraPoioIgNC5wbmdeXl50YW5naGRmc15eXjgwYjc4YzQzMmE2NzA4ZGY5MDRmNzI5M2VhNjg1NzExXl5edGFuZ2hkZnNeXl43ODQ0MTg$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lraPoioIgNS5wbmdeXl50YW5naGRmc15eXmVmYzc0NjVhMGZiYzE1MGEwNWNhMTQ2YjE1MDg0NmZjXl5edGFuZ2hkZnNeXl44MzkwNTE$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lraPoioIgNi5wbmdeXl50YW5naGRmc15eXmEzOTQ1NWQ3YTExMjY3OWYxZGMyOGM3NjllNjZmOTQzXl5edGFuZ2hkZnNeXl43OTczNTc$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lraPoioIgNy5wbmdeXl50YW5naGRmc15eXjU3ZTliNWQ2ZWQzM2UxODJkYmEzMzRlN2EyZWZhNjIxXl5edGFuZ2hkZnNeXl43MjA1MzM$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lraPoioIgOC5wbmdeXl50YW5naGRmc15eXjhiZTUwN2U2NzNiMGRkNmM1MWQ4OGYxOTFhNzA2NWUzXl5edGFuZ2hkZnNeXl43NDc4OTk$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lraPoioIgOS5wbmdeXl50YW5naGRmc15eXjM0YzdiYzhiYjc5M2I2OWQ5NjIyYmIyNWViMzMzNWNjXl5edGFuZ2hkZnNeXl45OTg0Mzk$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lraPoioIgMTAucG5nXl5edGFuZ2hkZnNeXl4zNDZhZTkwNDhiNjBkZmZkNzhhNzY5ZjE5ZGU2NzY1ZV5eXnRhbmdoZGZzXl5eODA3Njc3$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lraPoioIgMTEucG5nXl5edGFuZ2hkZnNeXl43ZDRkNzc3MTRlNDQzNGZiMGQ1ZTZmZTczYjEwNGQ2Nl5eXnRhbmdoZGZzXl5eMzY1ODYy$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lraPoioIgMTIucG5nXl5edGFuZ2hkZnNeXl5hZjJjYTYwYTlmMmMxODk0YTIwNjNiMzY5NDE1NTE4NV5eXnRhbmdoZGZzXl5eOTE0MTMx$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lraPoioIgMTMucG5nXl5edGFuZ2hkZnNeXl45ZWM4MTNhZmExN2I4ZWUyNjc3MzA1NDcwNzg1Mjg3M15eXnRhbmdoZGZzXl5eOTIwNjg4$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lraPoioIgMTQucG5nXl5edGFuZ2hkZnNeXl43NjYzMWQzY2UxNTk2NzE2NDM2ZjY1OTg2NDU5YTJlMV5eXnRhbmdoZGZzXl5eODk3MjM4$&u=62051318",
			},
			"健康": []string{
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lgaXlurcgMC5wbmdeXl50YW5naGRmc15eXjNkOWVmNjRiNTZiZGFhYjdkZDEyODc0NzJlNjBlNDA3Xl5edGFuZ2hkZnNeXl41NTM2Njk$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lgaXlurcgMS5wbmdeXl50YW5naGRmc15eXmNlZDY0YWIyYjZkYjI5Y2IyYjk4NWVjZDcyZGMxMmQ4Xl5edGFuZ2hkZnNeXl43MzgwNTE$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lgaXlurcgMi5wbmdeXl50YW5naGRmc15eXjNmZTU5NjI3MjA3YmEwOGRjNThkYTI5NWIxMzhlZjJjXl5edGFuZ2hkZnNeXl41OTAyMjI$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lgaXlurcgMy5wbmdeXl50YW5naGRmc15eXjBhZmE2MTk4NzIxZDFkYzkxYmY1MWM2MTZmYjJkYjUxXl5edGFuZ2hkZnNeXl43NzY4NzU$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lgaXlurcgNC5wbmdeXl50YW5naGRmc15eXjlkNmI3MzAxMzJiNDdlMTYzNmY2OTE0NDU2M2U2ZTEyXl5edGFuZ2hkZnNeXl43MTIyNzE$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lgaXlurcgNS5wbmdeXl50YW5naGRmc15eXjY3ZGQzMmNmNzgxZGY1OGNlZTcwNzBkN2NhYzJhZTkyXl5edGFuZ2hkZnNeXl42MzU4NjI$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lgaXlurcgNi5wbmdeXl50YW5naGRmc15eXjc5ZWQyMjU5NDg1YzU1YzBlYTZmODgwNjNiZmNhOWMxXl5edGFuZ2hkZnNeXl43Mzc2NjE$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lgaXlurcgNy5wbmdeXl50YW5naGRmc15eXjA2NzdhZjM4NDIxOGRmMDg1N2NmNzkzZmRiMmUyNzdlXl5edGFuZ2hkZnNeXl45MzgwMDQ$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lgaXlurcgOC5wbmdeXl50YW5naGRmc15eXjNhYmRmNDhiNWQ4MzA3YmZjZjMzNDVmYWIxMTc0YWVkXl5edGFuZ2hkZnNeXl42MTAxMjg$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lgaXlurcgOS5wbmdeXl50YW5naGRmc15eXjE4ODg5M2M4ZjlhMmEyMTg2YzFkZjlmOTIxN2FlMjViXl5edGFuZ2hkZnNeXl44MzQwOTU$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lgaXlurcgMTAucG5nXl5edGFuZ2hkZnNeXl4wYmM0M2VkZjJiY2YzZmI4NmRjMzJkYzkwZDljNjgwNF5eXnRhbmdoZGZzXl5eNzQ3OTgy$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lgaXlurcgMTEucG5nXl5edGFuZ2hkZnNeXl40Y2Q2NTBmNTM0MGFkYzc0YzcyN2VlNzhjOTBhZWE1OV5eXnRhbmdoZGZzXl5eNjA0NjQ3$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lgaXlurcgMTIucG5nXl5edGFuZ2hkZnNeXl4xOTdjY2NkMjg2NDVhNTQ0ZTkwMmEyMDA2OWNhYjEwNF5eXnRhbmdoZGZzXl5eNjA2ODQx$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lgaXlurcgMTMucG5nXl5edGFuZ2hkZnNeXl41NjgzNDAzM2Y0YWMxYTZlMWJiYjRjMzViZmJhYjk2Y15eXnRhbmdoZGZzXl5eNDQ0Mzk5$&u=62051318",
				"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_lgaXlurcgMTQucG5nXl5edGFuZ2hkZnNeXl5jYjk0MTg5YTIzMjEyOWI3ZDQ3ZWFkNmU4M2ExOGY5OV5eXnRhbmdoZGZzXl5eNDEwMTMz$&u=62051318",
			},
		},
		OnCloud: map[string][]string{
			"推荐": []string{
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-aOqOiNkCAwLnBuZ15eXnRhbmdoZGZzXl5eZGI0ZjliZDg2NWE2ZmE0NmEyMjQ4NGRhYzcyYzYyNGNeXl50YW5naGRmc15eXjcxMTE0OA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-aOqOiNkCAxLnBuZ15eXnRhbmdoZGZzXl5eY2ZhOWE4OTZiNjQwNzc4ZDZkMGMxZTFiMzhjMzlmMWNeXl50YW5naGRmc15eXjkyMjc5Ng$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-aOqOiNkCAyLnBuZ15eXnRhbmdoZGZzXl5eYjA0ZDIxMjJiNWM5MWIzOTNlNDM1NzBmNTE0NThmOGVeXl50YW5naGRmc15eXjc4NjQxMA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-aOqOiNkCAzLnBuZ15eXnRhbmdoZGZzXl5eMTY4YTBmOWUyMzhiOTg4ODFiZWFmMWJiZDhhYzEwYmZeXl50YW5naGRmc15eXjg0MDM2MA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-aOqOiNkCA0LnBuZ15eXnRhbmdoZGZzXl5eYjMxMTdhNTAxMjEwZmI5NjdjMjIwYWFkYzIyMWNiMGZeXl50YW5naGRmc15eXjM2Mjg4Mw$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-aOqOiNkCA1LnBuZ15eXnRhbmdoZGZzXl5eM2U3MDZkMjk3NDQxMDE5MGNkMmQ3OWQ1NWU1MTIxZjFeXl50YW5naGRmc15eXjY4NzMyOA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-aOqOiNkCA2LnBuZ15eXnRhbmdoZGZzXl5eYjAyOWEwZjE2MmMwNzc2ZWE0ZTM2ZGM4MWEyOTRkZDheXl50YW5naGRmc15eXjg4NDUzOA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-aOqOiNkCA3LnBuZ15eXnRhbmdoZGZzXl5eNDg0MjFhYTc1MWQzYmQzMWQ0NzczMzFiMzJmYjI5NWNeXl50YW5naGRmc15eXjY5MTU5OQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-aOqOiNkCA4LnBuZ15eXnRhbmdoZGZzXl5eNjhiOTI5MmRmYTdkMWRhZDQ4OGUzMmY1OTI2MzEzODVeXl50YW5naGRmc15eXjU5Mzg5OQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-aOqOiNkCA5LnBuZ15eXnRhbmdoZGZzXl5eOWIyZDBlNzBhMDY2NTJkYjAyMWMyMzRhNzRmMzZlMzFeXl50YW5naGRmc15eXjg3NTM4Mw$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-aOqOiNkCAxMC5wbmdeXl50YW5naGRmc15eXjVhZTIzZWVlMTEwYjVjNWJjMTU1NTg5NjU2NzE2M2NjXl5edGFuZ2hkZnNeXl42OTQ2ODM$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-aOqOiNkCAxMS5wbmdeXl50YW5naGRmc15eXjgyZjY5NmM1NGUzZjNmYWQ1MGYzYTA5YTMxNDQ5NWU5Xl5edGFuZ2hkZnNeXl4zMTA4NzA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-aOqOiNkCAxMi5wbmdeXl50YW5naGRmc15eXmM1Njk0NDNiZWI2YTE3MWVkNGViMjhkMzQ1Y2IwNmY3Xl5edGFuZ2hkZnNeXl4zNDk3MDQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-aOqOiNkCAxMy5wbmdeXl50YW5naGRmc15eXjk5OTg4MTBiOTIzNTc0YjhiZTZmZmM4ZmI4YzA0NmI5Xl5edGFuZ2hkZnNeXl42MDU4MDQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-aOqOiNkCAxNC5wbmdeXl50YW5naGRmc15eXjdlMTRjMzJhYTUzNmFmZGI0ZmQwYmYxNDY1MDUxOWY0Xl5edGFuZ2hkZnNeXl43NTM3OTU$&u=1980293",
			},
			"商务": []string{
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WVhuWKoSAwLnBuZ15eXnRhbmdoZGZzXl5eZDU3YjIyY2ZmY2QyMzNiMDMyMDZhYzEwMTdhYTQ0MGVeXl50YW5naGRmc15eXjY5NTMzOA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WVhuWKoSAxLnBuZ15eXnRhbmdoZGZzXl5eZDNmYTJhNThhYmM0ZGM3ZjhkMDM1NzY1MWE5YzUzNDleXl50YW5naGRmc15eXjc0NzU3MQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WVhuWKoSAyLnBuZ15eXnRhbmdoZGZzXl5eMTBiOGNmM2Q3NTU3ZDIzNmZkMDdlOWIwMjNlZThhNDFeXl50YW5naGRmc15eXjY4MjkxOQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WVhuWKoSAzLnBuZ15eXnRhbmdoZGZzXl5eYWJiMzU5ZmE5YjUzZTQxMmIyZDQ0N2NjYzIwNzY5Y2FeXl50YW5naGRmc15eXjQwOTQzMA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WVhuWKoSA0LnBuZ15eXnRhbmdoZGZzXl5eYWI0NWZlM2I5OTRiMjU2NmRjN2UzODIyMWQ0MTk3ZGNeXl50YW5naGRmc15eXjc0MjQwNw$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WVhuWKoSA1LnBuZ15eXnRhbmdoZGZzXl5eZTQ3ZWU0M2YwZTFlOGM1OTgyYzc0ZTY0M2M4YWU5NmVeXl50YW5naGRmc15eXjY2NDE3NA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WVhuWKoSA2LnBuZ15eXnRhbmdoZGZzXl5eNTJiM2I5MTQxNzQ1YjEwMGY2Yzk1M2VhMjIzMjY2ZGReXl50YW5naGRmc15eXjY0MjM5MQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WVhuWKoSA3LnBuZ15eXnRhbmdoZGZzXl5eNmZhNzc5MzI3NzQ3MjQ0MmQzNTkxZmNkM2FhMmY2OGReXl50YW5naGRmc15eXjI2OTE5Mw$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WVhuWKoSA4LnBuZ15eXnRhbmdoZGZzXl5eYjVkMGJiMzQ1ZjAwZjdhMDBjYjk2M2YxZTRiNGI3YTFeXl50YW5naGRmc15eXjE4MDAwNA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WVhuWKoSA5LnBuZ15eXnRhbmdoZGZzXl5eYTc4YTg2OTlhODAzMDJjOTViZDU5YzdjZGM4OGJhMTFeXl50YW5naGRmc15eXjk0MzQ5$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WVhuWKoSAxMC5wbmdeXl50YW5naGRmc15eXmZhNmVlYzVmZWQxZDZiMDRmNjZkOGUzYTEzZTE4NzhlXl5edGFuZ2hkZnNeXl4xMzM5MzU$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WVhuWKoSAxMS5wbmdeXl50YW5naGRmc15eXmI2NDg4Zjk5ODljOWFmYTA4ZWMyNjY5YThjMjg1YjYwXl5edGFuZ2hkZnNeXl44ODU3NA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WVhuWKoSAxMi5wbmdeXl50YW5naGRmc15eXjViZjRjYTFlZGRhNTlmNjc0ZDJkNjQyM2QwMjlmMDNhXl5edGFuZ2hkZnNeXl43OTc1ODQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WVhuWKoSAxMy5wbmdeXl50YW5naGRmc15eXjMyOWZiYThmOGU3ZjhlOThjZTNmZDk5MzcxMGRiMmRjXl5edGFuZ2hkZnNeXl43MjcxNTg$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WVhuWKoSAxNC5wbmdeXl50YW5naGRmc15eXjFkMDRmYjRhMjkyMzYxM2I3ZTZmZjAzYjQyNDcxNzNhXl5edGFuZ2hkZnNeXl43MDM4MjU$&u=1980293",
			},
			"办公": []string{
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WKnuWFrCAwLnBuZ15eXnRhbmdoZGZzXl5eMTk0NWViNTAzYjYyZTlkNzE5NGNlOWY1MTE3MjNkMjZeXl50YW5naGRmc15eXjkyODA5MA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WKnuWFrCAxLnBuZ15eXnRhbmdoZGZzXl5eZmZjYTg3MmU3NmQ0OGVhM2MxOTllNDZkNDQxYzY4MmReXl50YW5naGRmc15eXjYzMjEwMA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WKnuWFrCAyLnBuZ15eXnRhbmdoZGZzXl5eNzRlY2EyYzJjNTgwNzQ1OGZjZmY4OTg4OTcxODJjZTZeXl50YW5naGRmc15eXjYxNTg2OQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WKnuWFrCAzLnBuZ15eXnRhbmdoZGZzXl5eZDQ2YmYxYjA5MGU1N2FhY2MwZGVhZjJkMTk5N2Y1OTdeXl50YW5naGRmc15eXjcyOTcxNg$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WKnuWFrCA0LnBuZ15eXnRhbmdoZGZzXl5eNjEzMGMzMDQ2M2ViNTE1MDcyNzg0ZDA3ZGU0OWMwYjFeXl50YW5naGRmc15eXjc3NzcyOQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WKnuWFrCA1LnBuZ15eXnRhbmdoZGZzXl5eYjJmNzkxZWFmZTNjNTJhMTY0NGJiZGVmZTZlOTMyZjFeXl50YW5naGRmc15eXjM3NTEzMA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WKnuWFrCA2LnBuZ15eXnRhbmdoZGZzXl5eYzNhMmEyMjZjYWQxZmZkOTNiMDBhYjkwYzJhN2Q4ZjZeXl50YW5naGRmc15eXjY2Mzg2NQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WKnuWFrCA3LnBuZ15eXnRhbmdoZGZzXl5eOWMzMDVhM2JmODcyODlkNmFiN2MzNzk4ZWI1YjZmODReXl50YW5naGRmc15eXjc0NjE2MA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WKnuWFrCA4LnBuZ15eXnRhbmdoZGZzXl5eNGNjNTA0NzgyOTAyMmJiNWFiYzg0ZDEyOWY2YmU1OWZeXl50YW5naGRmc15eXjc4NDk2OQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WKnuWFrCA5LnBuZ15eXnRhbmdoZGZzXl5eMzc4ODcxYjAyYmYyY2FhMmY5M2RhMzUyODRkMWYzMzdeXl50YW5naGRmc15eXjg2OTc5NA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WKnuWFrCAxMC5wbmdeXl50YW5naGRmc15eXjVhMmU1YTIzN2NmNzU4YmQ2MDkyZWJhM2M0OTk1ZmIzXl5edGFuZ2hkZnNeXl4zMjQwOTc$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WKnuWFrCAxMS5wbmdeXl50YW5naGRmc15eXjc2OGY4YjFmMmNlNmEyMzI3NTcyMDU2MDEyOWVjYTI4Xl5edGFuZ2hkZnNeXl43MjI4MTg$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WKnuWFrCAxMi5wbmdeXl50YW5naGRmc15eXmI2YjdkMmU2OTk1MTcwNDU0N2FkMTcxNWIyZjNhZjk5Xl5edGFuZ2hkZnNeXl42MDQwMjM$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WKnuWFrCAxMy5wbmdeXl50YW5naGRmc15eXmUwM2Y2OGJjMzFhZDRmNTBlOWVhZTk3MjM0MmFhNjM2Xl5edGFuZ2hkZnNeXl44NjA2NjI$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WKnuWFrCAxNC5wbmdeXl50YW5naGRmc15eXmI1YjY4YTdhNDBkOGRiYWM1YmQ5NWRkNWQ2ZjVmYmJlXl5edGFuZ2hkZnNeXl4zNDc0NDc$&u=1980293",
			},
			"节日": []string{
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-iKguaXpSAwLnBuZ15eXnRhbmdoZGZzXl5eMDQ4YjYwM2JjYjE1OWEyODJmNjdmNGY3NTk5MWRkN2VeXl50YW5naGRmc15eXjkxMzA3Nw$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-iKguaXpSAxLnBuZ15eXnRhbmdoZGZzXl5eYjA0ZjQyNzk1OTFkZTA0ODZlYmIzN2E3ZDZkOWExZmFeXl50YW5naGRmc15eXjY2ODM0Mw$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-iKguaXpSAyLnBuZ15eXnRhbmdoZGZzXl5eODhmMzRiNGYyY2IzYmVhMWU2M2M4NmZhY2NlZWZmYzJeXl50YW5naGRmc15eXjY5OTc5OA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-iKguaXpSAzLnBuZ15eXnRhbmdoZGZzXl5eYzdjMjNmYTBjY2E1MzdlYjRmMDg1NmY2N2E1ZjAzYzFeXl50YW5naGRmc15eXjc2MzkzNQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-iKguaXpSA0LnBuZ15eXnRhbmdoZGZzXl5eNjM1YjI3ODQ5YTJmMzZiZTJmMjliYzRjMzQxZTE1ZjJeXl50YW5naGRmc15eXjg0MDAxOQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-iKguaXpSA1LnBuZ15eXnRhbmdoZGZzXl5eNWVkMzBiYjg1ZGY4NzI4YzQyNmRjYTFiMDM5M2NkNmZeXl50YW5naGRmc15eXjYwMjE2NQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-iKguaXpSA2LnBuZ15eXnRhbmdoZGZzXl5eNTU1Zjc1YzgxN2RlYzlkMjJjOWIwZGE4NjFkYzNlZjFeXl50YW5naGRmc15eXjUzODc2MA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-iKguaXpSA3LnBuZ15eXnRhbmdoZGZzXl5eMDEzZTY0ZmNmNjExYzhkZDBmM2I3YTgwNmYwMzhjMjVeXl50YW5naGRmc15eXjY0MTc4Mw$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-iKguaXpSA4LnBuZ15eXnRhbmdoZGZzXl5eYjRiYTIxM2RmNWMwNzYwNmFlOTRiN2I4YmFmMmEzMDVeXl50YW5naGRmc15eXjU2OTU1OA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-iKguaXpSA5LnBuZ15eXnRhbmdoZGZzXl5eNzc3NmEyNWQ1ZmY1N2RhOWQ2ZDJmMmNiOTcxZThjYjNeXl50YW5naGRmc15eXjgxMDM5Ng$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-iKguaXpSAxMC5wbmdeXl50YW5naGRmc15eXjBiYzQzZWRmMmJjZjNmYjg2ZGMzMmRjOTBkOWM2ODA0Xl5edGFuZ2hkZnNeXl43NDc5ODI$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-iKguaXpSAxMS5wbmdeXl50YW5naGRmc15eXjMwMjVjOWNiZDFjZjYzOWE4M2QwMTUxOWE0MGE5NmNlXl5edGFuZ2hkZnNeXl43Njg4NTc$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-iKguaXpSAxMi5wbmdeXl50YW5naGRmc15eXjkzNzAwMDljMWE4Nzc3YmQyN2M0MmMzMmIxMDE5NGExXl5edGFuZ2hkZnNeXl41NDA0MTY$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-iKguaXpSAxMy5wbmdeXl50YW5naGRmc15eXjFmNzI4NGQwNzQxNTY4ODZkZmQ0ZGY0ZTdjODI4NGZkXl5edGFuZ2hkZnNeXl41MDg3NTI$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-iKguaXpSAxNC5wbmdeXl50YW5naGRmc15eXjVjNzgzZmM3ZDUyZTc2YWU0YjU5YzczZDJkNjEwOWZkXl5edGFuZ2hkZnNeXl44ODE5Njk$&u=1980293",
			},
			"运动": []string{
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-i_kOWKqCAwLnBuZ15eXnRhbmdoZGZzXl5eMWIxOWI5YTliM2QzY2IxNzdhYTBmMDJiYjMzNGM3YjReXl50YW5naGRmc15eXjc1NTU5MA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-i_kOWKqCAxLnBuZ15eXnRhbmdoZGZzXl5eNzEwZDRiZGI0YmNjYzViMTg4Njg0ZjBmNThjYTA5NmFeXl50YW5naGRmc15eXjg3MjY0Nw$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-i_kOWKqCAyLnBuZ15eXnRhbmdoZGZzXl5eODBmY2Q3NzRlNjQ3ZjY3NTljOTQzMDEyNWFiMDg0MmNeXl50YW5naGRmc15eXjg3MzMwMQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-i_kOWKqCAzLnBuZ15eXnRhbmdoZGZzXl5eNjBlNWQwNGJmZjdiNDRlZmJjZDFlYjc0ZWRlYjk2OTheXl50YW5naGRmc15eXjcwNjk4Mg$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-i_kOWKqCA0LnBuZ15eXnRhbmdoZGZzXl5eZDhjNTA0OWJhMzViOTRiMDk5MTRhZjRjNDBlODRiMzJeXl50YW5naGRmc15eXjU2OTQ4Nw$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-i_kOWKqCA1LnBuZ15eXnRhbmdoZGZzXl5eZmIwZjkzNjEyOTc2ZjcwNWM1M2IyZDUzNDE2YmIzYjNeXl50YW5naGRmc15eXjUyOTg3NA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-i_kOWKqCA2LnBuZ15eXnRhbmdoZGZzXl5eMjc3MTlkMTYwMGE1NGY4MGQyYjkxZWNjNmMzZmE5MjleXl50YW5naGRmc15eXjEwMjYzMDE$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-i_kOWKqCA3LnBuZ15eXnRhbmdoZGZzXl5eMTAzNzg5ZDkyMzBjN2Q4NWRjNDcyYTcyYmY4ZWQyNzNeXl50YW5naGRmc15eXjU1MDQzNA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-i_kOWKqCA4LnBuZ15eXnRhbmdoZGZzXl5eZTkwYzRlZjI1N2U4NzVmMWY3YjJiYzRkNDNmM2YxYjFeXl50YW5naGRmc15eXjcwODY4Ng$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-i_kOWKqCA5LnBuZ15eXnRhbmdoZGZzXl5eMzhjY2ViZWJkOThlMTQ0NTE0ODJjMmMyNGFhNjdlMTleXl50YW5naGRmc15eXjczMjkwMA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-i_kOWKqCAxMC5wbmdeXl50YW5naGRmc15eXmIxMDhiYTk0YzBhMTE5NzcwNDRlZjQ2NjQ0MjE2M2MzXl5edGFuZ2hkZnNeXl44OTAwMTk$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-i_kOWKqCAxMS5wbmdeXl50YW5naGRmc15eXmFmMjcxY2VmNzBlZDdlNWEzM2JhYmI0NDgxOWNmMzExXl5edGFuZ2hkZnNeXl4xMTM4NDc5$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-i_kOWKqCAxMi5wbmdeXl50YW5naGRmc15eXmM1MWNmZTk1NmM4MTk1M2Q2ZDJhMzQyNGY2MTRlMjk5Xl5edGFuZ2hkZnNeXl44Mzg2NDM$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-i_kOWKqCAxMy5wbmdeXl50YW5naGRmc15eXmU5ZmM4OTk5NzA3NDMwMTNiZDhhOTQ1MTg0NzEyODEwXl5edGFuZ2hkZnNeXl4xMDAzNzQ1$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-i_kOWKqCAxNC5wbmdeXl50YW5naGRmc15eXjY4YzgwZTBiYTY1Mjg1YjRiYTU5OGZhNjgxY2Y2MzNlXl5edGFuZ2hkZnNeXl43Nzg1NTI$&u=1980293",
			},
			"美食": []string{
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-e-jumjnyAwLnBuZ15eXnRhbmdoZGZzXl5eOGJiN2JiZWZjYjk5NjI3ZDA0YmFkNWM0ZjhjZGE5M2ZeXl50YW5naGRmc15eXjgzMTM3MQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-e-jumjnyAxLnBuZ15eXnRhbmdoZGZzXl5eZDE4OGUwOTE0N2E5NmNhZDQyMmQ5NmY2YjA2ZjE5NGReXl50YW5naGRmc15eXjc0MjIzNw$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-e-jumjnyAyLnBuZ15eXnRhbmdoZGZzXl5eNjE3NGUyMGYzNmUwNDk0ZTc1ZWRiNGZmMjI4NWYwNzJeXl50YW5naGRmc15eXjc4MTc4MA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-e-jumjnyAzLnBuZ15eXnRhbmdoZGZzXl5eY2E4ODZiMWI4MTQyOTlkNjhmOGI2YmFhMWQ0NzY1NWZeXl50YW5naGRmc15eXjYwNTYwMA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-e-jumjnyA0LnBuZ15eXnRhbmdoZGZzXl5eODYxMGQ3YjdkMDMyNjJhMTAxMjVmYjdmNjg5OTYxNzFeXl50YW5naGRmc15eXjgyOTI0OA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-e-jumjnyA1LnBuZ15eXnRhbmdoZGZzXl5eMDRjMjRlNDkxMjFjNGZmYzEzMjExY2I2MDRmZmEyMDVeXl50YW5naGRmc15eXjYwMjIzOA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-e-jumjnyA2LnBuZ15eXnRhbmdoZGZzXl5eMjg3ZGQ3ODIxMWE5YmQyYjRiNDkwNzRmMDcwMTk3YTVeXl50YW5naGRmc15eXjY5MzQwMQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-e-jumjnyA3LnBuZ15eXnRhbmdoZGZzXl5eYTk1NTY3NjMzNDkyMDU1ZmJkZmNhZDg0N2U0MzVjN2JeXl50YW5naGRmc15eXjc4NDU0Ng$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-e-jumjnyA4LnBuZ15eXnRhbmdoZGZzXl5eZGU4NDRiM2Q4Mzk2OTM3YTIzYWM2YzkxOWRmYmYxYmFeXl50YW5naGRmc15eXjEwMTkwMTU$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-e-jumjnyA5LnBuZ15eXnRhbmdoZGZzXl5eOTdkZGNjZDZjN2M5NzE1NzNiYWU2M2JmYjU0ZjJhMDZeXl50YW5naGRmc15eXjc3MzM1NQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-e-jumjnyAxMC5wbmdeXl50YW5naGRmc15eXjVmMDc5ZjUyOWU3NDlmZTNkYzk0ZDI4YjVjYzI3NjdhXl5edGFuZ2hkZnNeXl41OTcyMzk$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-e-jumjnyAxMS5wbmdeXl50YW5naGRmc15eXjRiZWQ4MGRhYzQxM2NmYTgxMTRiOWQzNjI5Y2U5ZDU4Xl5edGFuZ2hkZnNeXl45MzYzODk$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-e-jumjnyAxMi5wbmdeXl50YW5naGRmc15eXmQ2MWIzNTBjZDIwYWE2ZDU4NzRiMGE2NTZlZGU3YTYxXl5edGFuZ2hkZnNeXl44MjYzNjY$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-e-jumjnyAxMy5wbmdeXl50YW5naGRmc15eXjdhZjdkN2FlYzdhZWZlMzNiOTIzMGQ3ZjA0MTYwMjZhXl5edGFuZ2hkZnNeXl4xMDc4NjQ1$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-e-jumjnyAxNC5wbmdeXl50YW5naGRmc15eXjYxYzkzODdkYWExOWNkODg2YWQ0ZGYzN2Q2YjliNTk4Xl5edGFuZ2hkZnNeXl43ODY5NDg$&u=1980293",
			},
			"城市": []string{
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WfjuW4giAwLnBuZ15eXnRhbmdoZGZzXl5eNzFlOWNmMTU3Y2IxNDFjZDFkYjVhMzQ5NDYyMGE0NDNeXl50YW5naGRmc15eXjcwMTM3MQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WfjuW4giAxLnBuZ15eXnRhbmdoZGZzXl5eZjhlNDU3ZjI5NWM2ZmQxZDEzMDE3ZmJmMzc3NTk3MTleXl50YW5naGRmc15eXjc5NjcxOA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WfjuW4giAyLnBuZ15eXnRhbmdoZGZzXl5eNzIyNTJhN2VlYjg2YWM5ZTliM2MzODBhNTJhZjdjYmZeXl50YW5naGRmc15eXjgzNjkyMg$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WfjuW4giAzLnBuZ15eXnRhbmdoZGZzXl5eMzg1MmI2NzA0MGNkOTQ4NjZiNzA1MThhN2E5MGFlM2ReXl50YW5naGRmc15eXjg3MjUxNw$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WfjuW4giA0LnBuZ15eXnRhbmdoZGZzXl5eMTg3M2M2MDQwNGVjMDJkNzRmZmRiYzU0OGU3Nzg5MDNeXl50YW5naGRmc15eXjczMzgwMg$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WfjuW4giA1LnBuZ15eXnRhbmdoZGZzXl5eMjc3M2RhNGE0NjlkNTA0NjdmNjFiZjVmODI2NWM4YjZeXl50YW5naGRmc15eXjk3MjE0MQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WfjuW4giA2LnBuZ15eXnRhbmdoZGZzXl5eODFkOWUxYzExNTE5OTcxYmMwNGY0MzIzYzI4MjA1YmFeXl50YW5naGRmc15eXjc5NDIzNA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WfjuW4giA3LnBuZ15eXnRhbmdoZGZzXl5eZGYwZDUxYTNmZDBiOTMxMDRkZTVjZjU5YzBkNDc2OWJeXl50YW5naGRmc15eXjc3OTMyOQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WfjuW4giA4LnBuZ15eXnRhbmdoZGZzXl5eNDQ3ZjY3MTc3MjYxM2FmNDdmMWI4Mzg0ZTljYTc3ZWVeXl50YW5naGRmc15eXjg5OTcxNQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WfjuW4giA5LnBuZ15eXnRhbmdoZGZzXl5eMzlhODVmNDY0OTA1N2Y0MDUwNWE4ZWM0NjQ1ZDJjZTZeXl50YW5naGRmc15eXjcxOTg1Mw$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WfjuW4giAxMC5wbmdeXl50YW5naGRmc15eXmRkMGYzYzkwMjIwN2JiZWI1YmVhNDdjYmY1ZjlmZmZjXl5edGFuZ2hkZnNeXl44NjE3NTA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WfjuW4giAxMS5wbmdeXl50YW5naGRmc15eXmNjZTVkY2EzMzc2ZmRmNjkxYWI0NTM4OGEyZDExYTExXl5edGFuZ2hkZnNeXl44MzAyNzc$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WfjuW4giAxMi5wbmdeXl50YW5naGRmc15eXjRkNWM1OWNmZWMyY2NkZjIxZmNmZjk1YWVjMDdlODAwXl5edGFuZ2hkZnNeXl45MjQzNzE$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WfjuW4giAxMy5wbmdeXl50YW5naGRmc15eXjc1MzFjNzY2N2ZjNjRlYTg3NGQ0M2RkMGQ4N2ZiNTVkXl5edGFuZ2hkZnNeXl45NjA0NDI$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WfjuW4giAxNC5wbmdeXl50YW5naGRmc15eXmE5YzlmOTlkOTY4OTdlZjM4MDhhNzgxOTVhZDZhYzNjXl5edGFuZ2hkZnNeXl43MDE5Nzg$&u=1980293",
			},
			"风景": []string{
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-mjjuaZryAwLnBuZ15eXnRhbmdoZGZzXl5eMDc4ZDBkZDNiOWVjNmFlYThmYzI2NWMwMmU3NWYxZmZeXl50YW5naGRmc15eXjc1OTU3NQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-mjjuaZryAxLnBuZ15eXnRhbmdoZGZzXl5eYTQwZGNmMDE5NTA0MzYzOTdiNDk4NWE4NzgxNWYzZjJeXl50YW5naGRmc15eXjgxMjEwNA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-mjjuaZryAyLnBuZ15eXnRhbmdoZGZzXl5eZDJjYTU3MTJkNjMzMmRlY2RlMDI5M2RmNTMzYzEwMjJeXl50YW5naGRmc15eXjY4Mzc4Ng$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-mjjuaZryAzLnBuZ15eXnRhbmdoZGZzXl5eNzI4ODhlZTFjZDM3ZDI4ODJlYjVlMGU0N2U3ZThmMGJeXl50YW5naGRmc15eXjkxMjU1MA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-mjjuaZryA0LnBuZ15eXnRhbmdoZGZzXl5eZDEzNDRkMGE4MTQxM2Y5YWM2ZTFhZTg4NzU2MzczNTFeXl50YW5naGRmc15eXjc5MjMxMQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-mjjuaZryA1LnBuZ15eXnRhbmdoZGZzXl5eNThlZDAxYjA5ZjFiNDYwY2RjMDUzODA3MGZiYjA4YTNeXl50YW5naGRmc15eXjYzNDc3Mw$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-mjjuaZryA2LnBuZ15eXnRhbmdoZGZzXl5eZGIwYTE3YzY4ODE4MDk5YTdmY2VkYmRmMmVhZTJiNWNeXl50YW5naGRmc15eXjY5MzIxMA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-mjjuaZryA3LnBuZ15eXnRhbmdoZGZzXl5eNTE2ZmRkMDBjY2U3NWY4ODFlNjMwYTc1YmRlYmYzZjleXl50YW5naGRmc15eXjc3NDUwOQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-mjjuaZryA4LnBuZ15eXnRhbmdoZGZzXl5eZGQ3NDZiNGY5NGE3MzQxNDVkNTAzYTFjODUzNTg2ZTNeXl50YW5naGRmc15eXjY5NzM1MQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-mjjuaZryA5LnBuZ15eXnRhbmdoZGZzXl5eMDEwNGM1YWVmNmRjZTYyYTY0ZTE3YzUxZTdiZTcxNjNeXl50YW5naGRmc15eXjcwMTA3Ng$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-mjjuaZryAxMC5wbmdeXl50YW5naGRmc15eXjg1NjY5ODcxYmJiZjhlYWFlMTFkMDU5YzE1NzA2ODQzXl5edGFuZ2hkZnNeXl42MzYzMDU$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-mjjuaZryAxMS5wbmdeXl50YW5naGRmc15eXjM3ZTQxOTVmZDZlMWY0YTI1MzNjNTNkNDU3ZmUwYWJlXl5edGFuZ2hkZnNeXl43MzE3MTM$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-mjjuaZryAxMi5wbmdeXl50YW5naGRmc15eXmM5ZTkwOTBmNmFmM2FiMTlkN2RjODdhYmM5ZGFiNzdmXl5edGFuZ2hkZnNeXl4xMDcyNDA1$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-mjjuaZryAxMy5wbmdeXl50YW5naGRmc15eXjQ4NDZhMWViOGJkZmI1ZTcyNGQ1MDcwZjI4NDUzMzE5Xl5edGFuZ2hkZnNeXl41ODk1MjA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-mjjuaZryAxNC5wbmdeXl50YW5naGRmc15eXmNlZDE5ZTU0ZGQ0ZGE4ZGU3YzY2OTAzMmM4NDgwOWFiXl5edGFuZ2hkZnNeXl41OTA1NzI$&u=1980293",
			},
			"季节": []string{
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-Wto-iKgiAwLnBuZ15eXnRhbmdoZGZzXl5eNzM3Yzg2YmYyMDdiZjJmN2U4OWM3OTNiZWU1YTdkOTdeXl50YW5naGRmc15eXjc0ODA1NQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-Wto-iKgiAxLnBuZ15eXnRhbmdoZGZzXl5eYzJkZmQ5M2ZhYTQ5NjFiNGIyOGI0YWQzN2ViYjZkM2ZeXl50YW5naGRmc15eXjUyNjY0NA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-Wto-iKgiAyLnBuZ15eXnRhbmdoZGZzXl5eM2VhMjcyYTQyYThkMWJhMmUyNGZmYmVlOGJjMjU0MGJeXl50YW5naGRmc15eXjYzMjU5NQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-Wto-iKgiAzLnBuZ15eXnRhbmdoZGZzXl5eZmFkNjE1Zjk3NTJkNGY4MTdmNzhkMjYxMGFlNmZhYTdeXl50YW5naGRmc15eXjk2OTgxNQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-Wto-iKgiA0LnBuZ15eXnRhbmdoZGZzXl5eODBiNzhjNDMyYTY3MDhkZjkwNGY3MjkzZWE2ODU3MTFeXl50YW5naGRmc15eXjc4NDQxOA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-Wto-iKgiA1LnBuZ15eXnRhbmdoZGZzXl5eZWZjNzQ2NWEwZmJjMTUwYTA1Y2ExNDZiMTUwODQ2ZmNeXl50YW5naGRmc15eXjgzOTA1MQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-Wto-iKgiA2LnBuZ15eXnRhbmdoZGZzXl5eYTM5NDU1ZDdhMTEyNjc5ZjFkYzI4Yzc2OWU2NmY5NDNeXl50YW5naGRmc15eXjc5NzM1Nw$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-Wto-iKgiA3LnBuZ15eXnRhbmdoZGZzXl5eNTdlOWI1ZDZlZDMzZTE4MmRiYTMzNGU3YTJlZmE2MjFeXl50YW5naGRmc15eXjcyMDUzMw$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-Wto-iKgiA4LnBuZ15eXnRhbmdoZGZzXl5eOGJlNTA3ZTY3M2IwZGQ2YzUxZDg4ZjE5MWE3MDY1ZTNeXl50YW5naGRmc15eXjc0Nzg5OQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-Wto-iKgiA5LnBuZ15eXnRhbmdoZGZzXl5eMzRjN2JjOGJiNzkzYjY5ZDk2MjJiYjI1ZWIzMzM1Y2NeXl50YW5naGRmc15eXjk5ODQzOQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-Wto-iKgiAxMC5wbmdeXl50YW5naGRmc15eXjM0NmFlOTA0OGI2MGRmZmQ3OGE3NjlmMTlkZTY3NjVlXl5edGFuZ2hkZnNeXl44MDc2Nzc$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-Wto-iKgiAxMS5wbmdeXl50YW5naGRmc15eXjdkNGQ3NzcxNGU0NDM0ZmIwZDVlNmZlNzNiMTA0ZDY2Xl5edGFuZ2hkZnNeXl4zNjU4NjI$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-Wto-iKgiAxMi5wbmdeXl50YW5naGRmc15eXmFmMmNhNjBhOWYyYzE4OTRhMjA2M2IzNjk0MTU1MTg1Xl5edGFuZ2hkZnNeXl45MTQxMzE$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-Wto-iKgiAxMy5wbmdeXl50YW5naGRmc15eXjllYzgxM2FmYTE3YjhlZTI2NzczMDU0NzA3ODUyODczXl5edGFuZ2hkZnNeXl45MjA2ODg$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-Wto-iKgiAxNC5wbmdeXl50YW5naGRmc15eXjc2NjMxZDNjZTE1OTY3MTY0MzZmNjU5ODY0NTlhMmUxXl5edGFuZ2hkZnNeXl44OTcyMzg$&u=1980293",
			},
			"健康": []string{
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WBpeW6tyAwLnBuZ15eXnRhbmdoZGZzXl5eM2Q5ZWY2NGI1NmJkYWFiN2RkMTI4NzQ3MmU2MGU0MDdeXl50YW5naGRmc15eXjU1MzY2OQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WBpeW6tyAxLnBuZ15eXnRhbmdoZGZzXl5eY2VkNjRhYjJiNmRiMjljYjJiOTg1ZWNkNzJkYzEyZDheXl50YW5naGRmc15eXjczODA1MQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WBpeW6tyAyLnBuZ15eXnRhbmdoZGZzXl5eM2ZlNTk2MjcyMDdiYTA4ZGM1OGRhMjk1YjEzOGVmMmNeXl50YW5naGRmc15eXjU5MDIyMg$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WBpeW6tyAzLnBuZ15eXnRhbmdoZGZzXl5eMGFmYTYxOTg3MjFkMWRjOTFiZjUxYzYxNmZiMmRiNTFeXl50YW5naGRmc15eXjc3Njg3NQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WBpeW6tyA0LnBuZ15eXnRhbmdoZGZzXl5eOWQ2YjczMDEzMmI0N2UxNjM2ZjY5MTQ0NTYzZTZlMTJeXl50YW5naGRmc15eXjcxMjI3MQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WBpeW6tyA1LnBuZ15eXnRhbmdoZGZzXl5eNjdkZDMyY2Y3ODFkZjU4Y2VlNzA3MGQ3Y2FjMmFlOTJeXl50YW5naGRmc15eXjYzNTg2Mg$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WBpeW6tyA2LnBuZ15eXnRhbmdoZGZzXl5eNzllZDIyNTk0ODVjNTVjMGVhNmY4ODA2M2JmY2E5YzFeXl50YW5naGRmc15eXjczNzY2MQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WBpeW6tyA3LnBuZ15eXnRhbmdoZGZzXl5eMDY3N2FmMzg0MjE4ZGYwODU3Y2Y3OTNmZGIyZTI3N2VeXl50YW5naGRmc15eXjkzODAwNA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WBpeW6tyA4LnBuZ15eXnRhbmdoZGZzXl5eM2FiZGY0OGI1ZDgzMDdiZmNmMzM0NWZhYjExNzRhZWReXl50YW5naGRmc15eXjYxMDEyOA$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WBpeW6tyA5LnBuZ15eXnRhbmdoZGZzXl5eMTg4ODkzYzhmOWEyYTIxODZjMWRmOWY5MjE3YWUyNWJeXl50YW5naGRmc15eXjgzNDA5NQ$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WBpeW6tyAxMC5wbmdeXl50YW5naGRmc15eXjBiYzQzZWRmMmJjZjNmYjg2ZGMzMmRjOTBkOWM2ODA0Xl5edGFuZ2hkZnNeXl43NDc5ODI$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WBpeW6tyAxMS5wbmdeXl50YW5naGRmc15eXjRjZDY1MGY1MzQwYWRjNzRjNzI3ZWU3OGM5MGFlYTU5Xl5edGFuZ2hkZnNeXl42MDQ2NDc$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WBpeW6tyAxMi5wbmdeXl50YW5naGRmc15eXjE5N2NjY2QyODY0NWE1NDRlOTAyYTIwMDY5Y2FiMTA0Xl5edGFuZ2hkZnNeXl42MDY4NDE$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WBpeW6tyAxMy5wbmdeXl50YW5naGRmc15eXjU2ODM0MDMzZjRhYzFhNmUxYmJiNGMzNWJmYmFiOTZjXl5edGFuZ2hkZnNeXl40NDQzOTk$&u=1980293",
				"http://oncloud.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-WBpeW6tyAxNC5wbmdeXl50YW5naGRmc15eXmNiOTQxODlhMjMyMTI5YjdkNDdlYWQ2ZTgzYTE4Zjk5Xl5edGFuZ2hkZnNeXl40MTAxMzM$&u=1980293",
			},
		},
	}
	FeedIconsData.TestCloud = map[string]string{
		"task": "http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjM2NzI1MDUvOC_ku7vliqHlj7df5Zu-5qCHX-W5v-aSreS7u-WKoUAyeC5wbmdeXl50YW5naGRmc15eXmEzOTA0MzZmYjQwMmQ3YzgzNDk1ODAwNzQ4NjRmOTk2Xl5edGFuZ2hkZnNeXl44Mjg$&u=63672505",
	}
	FeedIconsData.OnCloud = map[string]string{
		"task": "http://beefs.quanshi.com:80/ucfserver/hddown?fid=MTk4MDI5My84L-S7u-WKoeWPt1_lm77moIdf5bm_5pKt5Lu75YqhNDAyeC5wbmdeXl50YW5naGRmc15eXmEzOTA0MzZmYjQwMmQ3YzgzNDk1ODAwNzQ4NjRmOTk2Xl5edGFuZ2hkZnNeXl44Mjg$&u=1980293",
	}
}

//Init 初始化配置
func Init(option map[string]string) error {
	//App 域名
	serverName, ok := option["server_name"]
	if !ok {
		return fmt.Errorf(`config "server_name" is undefined`)
	}
	ServerName = serverName
	return nil
}

//GetBoardAvatar 查询默认广播号头像
func GetBoardAvatar() []string {
	switch ServerName {
	case "testcloud", "testcloudb", "testcloud3":
		return BoardAvatarData.TestCloud
	case "oncloud", "oncloud2":
		return BoardAvatarData.OnCloud
	}
	return nil
}

//GetBbsThumb 查询默认广播封面
func GetBbsThumb() map[string][]string {
	switch ServerName {
	case "testcloud", "testcloudb", "testcloud3":
		return BbsThumbData.TestCloud
	case "oncloud", "oncloud2":
		return BbsThumbData.OnCloud
	}
	return nil
}

//GetFeedIcons 查询feed图标
func GetFeedIcons(iconName string) string {
	switch ServerName {
	case "testcloud", "testcloudb", "testcloud3":
		return FeedIconsData.TestCloud[iconName]
	case "oncloud", "oncloud2":
		return FeedIconsData.OnCloud[iconName]
	}
	return ""
}
