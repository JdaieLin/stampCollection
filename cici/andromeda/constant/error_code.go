package constant

const (
	ERR_UNKNOWN     = 100
	ERR_UNSUPPORTED = 101
	// 通信错误
	ERR_REQ = iota + 10100
	ERR_REQ_NET
	ERR_REQ_BODY
	ERR_REQ_JSON_ENCODE
	ERR_REQ_SERVER
	ERR_REQ_CLIENT_VERSOIN

	// 注册错误
	ERR_REG = iota + 20100
	ERR_REG_LOGIN_ID_NULL
	ERR_REG_LOGIN_ID_EXIST
	ERR_REG_NAME_NULL
	ERR_REG_NAME_EXIST
	ERR_REG_LOGIN_PASSWORD_NULL
	ERR_REG_MAIL_EXIST
	ERR_REG_PHONE_EXIST
	ERR_REG_IDCARD_EXIST

	// 登录错误
	ERR_LOGIN = iota + 20200
	ERR_LOGIN_LOGIN_ID_INVAID
	ERR_LOGIN_PASSWORD_INVAID
	ERR_LOGIN_FAILD_MANY

	// 邮票操作
	ERR_STAMP_INVAID = iota + 30100
)

var ERR_DICT = map[int]string{
	ERR_UNSUPPORTED:        "尚未支持的接口:%s",
	ERR_UNKNOWN:            "未知错误",
	ERR_REQ:                "请求处理失败",
	ERR_REQ_NET:            "网络错误",
	ERR_REQ_JSON_ENCODE:    "json序列化失败:%s",
	ERR_REQ_BODY:           "请求解析失败:%s",
	ERR_REQ_SERVER:         "服务端异常:%s",
	ERR_REQ_CLIENT_VERSOIN: "客户端版本不支持:%s",

	// 注册错误
	ERR_REG:                     "注册失败",
	ERR_REG_LOGIN_ID_NULL:       "登录id为空",
	ERR_REG_LOGIN_ID_EXIST:      "登录id已存在:%s",
	ERR_REG_NAME_NULL:           "用户昵称为空",
	ERR_REG_NAME_EXIST:          "用户昵称已存在:%s",
	ERR_REG_LOGIN_PASSWORD_NULL: "登录密码为空",
	ERR_REG_MAIL_EXIST:          "邮箱已存在:%s",
	ERR_REG_PHONE_EXIST:         "手机号已存在:%s",
	ERR_REG_IDCARD_EXIST:        "身份证已存在:%s",

	// 登录错误
	ERR_LOGIN:                 "登录失败",
	ERR_LOGIN_LOGIN_ID_INVAID: "用户名不存在",
	ERR_LOGIN_PASSWORD_INVAID: "密码错误",
	ERR_LOGIN_FAILD_MANY:      "登录失败太多次",

	//操作失败
	ERR_STAMP_INVAID: "邮票查找失败：%d",
}

func GetErrorInfo(code int) string {
	msg, ok := ERR_DICT[code]
	if !ok {
		msg = ERR_DICT[ERR_UNKNOWN]
	}
	return msg
}
