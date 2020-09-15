package handler

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/adayswait/mojo/db"
	"github.com/adayswait/mojo/global"
	"github.com/adayswait/mojo/mlog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
	"github.com/google/goexpect"
	"regexp"
	"time"
)

var sessions *session.Session

func init() {
	sessions = session.New()
}

type SecInfo struct {
	User   string `json:"user"`
	Passwd string `json:"passwd"`
}

func Login(c *fiber.Ctx) error {
	var secInfo SecInfo
	var accessToken string
	if err := c.QueryParser(&secInfo); err == nil {
		if len(secInfo.User) != 0 && len(secInfo.Passwd) != 0 {
			//优先使用用户名和密码进行验证
			tempMD5 := md5.New()
			tempMD5.Write([]byte(secInfo.Passwd))
			tempMD5.Write([]byte(global.MD5_SALT))
			passwdMD5 := hex.EncodeToString(tempMD5.Sum(nil))
			eau := db.Auth(secInfo.User, passwdMD5)
			if eau != nil {
				return c.JSON(fiber.Map{"code": global.RET_ERR_USER_PASSWD,
					"data": eau.Error()})
			}
			token, etk := db.GetAccessToken(secInfo.User)
			if etk != nil {
				return c.JSON(fiber.Map{"code": global.RET_ERR_DB,
					"data": etk.Error()})
			}
			accessToken = token
		}
		//取不到user和passwd, 后面尝试access_token登录
	}
	if len(accessToken) == 0 {
		//用户名密码验证未成功, 再尝试使用access_token进行验证
		accessToken = c.Cookies(global.ACCESS_TOKEN)
	} else {
		cookie := new(fiber.Cookie)
		cookie.Name = global.ACCESS_TOKEN
		cookie.Expires = time.Now().Add(365 * 24 * time.Hour) // maybe too long
		cookie.Value = accessToken
		c.Cookie(cookie)
	}

	if len(accessToken) == 0 {
		return c.JSON(fiber.Map{"code": global.RET_ERR_ACCESS_TOKEN,
			"data": "no access_token found"})

	}
	userInfo, elg := db.Login(accessToken)
	if elg != nil {
		return c.JSON(fiber.Map{"code": global.RET_ERR_ACCESS_TOKEN,
			"data": elg.Error()})

	}
	data, _ := json.Marshal(userInfo)

	//create session
	go (func() {
		store := sessions.Get(c)
		store.Set(global.SESSION_KEY_USER, userInfo.User)
		store.Set(global.SESSION_KEY_GROUP, userInfo.Group)
		defer store.Save()
	})()
	return c.JSON(fiber.Map{"code": global.RET_OK, "data": string(data)})
}

func Logout(c *fiber.Ctx) error {
	store := sessions.Get(c)
	store.Destroy()

	//todo
	cookie := new(fiber.Cookie)
	cookie.Name = global.ACCESS_TOKEN
	cookie.Value = ""
	c.Cookie(cookie)

	defer store.Save()
	return c.JSON(fiber.Map{"code": global.RET_OK, "data": nil})
}

func Register(c *fiber.Ctx) error {
	var secInfo SecInfo
	err := c.QueryParser(&secInfo)
	if err != nil {
		return c.JSON(fiber.Map{"code": global.RET_ERR_HTTP_QUERY,
			"data": err.Error()})

	}
	if len(secInfo.User) == 0 || len(secInfo.Passwd) == 0 {
		return c.JSON(fiber.Map{"code": global.RET_ERR_INPUT,
			"data": "user or password is empty"})

	}
	tempMD5 := md5.New()
	tempMD5.Write([]byte(secInfo.Passwd))
	tempMD5.Write([]byte(global.MD5_SALT))
	passwdMD5 := hex.EncodeToString(tempMD5.Sum(nil))

	edb := db.Register(secInfo.User, passwdMD5)
	if edb != nil {
		return c.JSON(fiber.Map{"code": global.RET_ERR_DB, "data": edb.Error()})
	}
	accessToken, etk := db.GetAccessToken(secInfo.User)
	if etk != nil {
		return c.JSON(fiber.Map{"code": global.RET_ERR_DB, "data": etk.Error()})
	}
	cookie := new(fiber.Cookie)
	cookie.Name = global.ACCESS_TOKEN
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Value = accessToken
	c.Cookie(cookie)
	return c.JSON(fiber.Map{"code": global.RET_OK, "data": nil})
}

func ViewDB(c *fiber.Ctx) error {
	store := sessions.Get(c)
	// user := store.Get(global.SESSION_KEY_USER)
	group := store.Get(global.SESSION_KEY_GROUP)
	if group == nil {
		return c.JSON(fiber.Map{"code": global.RET_ERR_SESSION_INVALID,
			"data": "session invalid"})
	}
	if int(group.(int64)) > int(global.GROUP_UNDEF) {
		return c.JSON(fiber.Map{"code": global.RET_ERR_NO_RIGHT,
			"data": "no right to do this"})
	}
	table := c.Params("table")
	if len(table) == 0 {
		ret, err := db.Buckets()
		if err != nil {
			return c.JSON(fiber.Map{"code": global.RET_ERR_DB, "data": err.Error()})
		}
		return c.JSON(fiber.Map{"code": global.RET_OK, "data": ret})
	}
	key := c.Params("key")
	if len(key) == 0 {
		ret, err := db.Keys(table)
		if err != nil {
			return c.JSON(fiber.Map{"code": global.RET_ERR_DB, "data": err.Error()})
		}
		return c.JSON(fiber.Map{"code": global.RET_OK, "data": ret})
	} else {
		ret, err := db.Get(table, key)
		if err != nil {
			return c.JSON(fiber.Map{"code": global.RET_ERR_DB, "data": err.Error()})
		} else {
			return c.JSON(fiber.Map{"code": global.RET_OK, "data": string(ret)})
		}
	}
}

func UpdateDB(c *fiber.Ctx) error {
	store := sessions.Get(c)
	// user := store.Get(global.SESSION_KEY_USER)
	group := store.Get(global.SESSION_KEY_GROUP)
	if group == nil {
		return c.JSON(fiber.Map{"code": global.RET_ERR_SESSION_INVALID,
			"data": "session invalid"})
	}
	if int(group.(int64)) > int(global.GROUP_ADMIN) {
		return c.JSON(fiber.Map{"code": global.RET_ERR_NO_RIGHT,
			"data": "no right to do this"})
	}
	table := c.Params("table")
	key := c.Params("key")
	if len(table) == 0 {
		return c.JSON(fiber.Map{"code": global.RET_ERR_URL_PARAM,
			"data": "request param err"})
	}

	body := struct {
		Value string `json:"value"`
	}{}

	if errBp := c.BodyParser(&body); errBp != nil {
		return c.JSON(fiber.Map{"code": global.RET_ERR_BODY_PARAM,
			"data": errBp.Error()})
	}
	err := db.Set(table, key, body.Value)
	if err != nil {
		return c.JSON(fiber.Map{"code": global.RET_ERR_DB,
			"data": err.Error()})
	}
	return c.JSON(fiber.Map{"code": global.RET_OK, "data": nil})
}

func DeleteDB(c *fiber.Ctx) error {
	store := sessions.Get(c)
	// user := store.Get(global.SESSION_KEY_USER)
	group := store.Get(global.SESSION_KEY_GROUP)
	if group == nil {
		return c.JSON(fiber.Map{"code": global.RET_ERR_SESSION_INVALID,
			"data": "session invalid"})
	}
	if int(group.(int64)) > int(global.GROUP_WHOSYOURDADDY) {
		return c.JSON(fiber.Map{"code": global.RET_ERR_NO_RIGHT,
			"data": "no right to do this"})
	}
	table := c.Params("table")
	key := c.Params("key")
	if len(table) == 0 || len(key) == 0 {
		return c.JSON(fiber.Map{"code": global.RET_ERR_URL_PARAM,
			"data": "request param err"})
	}
	err := db.Delete(table, key)
	if err != nil {
		return c.JSON(fiber.Map{"code": global.RET_ERR_DB,
			"data": err.Error()})
	}
	return c.JSON(fiber.Map{"code": global.RET_OK, "data": nil})
}

func Test(c *fiber.Ctx) error {
	const timeout = time.Minute
	e, _, err := expect.Spawn("node", -1)
	defer e.Close()
	if err != nil {
		return c.JSON(fiber.Map{"code": global.RET_ERR_SPAWN,
			"data": "cmd"})
	}
	checkedOutRE := regexp.MustCompile(">")
	ret, _, err2 := e.Expect(checkedOutRE, timeout)
	mlog.Log(ret, err2)
	e.Send("1+1\n")
	ret2, _, err3 := e.Expect(checkedOutRE, timeout)
	mlog.Log(ret2, err3)
	return c.JSON(fiber.Map{"code": global.RET_OK,
		"data": ret})
}

var NewDB = UpdateDB
