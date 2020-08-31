package handler

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/adayswait/mojo/cmd"
	"github.com/adayswait/mojo/db"
	"github.com/adayswait/mojo/global"
	"github.com/adayswait/mojo/mlog"
	"github.com/adayswait/mojo/utils"
	"github.com/gofiber/fiber"
	"github.com/gofiber/session"
	"github.com/google/goexpect"
	"regexp"
	"strconv"
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

func Login(c *fiber.Ctx) {
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
				c.JSON(fiber.Map{"code": global.RET_ERR_USER_PASSWD,
					"data": eau.Error()})
				return
			}
			token, etk := db.GetAccessToken(secInfo.User)
			if etk != nil {
				c.JSON(fiber.Map{"code": global.RET_ERR_DB,
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
		c.JSON(fiber.Map{"code": global.RET_ERR_ACCESS_TOKEN,
			"data": "no access_token found"})
		return
	}
	userInfo, elg := db.Login(accessToken)
	if elg != nil {
		c.JSON(fiber.Map{"code": global.RET_ERR_ACCESS_TOKEN,
			"data": elg.Error()})
		return
	}
	data, _ := json.Marshal(userInfo)
	c.JSON(fiber.Map{"code": global.RET_OK, "data": string(data)})

	//create session
	(func() {
		store := sessions.Get(c)
		store.Set(global.SESSION_KEY_USER, userInfo.User)
		store.Set(global.SESSION_KEY_GROUP, userInfo.Group)
		defer store.Save()
	})()
}

func Logout(c *fiber.Ctx) {
	store := sessions.Get(c)
	store.Destroy()

	//todo
	cookie := new(fiber.Cookie)
	cookie.Name = global.ACCESS_TOKEN
	cookie.Value = ""
	c.Cookie(cookie)

	defer store.Save()
	c.JSON(fiber.Map{"code": global.RET_OK, "data": nil})
}

func Register(c *fiber.Ctx) {
	var secInfo SecInfo
	err := c.QueryParser(&secInfo)
	if err != nil {
		c.JSON(fiber.Map{"code": global.RET_ERR_HTTP_QUERY,
			"data": err.Error()})
		return
	}
	if len(secInfo.User) == 0 || len(secInfo.Passwd) == 0 {
		c.JSON(fiber.Map{"code": global.RET_ERR_INPUT,
			"data": "user or password is empty"})
		return
	}
	tempMD5 := md5.New()
	tempMD5.Write([]byte(secInfo.Passwd))
	tempMD5.Write([]byte(global.MD5_SALT))
	passwdMD5 := hex.EncodeToString(tempMD5.Sum(nil))

	edb := db.Register(secInfo.User, passwdMD5)
	if edb != nil {
		c.JSON(fiber.Map{"code": global.RET_ERR_DB, "data": edb.Error()})
		return
	}
	accessToken, etk := db.GetAccessToken(secInfo.User)
	if etk != nil {
		c.JSON(fiber.Map{"code": global.RET_ERR_DB, "data": etk.Error()})
		return
	}
	cookie := new(fiber.Cookie)
	cookie.Name = global.ACCESS_TOKEN
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Value = accessToken
	c.Cookie(cookie)
	c.JSON(fiber.Map{"code": global.RET_OK, "data": nil})
}

func ViewDB(c *fiber.Ctx) {
	store := sessions.Get(c)
	// user := store.Get(global.SESSION_KEY_USER)
	group := store.Get(global.SESSION_KEY_GROUP)
	if group == nil {
		c.JSON(fiber.Map{"code": global.RET_ERR_SESSION_INVALID,
			"data": "session invalid"})
		return
	}
	if int(group.(int64)) > int(global.GROUP_UNDEF) {
		c.JSON(fiber.Map{"code": global.RET_ERR_NO_RIGHT,
			"data": "no right to do this"})
		return
	}
	table := c.Params("table")
	if len(table) == 0 {
		ret, err := db.Buckets()
		if err != nil {
			c.JSON(fiber.Map{"code": global.RET_ERR_DB, "data": err.Error()})
			return
		}
		c.JSON(fiber.Map{"code": global.RET_OK, "data": ret})
		return
	}
	key := c.Params("key")
	if len(key) == 0 {
		ret, err := db.Keys(table)
		if err != nil {
			c.JSON(fiber.Map{"code": global.RET_ERR_DB, "data": err.Error()})
			return
		}
		c.JSON(fiber.Map{"code": global.RET_OK, "data": ret})
		return
	} else {
		ret, err := db.Get(table, key)
		if err != nil {
			c.JSON(fiber.Map{"code": global.RET_ERR_DB, "data": err.Error()})
			return
		} else {
			c.JSON(fiber.Map{"code": global.RET_OK, "data": string(ret)})
			return
		}
	}
}

func UpdateDB(c *fiber.Ctx) {
	store := sessions.Get(c)
	// user := store.Get(global.SESSION_KEY_USER)
	group := store.Get(global.SESSION_KEY_GROUP)
	if group == nil {
		c.JSON(fiber.Map{"code": global.RET_ERR_SESSION_INVALID,
			"data": "session invalid"})
		return
	}
	if int(group.(int64)) > int(global.GROUP_ADMIN) {
		c.JSON(fiber.Map{"code": global.RET_ERR_NO_RIGHT,
			"data": "no right to do this"})
		return
	}
	table := c.Params("table")
	key := c.Params("key")
	if len(table) == 0 {
		c.JSON(fiber.Map{"code": global.RET_ERR_URL_PARAM,
			"data": "request param err"})
		return
	}

	body := struct {
		Value string `json:"value"`
	}{}

	if errBp := c.BodyParser(&body); errBp != nil {
		c.JSON(fiber.Map{"code": global.RET_ERR_BODY_PARAM,
			"data": errBp.Error()})
		return
	}
	err := db.Set(table, key, body.Value)
	if err != nil {
		c.JSON(fiber.Map{"code": global.RET_ERR_DB,
			"data": err.Error()})
		return
	}
	c.JSON(fiber.Map{"code": global.RET_OK, "data": nil})
}

func DeleteDB(c *fiber.Ctx) {
	store := sessions.Get(c)
	// user := store.Get(global.SESSION_KEY_USER)
	group := store.Get(global.SESSION_KEY_GROUP)
	if group == nil {
		c.JSON(fiber.Map{"code": global.RET_ERR_SESSION_INVALID,
			"data": "session invalid"})
		return
	}
	if int(group.(int64)) > int(global.GROUP_ADMIN) {
		c.JSON(fiber.Map{"code": global.RET_ERR_NO_RIGHT,
			"data": "no right to do this"})
		return
	}
	table := c.Params("table")
	key := c.Params("key")
	if len(table) == 0 || len(key) == 0 {
		c.JSON(fiber.Map{"code": global.RET_ERR_URL_PARAM,
			"data": "request param err"})
		return
	}
	err := db.Delete(table, key)
	if err != nil {
		c.JSON(fiber.Map{"code": global.RET_ERR_DB,
			"data": err.Error()})
		return
	}
	c.JSON(fiber.Map{"code": global.RET_OK, "data": nil})
}

type ParamSvn struct {
	RepoUrl  string `json:"repourl"`
	Limit    string `json:"limit"`
	Period   string `json:"period"`
	Revision string `json:"revision"`
}

func CommitHistory(c *fiber.Ctx) {
	const timeout = time.Minute
	var paramSvn ParamSvn
	err := c.QueryParser(&paramSvn)
	if err != nil {
		c.JSON(fiber.Map{"code": global.RET_ERR_HTTP_QUERY,
			"data": err.Error()})
		return
	}
	if len(paramSvn.RepoUrl) == 0 {
		c.JSON(fiber.Map{"code": global.RET_ERR_URL_PARAM,
			"data": `can't find param repourl`})
		return
	}
	var cmd string
	if len(paramSvn.Revision) != 0 {
		cmd = fmt.Sprintf("svn log -r%s %s",
			paramSvn.Revision, paramSvn.RepoUrl)
	} else if len(paramSvn.Period) != 0 {
		cmd = fmt.Sprintf("svn log -r %s -q %s",
			paramSvn.Period, paramSvn.RepoUrl)
	} else {
		if len(paramSvn.Limit) == 0 {
			paramSvn.Limit = "10"
		}
		cmd = fmt.Sprintf("svn log -l %s -q %s",
			paramSvn.Limit, paramSvn.RepoUrl)
	}
	e, _, err := expect.Spawn(cmd, -1)
	defer e.Close()
	if err != nil {
		c.JSON(fiber.Map{"code": global.RET_ERR_SPAWN,
			"data": cmd})
		return
	}
	ret, _, _ := e.Expect(nil, timeout)
	c.JSON(fiber.Map{"code": global.RET_OK,
		"data": ret})
	return

}

func SubmitDep(c *fiber.Ctx) {
	subDepParam := struct {
		DepId float64 `json:"depid"`
		Force bool    `json:"force"`
	}{}
	if err := c.QueryParser(&subDepParam); err == nil {
		if subDepParam.DepId != 0 {
			depid := strconv.Itoa(int(subDepParam.DepId)) //float64
			depInfoInDB, _ := db.Get(global.BUCKET_OPS_DEPBIL, depid)
			depInfo := global.DepInfo{DepId: depid}
			json.Unmarshal(depInfoInDB, &depInfo)

			go cmd.SvnDep(depInfo, subDepParam.Force)
			c.JSON(fiber.Map{"code": global.RET_OK,
				"data": "request submitted"})
			return
		} else {
			c.JSON(fiber.Map{"code": global.RET_ERR_URL_PARAM,
				"data": `can't find param depid`})
			return
		}
	} else {
		c.JSON(fiber.Map{"code": global.RET_ERR_HTTP_QUERY,
			"data": err.Error()})
		return
	}
}

func ProgressList(c *fiber.Ctx) {
	ret := []string{}
	global.Depuuid2DepStatus.Range(func(k, v interface{}) bool {
		depInfo, _ := global.Depuuid2DepInfo.Load(k.(string))
		//start time
		ret = append(ret, strconv.FormatInt(
			depInfo.(global.DepInfo).StartTime, 10))
		ret = append(ret, depInfo.(global.DepInfo).DepId) //depid
		ret = append(ret, k.(string))                     //depuuid
		ret = append(ret, strconv.Itoa(v.(int)))          //dep status

		awakeTime, exist := global.DepTypeAwakeTime.Load(
			depInfo.(global.DepInfo).Type)
		if exist {
			ret = append(ret,
				strconv.FormatInt(awakeTime.(int64)-time.Now().Unix(), 10))
		} else {
			ret = append(ret, "0")
		}

		return true
	})
	c.JSON(fiber.Map{"code": global.RET_OK, "data": ret})
	return
}

func DeleteDep(c *fiber.Ctx) {
	const timeout = time.Minute
	subDepParam := struct {
		DepId string `json:"depid"`
	}{}
	if err := c.BodyParser(&subDepParam); err == nil {

	} else {
		c.JSON(fiber.Map{"code": global.RET_ERR_HTTP_QUERY,
			"data": err.Error()})
		return
	}
}

func BreakDep(c *fiber.Ctx) {
	const breakSecond = 120
	const renewSecond = 60
	//op = view(查看); break(打断); renew(续期); cancel(取消所有break操作,立即重启)
	breakDepParam := struct {
		DepUuid string `json:"depuuid"`
		Op      string `json:"op"`
	}{}
	err := c.QueryParser(&breakDepParam)
	if err != nil {
		c.JSON(fiber.Map{"code": global.RET_ERR_HTTP_QUERY,
			"data": err.Error()})
		return
	}
	if len(breakDepParam.DepUuid) == 0 {
		c.JSON(fiber.Map{"code": global.RET_ERR_URL_PARAM,
			"err": `can't find url param depuuid`})
		return
	}
	depStatus, existStatus := global.Depuuid2DepStatus.Load(breakDepParam.DepUuid)
	retcode := global.RET_OK
	var reterr string
	if existStatus {
		if depStatus.(int) >= global.DEP_STATUS_STOP_SERVER {
			retcode = global.RET_ERR_DEP_BREAK_TOO_LATE
			reterr = `it's too late to break this deployment`
		}
	} else {
		c.JSON(fiber.Map{"code": global.RET_ERR_DEP_EXPIRED,
			"err": `deployment status expired`})
		return
	}
	depInfo, existInfo := global.Depuuid2DepInfo.Load(breakDepParam.DepUuid)
	if !existInfo {
		c.JSON(fiber.Map{"code": global.RET_ERR_DEP_EXPIRED,
			"err": `deployment info expired`})
		return
	}
	nowSec := time.Now().Unix()
	awakeTime, loaded := global.DepTypeAwakeTime.Load(depInfo.(global.DepInfo).Type)
	if !loaded {
		if retcode == global.RET_OK {
			retcode = global.RET_ERR_FORCE_DEP_CANT_BREAK
			reterr = `can't break forced deployment`
		}
		awakeTime = (int64)(0)
	}

	tooLate := false
	if awakeTime.(int64) <= nowSec {
		tooLate = true
		if retcode == global.RET_OK {
			retcode = global.RET_ERR_DEP_BREAK_TOO_LATE
			reterr = `it's too late to break this deployment`
		}
	}
	if !tooLate {
		if breakDepParam.Op == "break" {
			//打断操作, 不足2分钟的补足2分钟, 超过的不作任何处理
			if awakeTime.(int64)-nowSec < breakSecond {
				awakeTime = nowSec + breakSecond
				global.DepTypeAwakeTime.Store(depInfo.(global.DepInfo).Type,
					awakeTime)
			}
		} else if breakDepParam.Op == "renew" {
			//续期操作
			awakeTime = awakeTime.(int64) + renewSecond
			global.DepTypeAwakeTime.Store(depInfo.(global.DepInfo).Type,
				awakeTime)
		} else if breakDepParam.Op == "cancel" {
			//取消操作
			global.DepTypeAwakeTime.Delete(depInfo.(global.DepInfo).Type)
		} else {
			//view only, nop
		}
	}
	if breakDepParam.Op == "view" {
		retcode = global.RET_OK
		reterr = ""
	}

	ret := []string{}

	//type
	ret = append(ret, depInfo.(global.DepInfo).Type)

	//countdown
	if tooLate {
		ret = append(ret, "0")
	} else {
		ret = append(ret, strconv.FormatInt(awakeTime.(int64)-nowSec, 10))
	}

	//status
	ret = append(ret, strconv.Itoa(depStatus.(int)))
	c.JSON(fiber.Map{"code": retcode,
		"data": ret, "err": reterr})

	return
}

func Chat(c *fiber.Ctx) {
	store := sessions.Get(c)
	user := store.Get(global.SESSION_KEY_USER)
	group := store.Get(global.SESSION_KEY_GROUP)
	if group == nil {
		c.JSON(fiber.Map{"code": global.RET_ERR_SESSION_INVALID,
			"data": "session invalid"})
		return
	}

	to := c.Params("to")
	chatInfo := struct {
		Message string `json:"message"`
	}{}
	if errBp := c.BodyParser(&chatInfo); errBp != nil {
		c.JSON(fiber.Map{"code": global.RET_ERR_BODY_PARAM,
			"data": errBp.Error()})
		return
	}
	mlog.Log(user, "chat to", to, "message:", chatInfo.Message)

	formatMsg := fmt.Sprintf(global.DINGDING_TEXT_MSG_PATTERN,
		"👻 "+chatInfo.Message)
	var errh error
	var reth []byte
	if to == "group" {
		reth, errh = utils.HttpPost(utils.GetDingdingWebhook(), formatMsg)
	} else if to == "dev" {
		reth, errh = utils.HttpPost(utils.GetDeveloperWebhook(), formatMsg)
	} else {
		if errBp := c.BodyParser(&chatInfo); errBp != nil {
			c.JSON(fiber.Map{"code": global.RET_ERR_URL_PARAM,
				"data": "err param 'to':" + to})
			return
		}
	}

	if errh != nil {
		c.JSON(fiber.Map{"code": global.RET_ERR_HTTP_REQUEST,
			"data": errh.Error()})
	} else {
		mlog.Log("chat webhook ret:\r\n", string(reth))
		c.JSON(fiber.Map{"code": global.RET_OK,
			"data": nil})
	}

}

func Test(c *fiber.Ctx) {
	const timeout = time.Minute
	e, _, err := expect.Spawn("node", -1)
	defer e.Close()
	if err != nil {
		c.JSON(fiber.Map{"code": global.RET_ERR_SPAWN,
			"data": "cmd"})
		return
	}
	checkedOutRE := regexp.MustCompile(">")
	ret, _, err2 := e.Expect(checkedOutRE, timeout)
	mlog.Log(ret, err2)
	e.Send("1+1\n")
	ret2, _, err3 := e.Expect(checkedOutRE, timeout)
	mlog.Log(ret2, err3)
	c.JSON(fiber.Map{"code": global.RET_OK,
		"data": ret})
}

var NewDB = UpdateDB
