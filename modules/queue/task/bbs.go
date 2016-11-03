package task

//Bbs 图文广播
type Bbs struct {
	base
}

//Run 启动任务处理
func (b *Bbs) Run() error {
	return nil
}
func init() {
	Register("bbs", new(Bbs))
}
