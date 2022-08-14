package sysinfo

import (
	"fmt"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/goiste/web_gui_example/src/models"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

func GetHostInfo() (hostData models.Host) {
	info, _ := host.Info()
	hostData.OS = cases.Title(language.English).String(info.OS)
	hostData.Arch = info.KernelArch
	hostData.Platform = cases.Title(language.English).String(info.Platform)
	hostData.Version = info.PlatformVersion
	hostData.Processes = info.Procs

	hostData.BootAt = getBootAt()
	hostData.Uptime = getUptimeString()

	return
}

func GetCPUInfo() (cpuData models.CPU) {
	cores, _ := cpu.Counts(false)
	threads, _ := cpu.Counts(true)
	info, _ := cpu.Info()

	cpuData.Name = info[0].ModelName
	cpuData.Cores = cores
	cpuData.Threads = threads
	cpuData.UsedPercent = getCPUPercent(0)

	return
}

func GetMemInfo() (memData models.Mem) {
	memory, _ := mem.VirtualMemory()

	memData.Total = formatSize(memory.Total)
	memData.Used = formatSize(memory.Used)
	memData.Free = formatSize(memory.Available)
	memData.UsedPercent = int64(memory.UsedPercent)

	return
}

func GetUpdate(d time.Duration) (upd models.Update) {
	hostInfo, _ := host.Info()
	memInfo, _ := mem.VirtualMemory()

	upd.Uptime = getUptimeString()
	upd.CPUUsedPercent = getCPUPercent(d)
	upd.MemUsed = formatSize(memInfo.Used)
	upd.MemFree = formatSize(memInfo.Available)
	upd.MemUsedPercent = int64(memInfo.UsedPercent)
	upd.Processes = hostInfo.Procs

	return
}

func getBootAt() time.Time {
	bootTime, _ := host.BootTime()
	return time.Unix(int64(bootTime), 0)
}

func getUptimeString() string {
	uptime, _ := host.Uptime()
	return fmt.Sprintf("%v", time.Second*time.Duration(uptime))
}

func getCPUPercent(d time.Duration) int64 {
	percent, _ := cpu.Percent(d, false)
	return int64(percent[0])
}
