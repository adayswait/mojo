package utils

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const local_DB_DIR = "db-dir"                       //内置db的数据文件存储
const local_LOG_DIR = "log-dir"                     //日志文件存储路径
const local_FS_DIR = "fs-dir"                       //web服务host文件系统路径
const local_WWW_DIR = "www-dir"                     //web服务host静态文件地址
const local_WEB_HOST = "web-host"                   //web服务host地址
const local_WEB_PORT = "web-port"                   //web服务port
const local_WEB_DOMAIN = "web-domain"               //web页面所在域
const local_PKG_PATH = "pkg-path"                   //版本打包存储地址
const local_EXCLUDE_FROM = "exclude-from"           //部署除外文件路径
const local_DINGDING_WEBHOOK = "dingding-webhook"   //钉钉通知地址
const local_DEVELOPER_WEBHOOK = "developer-webhook" //开发者反馈通知地址
const local_SPLAN_SWITCH_URL = "splan-switch-url"   //splan switch gm的http地址

type config struct {
	dbDir            string
	logDir           string
	fsDir            string
	wwwDir           string
	webHost          string
	webPort          uint16
	webDomain        string
	pkgPath          string
	excludeFrom      string
	dingdingWebhook  string
	developerWebhook string
	splanSwitchUrl   string
}

var mojoCnf config

func init() {
	ReadConfig()
}

//判断input字符串是否为非0自然数
func IsNaturalNumber(input string) (ret bool, err error) {
	return regexp.MatchString("^[1-9][0-9]*$", input)
}

func ReadConfig() {
	fd, err := os.Open("./mojo.cnf")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		if string(line[0]) == "#" { //开头为#表示此行为注释
			continue
		}
		cnfArr := strings.Split(line, " ")
		switch cnfArr[0] {
		case local_DB_DIR:
			mojoCnf.dbDir = cnfArr[1]
		case local_LOG_DIR:
			mojoCnf.logDir = cnfArr[1]
		case local_FS_DIR:
			mojoCnf.fsDir = cnfArr[1]
		case local_WWW_DIR:
			mojoCnf.wwwDir = cnfArr[1]
		case local_WEB_HOST:
			mojoCnf.webHost = cnfArr[1]
		case local_WEB_PORT:
			portuint64, _ := strconv.ParseUint(cnfArr[1], 10, 16)
			mojoCnf.webPort = uint16(portuint64)
		case local_WEB_DOMAIN:
			mojoCnf.webDomain = cnfArr[1]
		case local_PKG_PATH:
			mojoCnf.pkgPath = cnfArr[1]
		case local_EXCLUDE_FROM:
			mojoCnf.excludeFrom = cnfArr[1]
		case local_DINGDING_WEBHOOK:
			mojoCnf.dingdingWebhook = cnfArr[1]
		case local_DEVELOPER_WEBHOOK:
			mojoCnf.developerWebhook = cnfArr[1]
		case local_SPLAN_SWITCH_URL:
			mojoCnf.splanSwitchUrl = cnfArr[1]
		default:

		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func GetServiceHost() string {
	return mojoCnf.webHost
}

func GetDbDir() string {
	return mojoCnf.dbDir
}

func GetLogDir() string {
	return mojoCnf.logDir
}

func GetFsDir() string {
	return mojoCnf.fsDir
}

func Get3wDir() string {
	return mojoCnf.wwwDir
}

func GetWebPort() uint16 {
	return mojoCnf.webPort
}

func GetWebDomain() []string {
	if len(mojoCnf.webDomain) == 0 {
		mojoCnf.webDomain = "http://" + mojoCnf.webHost + ":" +
			strconv.Itoa(int(mojoCnf.webPort))
		return []string{mojoCnf.webDomain}
	}
	return strings.Split(mojoCnf.webDomain, ";")
}

func GetPkgPath() string {
	return mojoCnf.pkgPath
}

func GetExcludeFrom() string {
	return mojoCnf.excludeFrom
}

func GetDingdingWebhook() string {
	return mojoCnf.dingdingWebhook
}

func GetDeveloperWebhook() string {
	return mojoCnf.developerWebhook
}

func GetSplanSwitchUrl() string {
	return mojoCnf.splanSwitchUrl
}

func IsPathExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}
