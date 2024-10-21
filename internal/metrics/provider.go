package metrics

import (
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/process"
)

type StatProvider interface {
	CPUInfo() ([]cpu.InfoStat, error)
	CPUUsage(duration time.Duration) ([]float64, error)
	DiskUsage(path string) (*disk.UsageStat, error)
	DiskPartitions(all bool) ([]disk.PartitionStat, error)
	Info() (*host.InfoStat, error)
	MemoryInfo() (*mem.VirtualMemoryStat, error)
	IOCounters(pernic bool) ([]net.IOCountersStat, error)
	GetProcesses() ([]Process, error)
	SwapMemory() (*mem.SwapMemoryStat, error)
}

type Provider struct {
}

func (p *Provider) CPUInfo() ([]cpu.InfoStat, error) {
	return cpu.Info()
}

func (p *Provider) CPUUsage(duration time.Duration) ([]float64, error) {

	return cpu.Percent(duration, true)
}

func (p *Provider) DiskUsage(path string) (*disk.UsageStat, error) {
	return disk.Usage(path)
}

func (p *Provider) DiskPartitions(all bool) ([]disk.PartitionStat, error) {
	return disk.Partitions(all)
}
func (p *Provider) Info() (*host.InfoStat, error) {
	return host.Info()
}

// Struct to hold system info
type SystemInfoResult struct {
	Info *host.InfoStat
}

func (p *Provider) MemoryInfo() (*mem.VirtualMemoryStat, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	return v, nil
	// v.Total/1024/1024, v.Used/1024/1024, v.Available/1024/1024, v.UsedPercent)
}
func (p *Provider) IOCounters(pernic bool) ([]net.IOCountersStat, error) {
	return net.IOCounters(pernic)
}

func (p *Provider) GetProcesses() ([]Process, error) {
	procs, err := process.Processes()
	if err != nil {
		return nil, err
	}

	var result []Process
	for _, proc := range procs {
		result = append(result, &RealProcess{proc: proc})
	}
	return result, nil
}
func (p *Provider) SwapMemory() (*mem.SwapMemoryStat, error) {
	return mem.SwapMemory()
}

func NewProvider() *Provider {
	return &Provider{}
}
