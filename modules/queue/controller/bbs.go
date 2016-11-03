package controller

import (
	"github.com/ccyun/daemon/example/app/common"
)

type Bbs struct {
}

func (b *Bbs) Run() error {
	return nil
}
func init() {
	common.Register("bbs", new(Bbs))
}
