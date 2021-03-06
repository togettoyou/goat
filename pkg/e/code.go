package e

// 例: 10001
// 第一位的1表示服务级错误 (1 为系统级错误；2 为普通错误，通常是由用户非法操作引起的)
// 第二位至第三位的00表示服务模块代码
// 最后两位01表示具体错误代码

var (
	OK          = &errno{code: 0, msg: "成功"}
	ServerError = &errno{code: 10001, msg: "服务器异常"}
	DBError     = &errno{code: 10002, msg: "DB异常"}
	BindError   = &errno{code: 10003, msg: "参数绑定异常"}

	ErrValidation = &errno{code: 20001, msg: "参数校验失败"}

	ErrNotLogin     = &errno{code: 20101, msg: "请登录"}
	ErrPassword     = &errno{code: 20102, msg: "密码错误"}
	ErrTokenGen     = &errno{code: 20103, msg: "令牌生成失败"}
	ErrTokenExpired = &errno{code: 20104, msg: "令牌已过期"}
	ErrTokenInvalid = &errno{code: 20105, msg: "令牌无效"}
	ErrTokenFailure = &errno{code: 20106, msg: "令牌验证失败"}
	// 可自行扩展...
)
