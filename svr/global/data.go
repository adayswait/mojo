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

var ProgressMap sync.Map
var BreakMap sync.Map
