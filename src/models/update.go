package models

type Update struct {
	Uptime         string `json:"uptime"`
	CPUUsedPercent int64  `json:"cpu_used_percent"`
	MemUsed        string `json:"mem_used"`
	MemFree        string `json:"mem_free"`
	MemUsedPercent int64  `json:"mem_used_percent"`
	Processes      uint64 `json:"processes"`
}
