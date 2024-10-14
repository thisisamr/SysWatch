package metrics_test

import (
	"testing"

	"github.com/thisisamr/SysWatch/internal/metrics"
	"github.com/thisisamr/SysWatch/internal/mocks"
)

func TestMem(t *testing.T) {
	result, err := metrics.GetMemInfo(&mocks.MockProvider{})
	if err != nil {

		t.Errorf("failed to get the memory Info")
	}
	if result.MemoryStat.Total != 12000 || result.MemoryStat.Used != 5000 || result.MemoryStat.Available != 7000 || result.MemoryStat.UsedPercent != 10 {

		t.Errorf("expected {Total: \"12000\", Used: 5000, UsedPercent: 10} but got \n {Total:\"%v\", Used: %v, UsedPercent: %f}", result.MemoryStat.Total, result.MemoryStat.Used, result.MemoryStat.UsedPercent)
	}

}
