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
	"regexp"
	"time"
)

var dingdingStr string

func init() {
	dingdingStr = `{
		"msgtype":"actionCard",
		"actionCard":{
			"title":"内网%s更新提醒",
			"text":"#### 内网%s将于一分钟后重启
					###### 打断或查看需公司内网或vpn 
					###### %s",
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
	path := utils.GetPkgPath()
	if len(path) == 0 {
		path = "."
	}
	path = fmt.Sprintf("%s/%s/", path, depInfo.DepId)
	coCmd := fmt.Sprintf("svn export --force -%s %s %s",
		depInfo.Rversion, depInfo.RepoUrl, path)
	cuCmd := fmt.Sprintf("svn cleanup %s", path)
	depuuid := uuid.New().String()
	depInfo.StartTime = time.Now().Unix()
	global.Depuuid2DepInfo.Store(depuuid, depInfo)
	global.Depuuid2DepStatus.Store(depuuid, global.DEP_STATUS_NOT_START)
	ecu, _, _ := expect.Spawn(cuCmd, -1)
	defer ecu.Close()
	ecu.Expect(regexp.MustCompile("$"), 2*time.Minute)
	mlog.Log("svn cleanup passed")
	global.Depuuid2DepStatus.Store(depuuid, global.DEP_STATUS_CHECKOUT)
	eco, _, _ := expect.Spawn(coCmd, -1)
	defer eco.Close()
	coStr := "Exported revision"
	ret, matched, errcoe := eco.Expect(regexp.MustCompile(coStr),
		5*time.Minute)
	if errcoe != nil {
		mlog.Log(coCmd, ret)
		global.Depuuid2DepStatus.Store(depuuid, global.DEP_STATUS_ERR_CHECKOUT)
		return
	}
	if len(matched) == 0 {
		mlog.Log(coCmd, ret)
		global.Depuuid2DepStatus.Store(depuuid, global.DEP_STATUS_ERR_CHECKOUT)
		return
	}
	mlog.Log("svn export passed")

	ecoplan, _, _ := expect.Spawn(path+"/config/import_json_from_design.sh", -1)
	// defer ecoplan.Close() //todo here
	mlog.Log(path + "config/import_json_from_design.sh")
	retcoplan, matchedcoplan, errcoplan := ecoplan.Expect(
		regexp.MustCompile("import complete"),
		5*time.Minute)
	if errcoplan != nil || len(matchedcoplan) == 0 {
		mlog.Log("import json err", retcoplan, matchedcoplan, errcoplan)
		global.Depuuid2DepStatus.Store(depuuid, global.DEP_STATUS_ERR_CHECKOUT)
		return
	} else {
		mlog.Log("import json passed")
	}

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
		syncCmd := []string{
			"rsync",
			"-zarv",
			"--copy-links",
			fmt.Sprintf("--rsh=ssh -p %s", macini.Port),
			fmt.Sprintf("--exclude-from=%s", utils.GetExcludeFrom()),
			path,
			fmt.Sprintf("%s@%s:%s", macini.User, idep[2], idep[3]),
		}
		ers, _, _ := expect.SpawnWithArgs(syncCmd, -1)
		rsStr := "password:"
		retrs, matchedrs, errrs :=
			ers.Expect(regexp.MustCompile(rsStr), 5*time.Minute)
		if len(matchedrs) == 1 && errrs == nil {
			ers.Send(macini.Passwd + "\n")
			retrsc, matchedrsc, errrsc := ers.Expect(
				regexp.MustCompile("speedup is"), 5*time.Minute)
			if errrsc == nil && len(matchedrsc) == 1 {
				mlog.Log("sync succeed", idep[2], idep[3])
			} else {
				mlog.Log("sync failed", idep[2], idep[3], retrsc, errrsc)
				continue
			}
		} else {
			mlog.Log("sync failed", syncCmd, retrs, matchedrs, errrs)
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
		retstart, _, estart := essh.Expect(regexp.MustCompile("启动"), 10*time.Second)
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
