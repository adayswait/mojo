package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/adayswait/mojo/db"
	"github.com/adayswait/mojo/global"
	"github.com/adayswait/mojo/mlog"
	"github.com/adayswait/mojo/utils"
	"github.com/google/goexpect"
	"github.com/google/uuid"
	"golang.org/x/crypto/ssh"
	"os/exec"
	"regexp"
	"time"
)

var dingdingStr string

func init() {
	dingdingStr = `{
		"msgtype":"actionCard",
		"actionCard":{
			"title":"内网%s更新提醒",
			"text":"#### 内网%s将于一分钟后重启\n ###### 打断或查看需公司内网或vpn\n ###### %s",
			"btnOrientation": "1",
			"btns": [
				{
					"title": "打断重启",
					"actionURL": "%s"
				},
				{
					"title": "查看发布状态",
					"actionURL": "%s"
				}
			]
		}
	}`
}

func SvnDep(depInfo global.DepInfo, force bool) {

	depuuid := uuid.New().String()
	depInfo.StartTime = time.Now().Unix()
	global.Depuuid2DepInfo.Store(depuuid, depInfo)
	global.Depuuid2DepStatus.Store(depuuid, global.DEP_STATUS_NOT_START)

	path := utils.GetPkgPath()
	if len(path) == 0 {
		path = "."
	}
	path = fmt.Sprintf("%s/%s/", path, depInfo.DepId)

	cleanupCmd := exec.Command("rm", "-rf", path)
	cleanupErr := cleanupCmd.Run()
	if cleanupErr != nil {
		mlog.Log("rm -rf", path, "err:", cleanupErr)
		return
	}
	// revision := "-r" + depInfo.Revision
	// svnExportCmd := exec.Command("svn", "export", "--force",
	// 	revision, depInfo.RepoUrl, path)
	// fmt.Println(svnExportCmd)
	// svnExportErr := svnExportCmd.Run()
	// if svnExportErr != nil {
	// 	mlog.Log("svn export err", svnExportErr, svnExportCmd)
	// 	return
	// }
	coCmd := fmt.Sprintf("svn export --force -r%s %s %s",
		depInfo.Revision, depInfo.RepoUrl, path)
	expectCo, _, _ := expect.Spawn(coCmd, -1)
	defer expectCo.Close()
	retCo, matchedCo, errCo := expectCo.Expect(
		regexp.MustCompile("Exported revision"),
		5*time.Minute)
	if errCo != nil {
		mlog.Log(coCmd, retCo, matchedCo, errCo)
		global.Depuuid2DepStatus.Store(depuuid, global.DEP_STATUS_ERR_CHECKOUT)
		return
	}
	if len(matchedCo) == 0 {
		mlog.Log(coCmd, retCo, matchedCo)
		global.Depuuid2DepStatus.Store(depuuid, global.DEP_STATUS_ERR_CHECKOUT)
		return
	}
	mlog.Log("svn export passed")

	runImportCmd := exec.Command(path + "/config/import_json_from_design.sh")
	runImportErr := runImportCmd.Run()
	if runImportErr != nil {
		mlog.Log(path+"/config/import_json_from_design.sh err:", runImportErr)
		return
	}
	mlog.Log("import json passed")

	if !force {
		//发布更新通知
		markdown := fmt.Sprintf(dingdingStr, depInfo.Type, depInfo.Type,
			time.Now().Format("2006-01-02 15:04:05"),
			fmt.Sprintf("%s/#/visitor/breakdep?depuuid=%s&op=break",
				utils.GetWebDomain(), depuuid),
			fmt.Sprintf("%s/#/visitor/breakdep?depuuid=%s&op=view",
				utils.GetWebDomain(), depuuid))

		reth, errh := utils.HttpPost(utils.GetDingdingWebhook(), markdown)
		if errh != nil {
			mlog.Log("deploye webhook err:\r\n", errh.Error())
		} else {
			mlog.Log("deploye webhook ret:\r\n", string(reth))
		}

		global.Depuuid2DepStatus.Store(depuuid, global.DEP_STATUS_SLEEP)
		global.DepTypeAwakeTime.Store(depInfo.Type, time.Now().Unix()+60)
		for {
			breakTime, exist := global.DepTypeAwakeTime.Load(depInfo.Type)
			if exist {
				gapTime := breakTime.(int64) - time.Now().Unix()
				if gapTime <= 0 {
					break
				}
				time.Sleep(time.Second)
			} else {
				break
			}
		}
	}

	depiniInDB, errd := db.Keys(global.BUCKET_OPS_DEPINI)
	if errd != nil {
		global.Depuuid2DepStatus.Store(depuuid, global.DEP_STATUS_ERR_GETINI)
		mlog.Log("errd", errd)
		return
	}
	maciniInDB, errm := db.Keys(global.BUCKET_OPS_MACINI)
	if errm != nil {
		mlog.Log("errm", errm)
		global.Depuuid2DepStatus.Store(depuuid, global.DEP_STATUS_ERR_GETINI)
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

	global.Depuuid2DepStatus.Store(depuuid, global.DEP_STATUS_SYNC)
	for i := 1; i < len(depiniInDB); i += 2 {
		var idep []string
		e := json.Unmarshal([]byte(depiniInDB[i]), &idep)
		if e != nil {
			continue
		}
		if idep[0] != depInfo.Type {
			continue
		}
		inList := false
		for c := 0; c < len(depInfo.List); c++ {
			if idep[1] == depInfo.List[c] {
				inList = true
				break
			}
		}
		if !inList && len(depInfo.List) != 0 {
			continue
		}
		macini, exist := macIni[idep[2]]
		if exist == false {
			continue
		}
		rsyncErr := Rsync(path, idep[3], macini.User, macini.Passwd,
			idep[2], macini.Port,
			[]string{"-zarv", "--copy-links",
				fmt.Sprintf("--exclude-from=%s", utils.GetExcludeFrom())},
			5*time.Minute)
		if rsyncErr != nil {
			continue
		}

		sshClt, errdial := ssh.Dial("tcp",
			fmt.Sprintf("%s:%s", idep[2], macini.Port),
			&ssh.ClientConfig{
				User:            macini.User,
				Auth:            []ssh.AuthMethod{ssh.Password(macini.Passwd)},
				HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			})
		if errdial != nil {
			mlog.Log("ssh dial failed", errdial)
			continue
		}
		defer sshClt.Close()
		mlog.Log("ssh dial passed")

		essh, _, errssh := expect.SpawnSSH(sshClt, time.Minute)
		if errssh != nil {
			mlog.Log(errssh)
			continue
		}
		defer essh.Close()
		logined := regexp.MustCompile("$")
		retlogin, _, elogin := essh.Expect(logined, 10*time.Second)
		if elogin != nil {
			mlog.Log("ssh login failed", elogin, retlogin)
			continue
		}
		mlog.Log("ssh login passed")
		essh.Send(idep[3] + "/stop.sh\n")
		retstop, _, estop := essh.Expect(regexp.MustCompile("$"),
			10*time.Second)
		if estop != nil {
			mlog.Log("stop old service failed", estop, retstop)
			continue
		} else {
			mlog.Log("stop old service passed")
		}
		essh.Send(idep[3] + "/start.sh\n")
		retstart, _, estart := essh.Expect(regexp.MustCompile("启动"),
			10*time.Second)
		if estart != nil {
			mlog.Log("start new service failed", estart, retstart)
			continue
		} else {
			mlog.Log("start new service passed")
		}
	}
	mlog.Log("dep all done")
	global.Depuuid2DepStatus.Store(depuuid, global.DEP_STATUS_SUCCESS)

	return
}
