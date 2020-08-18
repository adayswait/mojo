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
type DingDingBtn struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}
type DingDingActionCardContent struct {
	Title          string        `json:"title"`
	Text           string        `json:"text"`
	BtnOrientation string        `json:"btnOrientation"`
	Btns           []DingDingBtn `json:"btns"`
}
type DingDingActionCard struct {
	MsgType    string                    `json:"msgtype"`
	ActionCard DingDingActionCardContent `json:"actionCard"`
}
type SSHinfo struct {
	Port   string
	User   string
	Passwd string
}

type DepInfo struct {
	Type      string   `json:"type"`
	RepoUrl   string   `json:"repourl"`
	Revision  string   `json:"revision"`
	List      []string `json:list`
	StartTime int64
	DepId     string
}

var Depuuid2DepStatus sync.Map
var Depuuid2DepInfo sync.Map
var DepTypeAwakeTime sync.Map
