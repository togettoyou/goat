package e

// 例: 10001
// 第一位的1表示服务级错误 (1 为系统级错误；2 为普通错误，通常是由用户非法操作引起的)
// 第二位至第三位的00表示服务模块代码
// 最后两位01表示具体错误代码

var (
	OK                  = &errno{code: 0, msg: "成功"}
	InternalServerError = &errno{code: 10001, msg: "服务器异常"}
	DBError             = &errno{code: 10002, msg: "DB异常"}

	ErrValidation = &errno{code: 20001, msg: "参数校验失败"}
	ErrBind       = &errno{code: 20002, msg: "参数绑定异常"}

	ErrNotLogin = &errno{code: 20101, msg: "请登录"}
	// 可自行扩展...
)
