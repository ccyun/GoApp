package controller

//Index 首页
type Index struct {
	Base
}

//Index 首页
func (I *Index) Index() {
	I.Ctx.WriteString("hello beego")

}
