package handler

import (
	"encoding/json"
	"fmt"
	"github.com/adayswait/mojo/cmd"
	"github.com/adayswait/mojo/db"
	"github.com/adayswait/mojo/global"
	"github.com/adayswait/mojo/mlog"
	"github.com/adayswait/mojo/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/goexpect"
	"golang.org/x/crypto/ssh"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var onlineHotUpdate = "pveBoss.json;pveDrop.json;pveEvent.json;festivalCarnivalReward.json;item.json;petEgg.json;petEvo.json;petFight.json;dropUnit.json;petLearnSkill.json;prizePool.json;eventManage.json;pet.json;exp.json;petToken.json;petSkillLvUp.json;lottery.json;activityTask.json;mapMonster.json;evoType.json;obtainGaiya.json;timeTravel.json;extremeChallenge.json;itemElemType.json;extremeLimit.json;fightSkills.json;shop.json;battleshop.json;laddershop.json;mysteriousShop.json;startrekshop.json;shopmanage.json;buyEnergy.json;releaseShop.json;expMechine.json;teamFight.json;monthSign.json;petCollect.json;seerQuiz.json;kxPackage.json;eventBossChallenge.json;eventPetAnaysis.json;activityExchange.json;achieve.json;military.json;treasureHunt.json;treasureWish.json;dailyTask.json;activityPoint.json;mainTask.json;mainTaskData.json;dailyDelegateTask.json;teamDonate.json;teamShop.json;teamShopPackage.json;teamProductionItems.json;teamTask.json;teamTech.json;petImprovement.json;improvementTime.json;encourage.json;weekSign.json;firstRecharge.json;getPhysical.json;suit.json;itemExchange.json;pvpLevel.json;equip.json;plantDescription.json;pvePlanet.json;pveSPT.json;features.json;eggExchange.json;title.json;mapItem.json;geneRecombine.json;vipPrivilege.json;starTrekBuy.json;pvePlanetBuy.json;vipLevelAward.json;vipWeekAward.json;monthCardFix.json;additive_attribute.json;natureRandom.json;petClass.json;talent.json;carnival.json;teamDonate.json;featureRandom.json;talentRandom.json;sevenPackage.json;buyCoin.json;petTest.json;starTrek.json;rankList.json;sns.json;invite.json;medalCondition.json;medal.json;fightBoss.json;braveTower.json;coolDown.json;springFestivalBag.json;dailyPayEvent.json;partyBoss.json;studyPoint.json;elemTypes.json;fightRelation.json;mainPetTraining.json;leadAwaken.json;springRedbags.json;darkArena.json;activation.json;activityWeeklyStory.json;activityPetExplore.json;attr.json;newInvention.json;activityPetTalent.json;superPetFight.json;eventBossBottom.json;extraDrop.json;petSkin.json;skinBuy.json;puniChallenge.json;skillRune.json;activityTimeLimit.json;shopMass.json;shopConsole.json;childrensDayRecharge.json;shopPackage.json;newcomerPetCollect.json;activityCard.json;newPrizePool.json;newLottery.json;battleLadder.json;roomUnit.json;castPvp.json;activityTaskReset.json;battleCondition.json;mapBossChallenge.json;mapBossBottom.json;mapBossRank.json;petTrain.json;VIPshop.json;activityLuckdraw.json;activityPetStory.json;rechargeBenefits.json;battlePass.json;battlePassTask.json;payPrizeBrandNew.json;firstPrizePool.json;dropExchange.json;checkIn.json;costPack.json;touchEvent.json;monthPet.json;pushPresent.json;activityPool.json;activityPoolFloor.json;activityLuckFloor.json;activityFreeLuckydraw.json;randomMonster.json;holidayShop.json;payEquipEvent.json"
var battleHotUpdate = "fightSkills.json;fightAffects.json;fightRelation.json;elemTypes.json;fightFactors.json"

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

	reth, errh := utils.HttpPost(utils.GetSplanSwitchUrl()+"/global_mail_proc",
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
	user := store.Get(global.SESSION_KEY_USER)
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

	body := struct {
		Modules   string `json:"modules"`    //0 表示全服
		Opt       string `json:"opt"`        //reload_json 代表重载配表
		JsonFiles string `json:"json_files"` //要重载的配表的名字,以;分隔
	}{Modules: "0", Opt: "reload_json"}

	var router string
	if strings.Contains(module, "online") {
		body.JsonFiles = onlineHotUpdate
		router = "/deal_online_cmd"
	}
	if strings.Contains(module, "battle") {
		body.JsonFiles = battleHotUpdate
		router = "/deal_battle_cmd"
	}
	bodyData, _ := json.Marshal(body)

	reth, errh := utils.HttpPost(utils.GetSplanSwitchUrl()+router,
		string(bodyData))
	if errh != nil {
		mlog.Log("splan update config err:\r\n", errh.Error())
		c.JSON(fiber.Map{"code": global.RET_OK, "data": string(reth)})
		return
	}

	if !strings.Contains(string(reth), `"code":0`) {
		mlog.Log("splan update config err:\r\n", errh.Error())
		c.JSON(fiber.Map{"code": global.RET_ERR, "data": string(reth)})
		return
	}

	mlog.Log("splan update config ret:\r\n", string(reth))

	param := struct {
		Notify bool `json:"notify"`
	}{}
	if errBp := c.BodyParser(&param); errBp != nil {
		param.Notify = false
	}
	if param.Notify {
		dingMsg := fmt.Sprintf("⚠ %s提交的热更%s配表请求已执行完成", user, module)
		formatMsg := fmt.Sprintf(global.DINGDING_TEXT_MSG_PATTERN, dingMsg)
		retd, errd := utils.HttpPost(utils.GetDingdingWebhook(), formatMsg)
		mlog.Log("hot update webhook ret:\r\n", string(retd), errd)
	}

	c.JSON(fiber.Map{"code": global.RET_OK, "data": string(reth)})
}

func SplanChangeTime(c *fiber.Ctx) {
	store := sessions.Get(c)
	user := store.Get(global.SESSION_KEY_USER)
	body := struct {
		Ip   string `json:"ip"`
		Time string `json:"time"`
	}{}

	if errBp := c.BodyParser(&body); errBp != nil {
		c.JSON(fiber.Map{"code": global.RET_ERR_BODY_PARAM,
			"data": errBp.Error()})
		return
	}

	maciniInDB, errm := db.Keys(global.BUCKET_OPS_MACINI)
	if errm != nil {
		mlog.Log("SplanChangeTime errm", errm)
		c.JSON(fiber.Map{"code": global.RET_ERR_DB,
			"data": errm.Error()})
		return
	}
	var sshport, sshuser, sshpasswd string

	for i := 1; i < len(maciniInDB); i += 2 {
		var imac []string
		e := json.Unmarshal([]byte(maciniInDB[i]), &imac)
		if e != nil {
			continue
		}
		if imac[0] == body.Ip {
			sshport = imac[1]
			sshuser = imac[2]
			sshpasswd = imac[3]
			break
		}
	}
	sshClt, errdial := ssh.Dial("tcp",
		fmt.Sprintf("%s:%s", body.Ip, sshport),
		&ssh.ClientConfig{
			User:            sshuser,
			Auth:            []ssh.AuthMethod{ssh.Password(sshpasswd)},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		})
	if errdial != nil {
		mlog.Log("ssh dial failed", errdial)
		c.JSON(fiber.Map{"code": global.RET_ERR_SPAWN,
			"data": errdial.Error()})
		return
	}
	defer sshClt.Close()
	mlog.Log("ssh dial passed")

	essh, _, errssh := expect.SpawnSSH(sshClt, time.Minute)
	if errssh != nil {
		c.JSON(fiber.Map{"code": global.RET_ERR_SPAWN,
			"data": errssh.Error()})
		return
	}
	defer essh.Close()
	logined := regexp.MustCompile("$")
	retlogin, _, elogin := essh.Expect(logined, 10*time.Second)
	if elogin != nil {
		mlog.Log("ssh login failed", elogin, retlogin)
		c.JSON(fiber.Map{"code": global.RET_ERR_SPAWN,
			"data": elogin.Error()})
		return
	}
	mlog.Log("ssh login passed", retlogin)
	essh.Send(fmt.Sprintf("sudo date -s '%s'\n", body.Time))
	retsd, _, esd := essh.Expect(regexp.MustCompile(
		fmt.Sprintf("password for %s:", sshuser)),
		10*time.Second)
	if esd != nil {
		mlog.Log("sudo date -s failed", esd, retsd)
		c.JSON(fiber.Map{"code": global.RET_ERR_SPAWN,
			"data": elogin.Error()})
		return
	} else {
		mlog.Log("sudo passed", retsd)
	}
	essh.Send(fmt.Sprintf("%s\n", sshpasswd))
	retok, matched, eok := essh.Expect(regexp.MustCompile("$"),
		5*time.Second)
	if eok != nil {
		mlog.Log("change server time error", eok, retok)
		c.JSON(fiber.Map{"code": global.RET_ERR_SPAWN,
			"data": eok.Error()})
		return
	} else {
		mlog.Log(user, "change server time ok", body.Time, retok, matched)
	}

	dingMsg := fmt.Sprintf("⚠ %s已将服务器时间修改为%s", user, body.Time)
	formatMsg := fmt.Sprintf(global.DINGDING_TEXT_MSG_PATTERN, dingMsg)
	retd, errd := utils.HttpPost(utils.GetDingdingWebhook(), formatMsg)
	mlog.Log("change server webhook ret:\r\n", string(retd), errd)

	c.JSON(fiber.Map{"code": global.RET_OK, "data": nil})
	return
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

func CreateDep(c *fiber.Ctx) {
	store := sessions.Get(c)
	// user := store.Get(global.SESSION_KEY_USER)
	group := store.Get(global.SESSION_KEY_GROUP)
	if group == nil {
		c.JSON(fiber.Map{"code": global.RET_ERR_SESSION_INVALID,
			"data": "session invalid"})
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
	err := db.Set(global.BUCKET_OPS_DEPBIL, "", body.Value)
	if err != nil {
		c.JSON(fiber.Map{"code": global.RET_ERR_DB,
			"data": err.Error()})
		return
	}
	c.JSON(fiber.Map{"code": global.RET_OK, "data": nil})
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
