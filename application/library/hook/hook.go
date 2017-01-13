package hook

//hooks 钩子
type hooks struct {
	AppRunStart []func() error
}

//Hooks 钩子
var Hooks hooks

//AppRunStart App运行开始
func AppRunStart() error {
	for _, f := range Hooks.AppRunStart {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
