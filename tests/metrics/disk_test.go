package metrics_test

import (
	"testing"

	"github.com/shirou/gopsutil/v4/disk"
	"github.com/thisisamr/SysWatch/internal/metrics"
	"github.com/thisisamr/SysWatch/internal/mocks"
)

func TestGetDiskUsage(t *testing.T) {
	path := "/"

	result, err := metrics.GetDiskUsage(&mocks.MockProvider{}, path)
	if err != nil {
		t.Fatalf("Failed to get disk usage: %v", err)
	}

	expectedUsage := &disk.UsageStat{
		Path:        path,
		Total:       1000000,
		Free:        500000,
		Used:        500000,
		UsedPercent: 50.0,
	}

	if result.Usage.Total != expectedUsage.Total || result.Usage.UsedPercent != expectedUsage.UsedPercent {
		t.Errorf("Expected disk usage %v, got %v", expectedUsage, result.Usage)
	}

	if len(result.Partitions) != 1 || result.Partitions[0].Mountpoint != "/" {
		t.Errorf("Expected partition '/', got %v", result.Partitions)
	}
}
