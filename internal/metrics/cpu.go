package metrics

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

type CpuInfoResult struct {
	CpuInfo  []cpu.InfoStat
	CpuUsage []float64
}

func GetCPUInfo(provider StatProvider) (*CpuInfoResult, error) {
	info, err := provider.CPUInfo()
	if err != nil {
		return &CpuInfoResult{}, err
	}
	usage, err := provider.CPUUsage(time.Second)
	if err != nil {
		fmt.Println("Error:", err)
		return &CpuInfoResult{}, err
	}

	return &CpuInfoResult{CpuInfo: info, CpuUsage: usage}, nil

}
