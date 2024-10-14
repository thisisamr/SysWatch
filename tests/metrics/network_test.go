package metrics_test

import (
	"testing"

	"github.com/shirou/gopsutil/v4/net"
	"github.com/thisisamr/SysWatch/internal/metrics"
	"github.com/thisisamr/SysWatch/internal/mocks"
)

func TestGetNetUsage(t *testing.T) {

	result, err := metrics.GetNetUsage(&mocks.MockProvider{})
	if err != nil {
		t.Fatalf("Failed to get network usage: %v", err)
	}

	if len(result.Counters) == 0 {
		t.Fatal("Expected at least one network interface")
	}

	expected := net.IOCountersStat{Name: "eth0", BytesRecv: 1000, BytesSent: 500}
	actual := result.Counters[0]

	if actual.Name != expected.Name || actual.BytesRecv != expected.BytesRecv || actual.BytesSent != expected.BytesSent {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
