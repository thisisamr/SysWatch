package metrics_test

import (
	"testing"

	"github.com/thisisamr/SysWatch/internal/metrics"
	"github.com/thisisamr/SysWatch/internal/mocks"
)

func TestGetMemoryInfo(t *testing.T) {

	result, err := metrics.GetMemoryInfo(&mocks.MockProvider{})
	if err != nil {
		t.Fatalf("Failed to get memory info: %v", err)
	}

	if result.SwapStat.Total != 4000000 || result.SwapStat.UsedPercent != 50.0 {
		t.Errorf("Unexpected swap stats: %v", result.SwapStat)
	}
}
