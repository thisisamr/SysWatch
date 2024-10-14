package metrics_test

import (
	"testing"

	"github.com/shirou/gopsutil/v4/host"
	"github.com/thisisamr/SysWatch/internal/metrics"
	"github.com/thisisamr/SysWatch/internal/mocks"
)

func TestGetSystemInfo(t *testing.T) {

	result, err := metrics.GetSystemInfo(&mocks.MockProvider{})
	if err != nil {
		t.Fatalf("Failed to get system info: %v", err)
	}

	expected := &host.InfoStat{
		Hostname: "test-host",
		OS:       "linux",
		Platform: "ubuntu",
		Uptime:   123456,
	}

	if result.Info.Hostname != expected.Hostname || result.Info.OS != expected.OS || result.Info.Platform != expected.Platform {
		t.Errorf("Expected %v, got %v", expected, result.Info)
	}
}
