package handler

import (
	"encoding/json"
	"github.com/adayswait/mojo/global"
	"github.com/adayswait/mojo/mlog"
	"github.com/adayswait/mojo/utils"
	"github.com/gofiber/fiber"
	"strconv"
	"time"
)

type MailParam struct {
	Opt          string `json:"opt"`
	ActiveTime   string `json:"activetime"`
	Sender       string `json:"sender"`
	Title        string `json:"title"`
	RegEndTime   string `json:"regendtime"`
	User         string `json:"user"`
	Gmail_file   string `json:"gmail_file"`
	RegStartTime string `json:"regstarttime"`
	Refresh      string `json:"refresh"`
	MailType     string `json:"mailtype"`
	Content      string `json:"content"`
	Addition     string `json:"addition"`
	Switch       string `json:"switch"`
	DeadTime     string `json:"deadtime"`
	SendTime     string `json:"sendtime"`
	SwitchKey    string `json:"switch_key"`
	Attachment   string `json:"attachment"`
}

func SplanMail(c *fiber.Ctx) {
	store := sessions.Get(c)
	// user := store.Get(global.SESSION_KEY_USER)
	group := store.Get(global.SESSION_KEY_GROUP)
	if group == nil {
		c.JSON(fiber.Map{"code": global.RET_ERR_SESSION_INVALID,
			"data": "session invalid"})
		return
	}
	if int(group.(int64)) > int(global.GROUP_USER) {
		c.JSON(fiber.Map{"code": global.RET_ERR_NO_RIGHT,
			"data": "no right to do this"})
		return
	}
	body := MailParam{}
	if errBp := c.BodyParser(&body); errBp != nil {
		c.JSON(fiber.Map{"code": global.RET_ERR_BODY_PARAM,
			"data": errBp.Error()})
		return
	}
	body.SendTime = strconv.FormatInt(time.Now().Unix(), 10)
	body.Opt = "add_global_mail"
	mlog.Log("content:", string(body.Content))
	bodyData, _ := json.Marshal(body)

	reth, errh := utils.HttpPost("http://10.1.1.43:21010/global_mail_proc",
		string(bodyData))
	if errh != nil {
		mlog.Log("splan add global mail err:\r\n", errh.Error())
	} else {
		mlog.Log("splan add global mail ret:\r\n", string(reth))
	}

	c.JSON(fiber.Map{"code": global.RET_OK, "data": nil})
}
