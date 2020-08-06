package utils

import "regexp"
import "os"
import "bufio"
import "strings"
import "strconv"

const local_DB_DATA_DIR = "db-data-dir"
const local_LOG_DIR = "log-dir"
const local_HOST = "host"
const local_PORT = "port"
const local_REPO_PATH = "repo-path"

type config struct {
	dbDataDir string
	logDir    string
	host      string
	port      uint16
	repoPath  string
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
		case local_DB_DATA_DIR:
			mojoCnf.dbDataDir = cnfArr[1]
		case local_LOG_DIR:
			mojoCnf.logDir = cnfArr[1]
		case local_HOST:
			mojoCnf.host = cnfArr[1]
		case local_PORT:
			portuint64, _ := strconv.ParseUint(cnfArr[1], 10, 16)
			mojoCnf.port = uint16(portuint64)
		case local_REPO_PATH:
			mojoCnf.repoPath = cnfArr[1]
		default:

		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func GetDbDirCnf() string {
	return mojoCnf.dbDataDir
}

func GetLogDirCnf() string {
	return mojoCnf.logDir
}

func GetListeningPort() uint16 {
	return mojoCnf.port
}

func GetRepoPath() string {
	return mojoCnf.repoPath
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
