package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/adayswait/mojo/db"
	"github.com/adayswait/mojo/global"
	"github.com/adayswait/mojo/utils"
	"github.com/google/goexpect"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"golang.org/x/crypto/ssh"
	"regexp"
	"time"
)

var dingdingStr string

func init() {
	dingdingStr = "{" +
		`"msgtype":"actionCard",` +
		`"actionCard":{` +
		`"title":"更新提醒",` +
		`"text":"#### **%s**将于一分钟后重启\n ###### %s",` +
		`"btnOrientation": "1",` +
		`"btns": [` +
		"{" +
		`"title": "打断重启", ` +
		`"actionURL": "%s"` +
		"}," +
		"{" +
		`"title": "查看发布状态", ` +
		`"actionURL": "%s"` +
		"}]" +
		"}" +
		"}"
}

func SvnDep(depInfo global.DepInfo, force bool) {
	path := utils.GetRepoPath()
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
	fmt.Println("svn cleanup passed")
	global.Depuuid2DepStatus.Store(depuuid, global.DEP_STATUS_CHECKOUT)
	eco, _, _ := expect.Spawn(coCmd, -1)
	defer eco.Close()
	coStr := "Exported revision"
	ret, matched, errcoe := eco.Expect(regexp.MustCompile(coStr),
		5*time.Minute)
	if errcoe != nil {
		fmt.Println(coCmd, ret)
		global.Depuuid2DepStatus.Store(depuuid, global.DEP_STATUS_ERR_CHECKOUT)
		return
	}
	if len(matched) == 0 {
		fmt.Println(coCmd, ret)
		global.Depuuid2DepStatus.Store(depuuid, global.DEP_STATUS_ERR_CHECKOUT)
		return
	}
	fmt.Println("svn export passed")

	ecoplan, _, _ := expect.Spawn(path+"/config/import_json_from_design.sh", -1)
	// defer ecoplan.Close() //todo here
	fmt.Println(path + "config/import_json_from_design.sh")
	retcoplan, matchedcoplan, errcoplan := ecoplan.Expect(
		regexp.MustCompile("import complete"),
		5*time.Minute)
	if errcoplan != nil || len(matchedcoplan) == 0 {
		fmt.Println("import json err", retcoplan, matchedcoplan, errcoplan)
		global.Depuuid2DepStatus.Store(depuuid, global.DEP_STATUS_ERR_CHECKOUT)
		return
	} else {
		fmt.Println("import json passed")
	}

	if !force {
		//发布更新通知
		req := &fasthttp.Request{}
		req.SetRequestURI(utils.GetDingdingWebhook())

		markdown := fmt.Sprintf(dingdingStr, depInfo.Type,
			time.Now().Format("2006-01-02 15:04:05"),
			fmt.Sprintf("http://%s:%d/api/visitor/breakdep?depuuid=%s",
				utils.GetServerHost(),
				utils.GetListeningPort(),
				depuuid),
			fmt.Sprintf("http://10.1.1.248:8080/#/visitor/viewdep?depuuid=%s", depuuid))
		req.SetBody([]byte(markdown))

		// 默认是application/x-www-form-urlencoded
		req.Header.SetContentType("application/json")
		req.Header.SetMethod("POST")

		resp := &fasthttp.Response{}

		client := &fasthttp.Client{}
		if err := client.Do(req, resp); err != nil {
			fmt.Println("请求失败:", err.Error())
			return
		}
		b := resp.Body()

		fmt.Println("dingding webhook ret:\r\n", string(b))
		global.Depuuid2DepStatus.Store(depuuid, global.DEP_STATUS_SLEEP)
		global.DepTypeAwakeTime.Store(depInfo.Type, time.Now().Unix()+10)
		time.Sleep(time.Second * 10)
		for {
			breakTime, exist := global.DepTypeAwakeTime.Load(depInfo.Type)
			if exist {
				gapTime := breakTime.(int64) - time.Now().Unix()
				if gapTime <= 0 {
					break
				}
				time.Sleep(time.Duration(gapTime) * time.Second)
			} else {
				break
			}
		}
	}

	depiniInDB, errd := db.Keys(global.BUCKET_OPS_DEPINI)
	if errd != nil {
		global.Depuuid2DepStatus.Store(depuuid, global.DEP_STATUS_ERR_GETINI)
		fmt.Println("errd", errd)
		return
	}
	maciniInDB, errm := db.Keys(global.BUCKET_OPS_MACINI)
	if errm != nil {
		fmt.Println("errm", errm)
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
				fmt.Println("sync succeed", idep[2], idep[3])
			} else {
				fmt.Println("sync failed", idep[2], idep[3], retrsc, errrsc)
				continue
			}
		} else {
			fmt.Println("sync failed", syncCmd, retrs, matchedrs, errrs)
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
			fmt.Println("ssh dial failed", errdial)
			continue
		}
		defer sshClt.Close()
		fmt.Println("ssh dial passed")

		essh, _, errssh := expect.SpawnSSH(sshClt, time.Minute)
		if errssh != nil {
			fmt.Println(errssh)
			continue
		}
		defer essh.Close()
		logined := regexp.MustCompile("$")
		retlogin, _, elogin := essh.Expect(logined, 10*time.Second)
		if elogin != nil {
			fmt.Println("ssh login failed", elogin, retlogin)
			continue
		}
		fmt.Println("ssh login passed")
		essh.Send(idep[3] + "/stop.sh\n")
		retstop, _, estop := essh.Expect(regexp.MustCompile("$"),
			10*time.Second)
		if estop != nil {
			fmt.Println("stop old service failed", estop, retstop)
			continue
		} else {
			fmt.Println("stop old service passed")
		}
		essh.Send(idep[3] + "/start.sh\n")
		retstart, _, estart := essh.Expect(regexp.MustCompile("启动"), 10*time.Second)
		if estart != nil {
			fmt.Println("start new service failed", estart, retstart)
			continue
		} else {
			fmt.Println("start new service passed")
		}
	}
	fmt.Println("dep all done")
	global.Depuuid2DepStatus.Store(depuuid, global.DEP_STATUS_SUCCESS)

	return
}
