package metrics

import (
	"github.com/shirou/gopsutil/v4/disk"
)

// Struct to hold disk usage results
type DiskUsageResult struct {
	Usage      *disk.UsageStat
	Partitions []disk.PartitionStat
}

func GetDiskUsage(provider StatProvider, path string) (*DiskUsageResult, error) {
	usage, err := provider.DiskUsage(path)
	if err != nil {
		return nil, err
	}

	partitions, err := provider.DiskPartitions(false)
	if err != nil {
		return nil, err
	}

	return &DiskUsageResult{
		Usage:      usage,
		Partitions: partitions,
	}, nil
}
