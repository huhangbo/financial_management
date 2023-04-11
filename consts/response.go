package consts

const (
	SuccessCode         = 200
	FailCode            = 400
	AuthErrorCode       = 401
	NotFoundCode        = 404
	SystemErrorCode     = 500
	UserExistCode       = 1001
	UserNotExistCode    = 1002
	PasswordErrorCode   = 1003
	ParamErrorCode      = 1004
	PermissionErrorCode = 1005
	EmptyCode           = 1006
)

const (
	SuccessMsg         = "成功"
	FailMsg            = "请求失败"
	AuthErrorMsg       = "用户认证失败"
	NotFoundMsg        = "未找到"
	SystemErrorRMsg    = "系统错误"
	UserExistMsg       = "用户已存在"
	UserNotExistMsg    = "用户不存在"
	PasswordErrorMsg   = "密码错误"
	ParamErrorMsg      = "参数提交错误"
	PermissionErrorMsg = "没有权限"
)

var (
	ResponseCode2Msg = map[int]string{
		SuccessCode:         SuccessMsg,
		FailCode:            FailMsg,
		AuthErrorCode:       AuthErrorMsg,
		NotFoundCode:        NotFoundMsg,
		SystemErrorCode:     SystemErrorRMsg,
		UserExistCode:       UserExistMsg,
		UserNotExistCode:    UserNotExistMsg,
		PasswordErrorCode:   PasswordErrorMsg,
		ParamErrorCode:      ParamErrorMsg,
		PermissionErrorCode: PermissionErrorMsg,
	}
)
