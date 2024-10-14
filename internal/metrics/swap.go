package metrics

import (
	"github.com/shirou/gopsutil/v4/mem"
)

// Interface for Memory metrics (update existing)
type MemStatProvider interface {
	// VirtualMemory() (*mem.VirtualMemoryStat, error)
	SwapMemory() (*mem.SwapMemoryStat, error)
}

// Update RealMemStatProvider to implement SwapMemory
type RealMemStatProvider struct{}

func (r RealMemStatProvider) VirtualMemory() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}

// Struct to hold memory info (update existing)
type MemInfoResult struct {
	SwapStat *mem.SwapMemoryStat
}

// Update GetMemoryInfo to include swap info
func GetMemoryInfo(provider MemStatProvider) (*MemInfoResult, error) {

	s, err := provider.SwapMemory()
	if err != nil {
		return nil, err
	}

	return &MemInfoResult{
		SwapStat: s,
	}, nil
}
