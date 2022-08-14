package sysinfo

import (
	"fmt"
)

const (
	_ = 1 << (iota * 10)
	kib
	mib
	gib
)

func formatSize(sizeInBytes uint64) string {
	switch {
	case sizeInBytes/gib > 0:
		return fmt.Sprintf("%.2f GiB", float64(sizeInBytes)/gib)
	case sizeInBytes/mib > 0:
		return fmt.Sprintf("%.2f MiB", float64(sizeInBytes)/mib)
	case sizeInBytes/kib > 0:
		return fmt.Sprintf("%.2f KiB", float64(sizeInBytes)/kib)
	default:
		return fmt.Sprintf("%d bytes", sizeInBytes)
	}
}
