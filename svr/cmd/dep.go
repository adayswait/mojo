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
	"strings"
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
			mlog.Errorf("before deploy exec cmd failed, path : %s, error : %s",
				beforeDeployShellPath, beforeDeployErr)
		} else {
			mlog.Info("before deploy exec cmd succeed")
		}
	} else {
		mlog.Warning("before deploy bd.sh not found")
	}

	//同步代码
	cleanupCmd := exec.Command("rm", "-rf", path)
	cleanupErr := cleanupCmd.Run()
	if cleanupErr != nil {
		mlog.Errorf("rm -rf %s failed, err : %s", path, cleanupErr)
		return
	}
	coCmd := fmt.Sprintf("svn export --force -r%s %s %s",
		depInfo.Revision, depInfo.RepoUrl, path)
	expectCo, _, _ := expect.Spawn(coCmd, -1)
	// defer expectCo.Close()
	retCo, matchedCo, errCo := expectCo.Expect(
		regexp.MustCompile("Exported revision"),
		5*time.Minute)
	if errCo != nil || len(matchedCo) == 0 {
		mlog.Errorf("expect run cmd : %s failed, ret : %s, match : %s, err : %s",
			coCmd, retCo, matchedCo, errCo)
		global.Depuuid2DepStatus.Store(depuuid, global.DEP_STATUS_ERR_CHECKOUT)
		return
	}
	mlog.Infof("expect run cmd : %s succeed", coCmd)

	//deploy后置任务,ad means after deploy
	afterDeployShellPath := path + "ad.sh"
	if utils.IsPathExist(afterDeployShellPath) &&
		!utils.IsDir(afterDeployShellPath) {
		afterDeployCmd := exec.Command("bash", afterDeployShellPath)
		afterDeployErr := afterDeployCmd.Run()
		if afterDeployErr != nil {
			mlog.Errorf("before deploy exec cmd failed, path : %s, error : %s",
				afterDeployShellPath, afterDeployErr)
		} else {
			mlog.Info("after deploy exec cmd succeed")
		}
	} else {
		mlog.Warning("after deploy ad.sh not found")
	}

	if !force {
		//发布更新通知
		markdown := fmt.Sprintf(dingdingStr, depInfo.Type, depInfo.Type,
			time.Now().Format("2006-01-02 15:04:05"),
			fmt.Sprintf("%s/#/visitor/breakdep?depuuid=%s&op=break",
				strings.Split(utils.GetWebDomain(), ",")[0], depuuid),
			fmt.Sprintf("%s/#/visitor/breakdep?depuuid=%s&op=view",
				strings.Split(utils.GetWebDomain(), ",")[0], depuuid))

		reth, errh := utils.HttpPost(utils.GetDingdingWebhook(), markdown)
		if errh != nil {
			mlog.Warningf("deploye webhook err : %s", errh.Error())
		} else {
			mlog.Infof("deploye webhook ret : %s", string(reth))
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
		mlog.Errorf("iterating over all the keys in bucket : %s, err : %s",
			global.DEP_STATUS_ERR_GETINI, errd)
		return
	}
	maciniInDB, errm := db.Keys(global.BUCKET_OPS_MACINI)
	if errm != nil {
		mlog.Errorf("iterating over all the keys in bucket : %s, err : %s",
			global.BUCKET_OPS_MACINI, errm)
		global.Depuuid2DepStatus.Store(depuuid, global.DEP_STATUS_ERR_GETINI)
		return
	}
	macIni := make(map[string]global.SSHinfo)
	for i := 1; i < len(maciniInDB); i += 2 {
		var imac []string
		e := json.Unmarshal([]byte(maciniInDB[i]), &imac)
		if e != nil {
			mlog.Errorf("json.Unmarshal : %s, err:%v", depiniInDB[i], e)
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

	var totalN = 0
	var failedN = 0

	global.Depuuid2DepStatus.Store(depuuid, global.DEP_STATUS_SYNC)
	for i := 1; i < len(depiniInDB); i += 2 {
		var idep []string
		e := json.Unmarshal([]byte(depiniInDB[i]), &idep)
		if e != nil {
			mlog.Errorf("json.Unmarshal : %s, err:%v", depiniInDB[i], e)
			continue
		}

		//过滤掉非本次发布类型的服务
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

		//过滤掉不再发布列表中的服务
		if !inList && len(depInfo.List) != 0 {
			continue
		}
		totalN += 1
		macini, exist := macIni[idep[2]]
		if exist == false {
			mlog.Errorf("macIni[%s] doesn't exist", idep[2])
			failedN += 1
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
			mlog.Errorf("ssh dial %s@%s:%s failed, err : %s",
				macini.User, idep[2], macini.Port, errdial)
			failedN += 1
			continue
		}
		defer sshClt.Close()
		mlog.Infof("ssh dial %s@%s:%s succeed",
			macini.User, idep[2], macini.Port)
		essh, _, errssh := expect.SpawnSSH(sshClt, time.Minute)
		if errssh != nil {
			failedN += 1
			mlog.Errorf("expect.SpawnSSH failed, err : %s", errssh)
			continue
		}
		defer essh.Close()
		logined := regexp.MustCompile("$")
		retlogin, _, elogin := essh.Expect(logined, 10*time.Second)
		if elogin != nil {
			mlog.Errorf("expect ssh login failed, ret : %s, err : %s",
				retlogin, elogin)
			failedN += 1
			continue
		}
		mlog.Infof("expect ssh login succeed, ret:%s", retlogin)

		mkdircmd := "mkdir -p " + idep[3] + " && echo mojomkdirok\n"
		essh.Send(mkdircmd)
		retmkdir, matchedmkdir, emkdir := essh.Expect(regexp.MustCompile("mojomkdirok"),
			10*time.Second)
		if emkdir != nil {
			mlog.Errorf("before rsync mkdir failed, "+
				"cmd:%s, ret:%s, match:%s, err:%s",
				mkdircmd, retmkdir, matchedmkdir, emkdir)
			failedN += 1
			continue
		} else {
			mlog.Info("before rsync mkdir succeed")
		}

		essh.Send("bash " + idep[3] + "/br.sh  && echo mojobrok\n")
		retbr, matchedbr, ebr := essh.Expect(regexp.MustCompile("mojobrok"),
			10*time.Second)
		if ebr != nil {
			//before release脚本执行错误不中断后续流程
			//因为首次发布的时候可能没有br.sh
			//todo 首先传输br.sh到指定位置
			mlog.Errorf("before release exec cmd failed, "+
				"path:%s, ret:%s, match:%s, err:%s",
				idep[3]+"/br.sh", retbr, matchedbr, ebr)
		} else {
			mlog.Info("before release exec cmd succeed")
		}

		rsyncErr := Rsync(path, idep[3], macini.User, macini.Passwd,
			idep[2], macini.Port,
			[]string{"-zarv", "--copy-links",
				fmt.Sprintf("--exclude-from=%s", utils.GetExcludeFrom())},
			5*time.Minute)
		if rsyncErr != nil {
			mlog.Errorf("rsync err:%v", rsyncErr)
			failedN += 1
			continue
		}

		var arcmd string
		if len(idep) > 4 && len(idep[4]) != 0 {
			arcmd = "bash " + idep[3] + "/ar.sh " + idep[4] + " && echo mojoarok\n"
		} else {
			arcmd = "bash " + idep[3] + "/ar.sh && echo mojoarok\n"
		}
		essh.Send(arcmd)
		retar, matchedar, ear := essh.Expect(regexp.MustCompile("mojoarok"),
			10*time.Second)
		if ear != nil {
			mlog.Errorf("before release exec cmd failed,"+
				"path:%s, ret:%s, match:%s, err:%s",
				arcmd, retar, matchedar, ear)
			failedN += 1
			continue
		} else {
			mlog.Info("after release exec cmd succeed")
		}
	}
	mlog.Infof("deploy id:%s run over,total:%d, failed:%d",
		depInfo.DepId, totalN, failedN)
	if totalN != 0 && failedN == 0 {
		global.Depuuid2DepStatus.Store(depuuid, global.DEP_STATUS_SUCCESS)
	} else {
		global.Depuuid2DepStatus.Store(depuuid, global.DEP_STATUS_ERR)
	}

	return
}
