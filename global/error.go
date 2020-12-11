package global

type AppErr struct {
	Code string
	Msg  string
}

var (
	ErrLoginFailed              = &AppErr{Code: "1000", Msg: "用户名或密码错误"}
	ErrAuthFailed               = &AppErr{Code: "1001", Msg: "鉴权失败"}
	ErrInvalidParam             = &AppErr{Code: "1002", Msg: "请求参数不正确"}
	ErrInvalidToken             = &AppErr{Code: "1003", Msg: "鉴权失败，请传入token"}
	ErrRequestValidationNotPass = &AppErr{Code: "1005", Msg: "请求有效性检查未通过"}
	ErrDataNotExist             = &AppErr{Code: "1006", Msg: "记录不存在"}
	ErrDataAlreadyExist         = &AppErr{Code: "1007", Msg: "记录已存在"}
	ErrTryLater                 = &AppErr{Code: "1008", Msg: "系统繁忙，请稍后重试"}
	ErrLogin                    = &AppErr{Code: "2001", Msg: "token失效，请重新登陆"}
	ErrAlreadyLogin             = &AppErr{Code: "2002", Msg: "请勿重复登录"}
	ErrEmailAlreadyExist        = &AppErr{Code: "2003", Msg: "邮箱已存在"}
	ErrUsernameAlreadyExist     = &AppErr{Code: "2004", Msg: "用户名已存在"}
	ErrPasswordNotTheSame       = &AppErr{Code: "2005", Msg: "新密码不一致"}
	ErrOldPasswordNotRight      = &AppErr{Code: "2006", Msg: "旧密码不正确"}
	ErrPhoneNotExited           = &AppErr{Code: "2007", Msg: "手机号不存在"}
	ErrUserNameNotExited        = &AppErr{Code: "2008", Msg: "用户名不存在"}
	ErrInvalidPassword          = &AppErr{Code: "2009", Msg: "密码错误"}
	ErrBusinessCheckNotPass     = &AppErr{Code: "2010", Msg: "业务检查未通过"}
	ErrInternal                 = &AppErr{Code: "9999", Msg: "系统错误"}
)
