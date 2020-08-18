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

	body := struct {
		Modules   string `json:"modules"`    //0 表示全服
		Opt       string `json:"opt"`        //reload_json 代表重载配表
		JsonFiles string `json:"json_files"` //要重载的配表的名字,以;分隔
	}{Modules: "0", Opt: "reload_json"}

	var router string
	if strings.Contains(module, "online") {
		body.JsonFiles = onlineHotUpdate
		router = "deal_online_cmd"
	}
	if strings.Contains(module, "battle") {
		body.JsonFiles = battleHotUpdate
		router = "deal_battle_cmd"
	}
	bodyData, _ := json.Marshal(body)

	reth, errh := utils.HttpPost("http://10.1.1.43:21010/"+router,
		string(bodyData))
	if errh != nil {
		mlog.Log("splan update config err:\r\n", errh.Error())
	} else {
		mlog.Log("splan update config ret:\r\n", string(reth))
	}

	c.JSON(fiber.Map{"code": global.RET_OK, "data": string(reth)})
}
