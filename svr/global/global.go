package global

type GroupType int
type RetCode int

//用户权限组
const (
	GROUP_ADMIN GroupType = iota
	GROUP_USER
	GROUP_UNDEF
)

//api返回code枚举
const (
	RET_OK RetCode = iota
	RET_ERR
	RET_ERR_DB
	RET_ERR_HTTP_QUERY
	RET_ERR_USER_PASSWD
	RET_ERR_ACCESS_TOKEN
	RET_ERR_CREATE_TOKEN
	RET_ERR_NO_RIGHT
	RET_ERR_SESSION_INVALID
)

//数据库表名
const BUCKET_USR_PASSWD = "usr:passwd"
const BUCKET_TOKEN_INFO = "token:info"
const BUCKET_USER_TOKEN = "user:token"

//cookie名
const ACCESS_TOKEN = "access_token"

//session key
const SESSION_KEY_USER = "user"
const SESSION_KEY_GROUP = "group"
