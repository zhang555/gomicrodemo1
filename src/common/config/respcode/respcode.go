package respcode

type CodeMsg struct {
	Code int32
	Msg  string
}

var (
	Unauth                = CodeMsg{Code: 300000, Msg: "没有登录，请先登录"}
	IdCanNot0             = CodeMsg{Code: 300888, Msg: "id不能为0"}
	UsernamePasswordError = CodeMsg{Code: 300882, Msg: "账号密码错误"}
)
