package global

import (
	"sync"
)

type Progress struct {
	DepId  string  `json:"depid"`
	Status float64 `json:"status"`
	Desc   string  `json:"desc"`
}

type DingDingMarkdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
type SSHinfo struct {
	Port   string
	User   string
	Passwd string
}

type DepInfo struct {
	Type      string   `json:"type"`
	RepoUrl   string   `json:"repourl"`
	Rversion  string   `json:"rversion"`
	List      []string `json:list`
	StartTime int64
	DepId     string
}

var Depuuid2DepStatus sync.Map
var Depuuid2DepInfo sync.Map
var DepTypeAwakeTime sync.Map
