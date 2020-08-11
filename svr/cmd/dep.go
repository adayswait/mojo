package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/adayswait/mojo/db"
	"github.com/adayswait/mojo/global"
	"github.com/adayswait/mojo/utils"
	"github.com/google/goexpect"
	"github.com/valyala/fasthttp"
	"golang.org/x/crypto/ssh"
	"regexp"
	"time"
)

var dingdingStr string

func init() {
	dingdingStr = "{" +
		`"msgtype":"markdown",` +
		`"markdown":{` +
		`"title":"更新提醒",` +
		`"text":"#### **%s**将于30秒后重启\n\n` +
		`> - [点击打断重启](http://10.1.1.248:8080)\n` +
		`###### %s [详情](http://10.1.1.248:8080)\n"` +
		"}" +
		"}"
}

func SvnDep(depid, deptype, coCmd, cuCmd, path string) {
	// depid := strconv.Itoa(int(subDepParam.DepId))
	global.ProgressMap.Store(depid, global.DEP_STATUS_NOT_START)
	ec, _, _ := expect.Spawn(cuCmd, -1)
	ec.Expect(regexp.MustCompile(`[\s\S]`), 2*time.Minute)
	global.ProgressMap.Store(depid, global.DEP_STATUS_CHECKOUT)
	e, _, err := expect.Spawn(coCmd, -1)
	defer e.Close()
	if err != nil {
		global.ProgressMap.Store(depid, global.DEP_STATUS_ERR_CHECKOUT)
		return
	}
	coStr := "Checked out revision"
	ret, matched, err := e.Expect(regexp.MustCompile(coStr),
		5*time.Minute)
	fmt.Println(coCmd, ret)
	if err != nil {
		global.ProgressMap.Store(depid, global.DEP_STATUS_ERR_CHECKOUT)
		return
	}
	if len(matched) == 0 {
		global.ProgressMap.Store(depid, global.DEP_STATUS_ERR_CHECKOUT)
		return
	}
	global.ProgressMap.Store(depid, global.DEP_STATUS_SYNC)
	depiniInDB, errd := db.Keys(global.BUCKET_OPS_DEPINI)
	if errd != nil {
		global.ProgressMap.Store(depid, global.DEP_STATUS_ERR_GETINI)
		fmt.Println("errd", errd)
		return
	}
	maciniInDB, errm := db.Keys(global.BUCKET_OPS_MACINI)
	if errm != nil {
		fmt.Println("errm", errm)
		global.ProgressMap.Store(depid, global.DEP_STATUS_ERR_GETINI)
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
		if idep[0] != deptype {
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
		fmt.Println(syncCmd)
		ers, _, _ := expect.SpawnWithArgs(syncCmd, -1)
		rsStr := "password:"
		retrs, matchedrs, errrs :=
			ers.Expect(regexp.MustCompile(rsStr), 5*time.Minute)
		if len(matchedrs) == 1 && errrs == nil {
			ers.Send(macini.Passwd + "\n")
			retrsc, matchedrsc, errrsc := ers.Expect(
				regexp.MustCompile("speedup is"), 5*time.Minute)
			if errrsc == nil && len(matchedrsc) == 1 {
				fmt.Println("sync succeed", retrsc)
			} else {
				fmt.Println("sync failed", retrsc, errrsc)
			}
		} else {
			fmt.Println(syncCmd, retrs, matchedrs, errrs)
		}

		sshClt, err := ssh.Dial("tcp",
			fmt.Sprintf("%s:%s", idep[2], macini.Port),
			&ssh.ClientConfig{
				User:            macini.User,
				Auth:            []ssh.AuthMethod{ssh.Password(macini.Passwd)},
				HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			})
		if err != nil {
			fmt.Println("ssh.Dial failed")
		}
		defer sshClt.Close()

		essh, _, errssh := expect.SpawnSSH(sshClt, time.Minute)
		if errssh != nil {
			fmt.Println(errssh)
		}
		defer essh.Close()
		logined := regexp.MustCompile("$")
		retlogin, _, _ := essh.Expect(logined, 10*time.Second)
		fmt.Println(retlogin)
		essh.Send(idep[3] + "/stop.sh\n")
		retstop, _, _ := essh.Expect(regexp.MustCompile("killed your"),
			10*time.Second)
		fmt.Println(retstop)
		essh.Send(idep[3] + "/start.sh\n")
		retstart, _, _ := essh.Expect(regexp.MustCompile("启动"), 10*time.Second)
		fmt.Println(retstart)
	}

	req := &fasthttp.Request{}
	req.SetRequestURI(utils.GetDingdingWebhook())
	markdown := fmt.Sprintf(dingdingStr, deptype,
		time.Now().Format("2006-01-02 15:04:05"))
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

	fmt.Println("result:\r\n", string(b))

	global.ProgressMap.Store(depid, global.DEP_STATUS_SUCCESS)

	return
}
