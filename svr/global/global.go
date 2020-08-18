package global

type GroupType int
type RetCode int

//用户权限组
const (
	GROUP_WHOSYOURDADDY GroupType = iota
	GROUP_ADMIN
	GROUP_USER
	GROUP_VIEW
	GROUP_UNDEF
)

//api返回code枚举
const (
	RET_OK RetCode = iota
	RET_ERR
	RET_ERR_DB
	RET_ERR_HTTP_QUERY
	RET_ERR_USER_PASSWD
	RET_ERR_INPUT //5
	RET_ERR_URL_PARAM
	RET_ERR_BODY_PARAM
	RET_ERR_ACCESS_TOKEN
	RET_ERR_CREATE_TOKEN
	RET_ERR_NO_RIGHT //10
	RET_ERR_SESSION_INVALID
	RET_ERR_SPAWN
	RET_ERR_SPAWN_EXPECT
	RET_ERR_SPAWN_EXPECT_MATCH
	RET_ERR_DEP_EXPIRED //15
	RET_ERR_DEP_BREAK_TOO_LATE
	RET_ERR_FORCE_DEP_CANT_BREAK
	RET_ERR_HTTP_REQUEST
	RET_ERR_OS_EXEC_CMD_RUN
)

const (
	DEP_STATUS_NOT_START int = iota
	DEP_STATUS_CHECKOUT
	DEP_STATUS_SLEEP
	DEP_STATUS_SYNC
	DEP_STATUS_STOP_SERVER
	DEP_STATUS_START_SERVER
	DEP_STATUS_SUCCESS
)
const DEP_STATUS_ERR int = -1
const DEP_STATUS_ERR_CHECKOUT = -2
const DEP_STATUS_ERR_GETINI = -3

//数据库表名
const BUCKET_ITEMS_DESC = "sys:items:desc"
const BUCKET_USR_PASSWD = "sys:usr:passwd"
const BUCKET_TOKEN_INFO = "sys:token:info"
const BUCKET_USER_TOKEN = "sys:user:token"
const BUCKET_USR_OPSLOG = "sys:usr:opslog"
const BUCKET_OPS_MACINI = "sys:ops:macini"
const BUCKET_OPS_DEPINI = "sys:ops:depini"
const BUCKET_OPS_DEVINI = "sys:ops:devini"
const BUCKET_OPS_DEPBIL = "sys:ops:depbil"

//cookie名
const ACCESS_TOKEN = "access_token"

//session key
const SESSION_KEY_USER = "user"
const SESSION_KEY_GROUP = "group"

const MD5_SALT = "mojo_salt"

const DINGDING_TEXT_MSG_PATTERN = `{"msgtype": "text","text": {"content": "%s"}}`
