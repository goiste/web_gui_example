package models

type Mem struct {
	Total       string `json:"total"`
	Used        string `json:"used"`
	Free        string `json:"free"`
	UsedPercent int64  `json:"used_percent"`
}
