package metrics_test

import (
	"testing"

	"github.com/thisisamr/SysWatch/internal/metrics"
	"github.com/thisisamr/SysWatch/internal/mocks"
)

func TestCpu(t *testing.T) {

	result, err := metrics.GetCPUInfo(&mocks.MockProvider{})
	if err != nil {
		t.Errorf("failed to get the Cpu Info")
	}
	for _, res := range result.CpuInfo {
		if res.ModelName != "Intel Core i7" || res.Cores != 8 || res.Mhz != 2400 {
			t.Errorf("expected {ModelName: \"Intel Core i7\", Cores: 8, Mhz: 2400} but got \n {ModelName:\"%v\", Cores: %v, Mhz: %v}", res.ModelName, res.Cores, res.Mhz)
		}
	}
}
