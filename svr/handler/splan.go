package handler

import (
	"encoding/json"
	"fmt"
	"github.com/adayswait/mojo/cmd"
	"github.com/adayswait/mojo/db"
	"github.com/adayswait/mojo/global"
	"github.com/adayswait/mojo/mlog"
	"github.com/adayswait/mojo/utils"
	"github.com/gofiber/fiber"
	"os/exec"
	"strconv"
	"strings"
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

func SplanUpdateConfig(c *fiber.Ctx) {
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
	module := c.Params("module")

	dbRet, errRepo := db.Get(global.BUCKET_OPS_DEVINI, module)
	if errRepo != nil {
		c.JSON(fiber.Map{"code": global.RET_ERR_DB, "data": errRepo.Error()})
		return
	}
	repoInfo := struct {
		URL string `json:"url"`
	}{}
	json.Unmarshal(dbRet, &repoInfo)

	dirPath := fmt.Sprintf("%s/hotupdate/", utils.GetPkgPath())
	exePath := fmt.Sprintf("%s/hotupdate/import_json_from_design.sh",
		utils.GetPkgPath())

	rmLastDirCmd := exec.Command("rm", "-rf", dirPath)
	rmLastDirErr := rmLastDirCmd.Run()
	if rmLastDirErr != nil {
		mlog.Log("rm -rf", dirPath, "err:", rmLastDirErr)
		c.JSON(fiber.Map{"code": global.RET_ERR_OS_EXEC_CMD_RUN,
			"data": rmLastDirErr.Error()})
		return
	}
	mkNewDirCmd := exec.Command("mkdir", dirPath)
	mkNewDirErr := mkNewDirCmd.Run()
	if mkNewDirErr != nil {
		mlog.Log("mkdir", dirPath, "err:", mkNewDirErr)
		c.JSON(fiber.Map{"code": global.RET_ERR_OS_EXEC_CMD_RUN,
			"data": mkNewDirErr.Error()})
		return
	}

	coConfigCmd := exec.Command("svn", "export", "--force",
		repoInfo.URL+"/config/import_json_from_design.sh", exePath)
	coConfigErr := coConfigCmd.Run()
	if coConfigErr != nil {
		mlog.Log(fmt.Sprintf("svn export --force %s %s",
			repoInfo.URL+"/config/import_json_from_design.sh", exePath),
			"err:", coConfigErr)
		c.JSON(fiber.Map{"code": global.RET_ERR_OS_EXEC_CMD_RUN,
			"data": coConfigErr.Error()})
		return
	}
	if strings.Contains(module, "online") {
		coConfigCmd2 := exec.Command("svn", "export", "--force",
			repoInfo.URL+"/config/new_json_files.txt",
			fmt.Sprintf("%s/hotupdate/new_json_files.txt",
				utils.GetPkgPath()))
		coConfigCmd2.Run()
	}

	runImportCmd := exec.Command(exePath)
	runImportErr := runImportCmd.Run()
	if runImportErr != nil {
		mlog.Log("run", exePath, "err:", runImportErr)
		c.JSON(fiber.Map{"code": global.RET_ERR_OS_EXEC_CMD_RUN,
			"data": runImportErr.Error()})
		return
	}

	//开启rsync
	depiniInDB, errd := db.Keys(global.BUCKET_OPS_DEPINI)
	if errd != nil {
		mlog.Log("errd", errd)
		c.JSON(fiber.Map{"code": global.RET_ERR_DB,
			"data": errd.Error()})
		return
	}
	maciniInDB, errm := db.Keys(global.BUCKET_OPS_MACINI)
	if errm != nil {
		mlog.Log("errm", errm)
		c.JSON(fiber.Map{"code": global.RET_ERR_DB,
			"data": errm.Error()})
		return
	}
	macIni := make(map[string]global.SSHinfo)
	for i := 1; i < len(maciniInDB); i += 2 {
		var imac []string
		e := json.Unmarshal([]byte(maciniInDB[i]), &imac)
		if e != nil {
			continue
		}
		_, exist := macIni[imac[0]]
		if exist == false {
			macIni[imac[0]] = global.SSHinfo{
				Port:   imac[1],
				User:   imac[2],
				Passwd: imac[3],
			}
		}
	}
	for i := 1; i < len(depiniInDB); i += 2 {
		var idep []string
		e := json.Unmarshal([]byte(depiniInDB[i]), &idep)
		if e != nil {
			continue
		}
		if idep[0] != module {
			continue
		}
		macini, exist := macIni[idep[2]]
		if exist == false {
			continue
		}
		rsyncErr := cmd.Rsync(dirPath, idep[3]+"/config/",
			macini.User, macini.Passwd,
			idep[2], macini.Port,
			[]string{"-zarv", "--copy-links",
				fmt.Sprintf("--exclude-from=%s", utils.GetExcludeFrom())},
			5*time.Minute)
		if rsyncErr != nil {
			continue
		}
	}

	c.JSON(fiber.Map{"code": global.RET_OK, "data": nil})
}
