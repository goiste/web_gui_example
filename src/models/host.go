package models

import (
	"time"
)

type Host struct {
	OS        string    `json:"os"`
	Arch      string    `json:"arch"`
	Platform  string    `json:"platform"`
	Version   string    `json:"version"`
	Processes uint64    `json:"processes"`
	BootAt    time.Time `json:"boot_at"`
	Uptime    string    `json:"uptime"`
}
