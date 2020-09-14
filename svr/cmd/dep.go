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

/*
部署发布需要四个脚本文件
bd.sh(可选) 拉取代码前执行 before deploy
ad.sh(可选) 拉取代码后执行 after deploy
br.sh(必需) 拉取代码前执行 before release 执行成功需在脚本末尾输出mojobrok
ar.sh(必需) 拉取代码后执行 after release  执行成功需在脚本末尾输出mojoarok
*/
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

	//deploy前置任务,bd means before deploy
	beforeDeployShellPath := path + "bd.sh"
	if utils.IsPathExist(beforeDeployShellPath) &&
		!utils.IsDir(beforeDeployShellPath) {
		beforeDeployCmd := exec.Command("bash", beforeDeployShellPath)
		beforeDeployErr := beforeDeployCmd.Run()
		if beforeDeployErr != nil {
			mlog.Log("before deploy cmd run failed :", beforeDeployShellPath,
				"error", beforeDeployErr)
		} else {
			mlog.Log("before deploy cmd run succeed")
		}
	} else {
		mlog.Log("before deploy cmd not found")
	}

	//同步代码
	cleanupCmd := exec.Command("rm", "-rf", path)
	cleanupErr := cleanupCmd.Run()
	if cleanupErr != nil {
		mlog.Log("rm -rf", path, "err:", cleanupErr)
		return
	}
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

	//deploy后置任务,ad means after deploy
	afterDeployShellPath := path + "ad.sh"
	if utils.IsPathExist(afterDeployShellPath) &&
		!utils.IsDir(afterDeployShellPath) {
		afterDeployCmd := exec.Command("bash", afterDeployShellPath)
		afterDeployErr := afterDeployCmd.Run()
		if afterDeployErr != nil {
			mlog.Log("after deploy cmd run failed :", afterDeployShellPath,
				"error", afterDeployErr)
		} else {
			mlog.Log("after deploy cmd run succeed")
		}
	} else {
		mlog.Log("after deploy cmd not found")
	}

	if !force {
		//发布更新通知
		markdown := fmt.Sprintf(dingdingStr, depInfo.Type, depInfo.Type,
			time.Now().Format("2006-01-02 15:04:05"),
			fmt.Sprintf("%s/#/visitor/breakdep?depuuid=%s&op=break",
				utils.GetWebDomain()[0], depuuid),
			fmt.Sprintf("%s/#/visitor/breakdep?depuuid=%s&op=view",
				utils.GetWebDomain()[0], depuuid))

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

		essh.Send("bash " + idep[3] + "/br.sh\n")
		retbr, _, ebr := essh.Expect(regexp.MustCompile("mojobrok"),
			10*time.Second)
		if ebr != nil {
			//before release脚本执行错误不中断后续流程
			//因为首次发布的时候可能没有br.sh
			//todo 首先传输br.sh到指定位置
			mlog.Log("before release cmd exec error", ebr, retbr)
		} else {
			mlog.Log("before release cmd exec succeed")
		}

		rsyncErr := Rsync(path, idep[3], macini.User, macini.Passwd,
			idep[2], macini.Port,
			[]string{"-zarv", "--copy-links",
				fmt.Sprintf("--exclude-from=%s", utils.GetExcludeFrom())},
			5*time.Minute)
		if rsyncErr != nil {
			continue
		}

		var arcmd string
		if len(idep) > 4 && len(idep[4]) != 0 {
			arcmd = "bash " + idep[3] + "/ar.sh " + idep[4] + "\n"
		} else {
			arcmd = "bash " + idep[3] + "/ar.sh\n"
		}
		essh.Send(arcmd)
		retar, _, ear := essh.Expect(regexp.MustCompile("mojoarok"),
			10*time.Second)
		if ear != nil {
			mlog.Log("after release cmd exec error", arcmd, ear, retar)
			continue
		} else {
			mlog.Log("after release cmd exec succeed", arcmd)
		}
	}
	mlog.Log(depInfo.DepId, "dep all done")
	global.Depuuid2DepStatus.Store(depuuid, global.DEP_STATUS_SUCCESS)

	return
}
