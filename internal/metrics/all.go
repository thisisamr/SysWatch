package metrics

import (
	"github.com/shirou/gopsutil/v4/process"
)

type Metrics struct {
	SystemInfo *SystemInfoResult
	CPUInfo    *CpuInfoResult    `json:"cpu_info"`
	MemoryInfo *MemoryStatResult `json:"memory_info"`
	DiskInfo   *DiskUsageResult  `json:"disk_info"`
	NetInfo    *NetUsageResult   `json:"net_info"`
	Processes  []*Proc           `json:"processes"`
}
type Proc struct {
	Name       string
	Pid        int32
	CpuPercent float64
	Mem        *process.MemoryInfoStat
}

func GatherAllMetrics(p StatProvider) (*Metrics, error) {
	// Gather system info
	hostInfo, err := GetSystemInfo(p)
	if err != nil {
		return nil, err
	}
	// Use channels to gather CPU and process info in parallel
	cpuChan := make(chan *CpuInfoResult)
	processChan := make(chan []*Proc)
	errorChan := make(chan error, 2)

	go func() {
		cpuInfo, err := GetCPUInfo(p)
		if err != nil {
			errorChan <- err
			return
		}
		cpuChan <- cpuInfo
	}()

	go func() {
		//REMOVED the sort here
		processes, err := GetTopProcesses(p, 20)
		if err != nil {
			errorChan <- err
			return
		}
		process_array := []*Proc{}
		for _, p := range processes {
			name, _ := p.Name()
			pid := p.Pid()
			cpu, _ := p.CPUPercent()
			mem, _ := p.MemoryInfo()
			process_array = append(process_array, &Proc{
				Name:       name,
				Pid:        pid,
				CpuPercent: cpu,
				Mem:        mem,
			})
		}
		processChan <- process_array
	}()

	// Gather memory info
	memInfo, err := GetMemInfo(p)
	if err != nil {
		return nil, err
	}

	// Gather disk usage
	diskInfo, err := GetDiskUsage(p, "/")
	if err != nil {
		return nil, err
	}

	// Gather network usage
	netInfo, err := GetNetUsage(p)
	if err != nil {
		return nil, err
	}

	// Wait for CPU and process info
	var cpuInfo *CpuInfoResult
	var process_array []*Proc

	for i := 0; i < 2; i++ {
		select {
		case cpuInfo = <-cpuChan:
		case process_array = <-processChan:
		case err := <-errorChan:
			return nil, err
		}
	}
	return &Metrics{
		SystemInfo: hostInfo,
		CPUInfo:    cpuInfo,
		MemoryInfo: memInfo,
		DiskInfo:   diskInfo,
		NetInfo:    netInfo,
		Processes:  process_array,
	}, err
}
