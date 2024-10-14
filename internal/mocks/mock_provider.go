package mocks

import (
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/process"
	"github.com/thisisamr/SysWatch/internal/metrics"
)

type MockProvider struct {
}

func (m *MockProvider) CPUInfo() ([]cpu.InfoStat, error) {
	return []cpu.InfoStat{
		{ModelName: "Intel Core i7", Cores: 8, Mhz: 2400},
	}, nil
}
func (m *MockProvider) CPUUsage(duration time.Duration) ([]float64, error) {

	return []float64{25.5}, nil
}

func (m *MockProvider) DiskUsage(path string) (*disk.UsageStat, error) {
	return &disk.UsageStat{
		Path:        path,
		Total:       1000000,
		Free:        500000,
		Used:        500000,
		UsedPercent: 50.0,
	}, nil
}

func (m *MockProvider) DiskPartitions(all bool) ([]disk.PartitionStat, error) {
	return []disk.PartitionStat{
		{Mountpoint: "/", Device: "/dev/sda1"},
	}, nil
}
func (m *MockProvider) Info() (*host.InfoStat, error) {
	return &host.InfoStat{
		Hostname: "test-host",
		OS:       "linux",
		Platform: "ubuntu",
		Uptime:   123456,
	}, nil
}
func (m *MockProvider) MemoryInfo() (*mem.VirtualMemoryStat, error) {
	return &mem.VirtualMemoryStat{
		Total:       12000,
		Used:        5000,
		Available:   7000,
		UsedPercent: 10.00,
	}, nil
}
func (m *MockProvider) IOCounters(pernic bool) ([]net.IOCountersStat, error) {
	return []net.IOCountersStat{
		{Name: "eth0", BytesRecv: 1000, BytesSent: 500},
	}, nil
}

// MockProcess simulates the Process interface for testing
type MockProcess struct {
	PidValue       int32
	NameFunc       func() (string, error)
	CPUPercentFunc func() (float64, error)
	MemoryInfoFunc func() (*process.MemoryInfoStat, error)
}

func (m *MockProcess) Pid() int32 {
	return m.PidValue
}

func (m *MockProcess) Name() (string, error) {
	return m.NameFunc()
}

func (m *MockProcess) CPUPercent() (float64, error) {
	return m.CPUPercentFunc()
}

func (m *MockProcess) MemoryInfo() (*process.MemoryInfoStat, error) {
	return m.MemoryInfoFunc()
}

// MockProcessProvider implements the ProcessProvider interface for testing
type MockProcessProvider struct{}

func (m *MockProvider) GetProcesses() ([]metrics.Process, error) {
	// Creating mock processes
	p1 := &MockProcess{
		PidValue:       1,
		NameFunc:       func() (string, error) { return "Process1", nil },
		CPUPercentFunc: func() (float64, error) { return 30.0, nil },
		MemoryInfoFunc: func() (*process.MemoryInfoStat, error) {
			return &process.MemoryInfoStat{RSS: 1000000}, nil
		},
	}

	p2 := &MockProcess{
		PidValue:       2,
		NameFunc:       func() (string, error) { return "Process2", nil },
		CPUPercentFunc: func() (float64, error) { return 20.0, nil },
		MemoryInfoFunc: func() (*process.MemoryInfoStat, error) {
			return &process.MemoryInfoStat{RSS: 2000000}, nil
		},
	}

	return []metrics.Process{p1, p2}, nil
}
func (m *MockProvider) SwapMemory() (*mem.SwapMemoryStat, error) {
	return &mem.SwapMemoryStat{
		Total:       4000000,
		Used:        2000000,
		Free:        2000000,
		UsedPercent: 50.0,
	}, nil
}
