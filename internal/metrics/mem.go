package metrics

import (
	"github.com/shirou/gopsutil/v4/mem"
)

type MemoryStatResult struct {
	MemoryStat *mem.VirtualMemoryStat
}

func GetMemInfo(memProvider StatProvider) (*MemoryStatResult, error) {
	v, err := memProvider.MemoryInfo()
	if err != nil {
		return nil, err
	}
	return &MemoryStatResult{
		MemoryStat: v,
	}, nil
}
