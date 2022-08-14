package models

type CPU struct {
	Name        string `json:"name"`
	Cores       int    `json:"cores"`
	Threads     int    `json:"threads"`
	UsedPercent int64  `json:"used_percent"`
}
