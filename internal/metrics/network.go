package metrics

import (
	"github.com/shirou/gopsutil/v4/net"
)

// Struct to hold network usage results
type NetUsageResult struct {
	Counters []net.IOCountersStat
}

func GetNetUsage(provider StatProvider) (*NetUsageResult, error) {
	counters, err := provider.IOCounters(true)
	if err != nil {
		return nil, err
	}

	return &NetUsageResult{
		Counters: counters,
	}, nil
}
