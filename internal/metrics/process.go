package metrics

import (
	"sort"

	"github.com/shirou/gopsutil/v4/process"
)

// Process interface abstracts real process methods for mocking and testing
type Process interface {
	Pid() int32
	Name() (string, error)
	CPUPercent() (float64, error)
	MemoryInfo() (*process.MemoryInfoStat, error)
}

// // RealProcess implements Process interface using the gopsutil process library

type RealProcess struct {
	proc *process.Process
}

func (p *RealProcess) Pid() int32 {
	return p.proc.Pid
}

func (p *RealProcess) Name() (string, error) {
	return p.proc.Name()
}

func (p *RealProcess) CPUPercent() (float64, error) {
	return p.proc.CPUPercent()
}

func (p *RealProcess) MemoryInfo() (*process.MemoryInfoStat, error) {
	return p.proc.MemoryInfo()
}

// GetTopProcesses returns the top N processes sorted by CPU or memory usage
func GetTopProcesses(provider StatProvider, count int) ([]Process, error) {
	processes, err := provider.GetProcesses()
	if err != nil {
		return nil, err
	}

	// switch sortBy {
	// case "cpu":
	// 	sortByCPU(processes)
	// case "memory":
	// 	sortByMemory(processes)
	// default:
	// 	return nil, fmt.Errorf("invalid sortBy value: %s", sortBy)
	// }

	// Return the top N processes
	if count == 0 {
		return processes, nil
	}
	if len(processes) > count {
		return processes[len(processes)-count:], nil
	}
	return processes, nil
}

// sortByCPU sorts processes by CPU usage in descending order
func sortByCPU(processes []Process) {
	sort.Slice(processes, func(i, j int) bool {
		cpuI, errI := processes[i].CPUPercent()
		cpuJ, errJ := processes[j].CPUPercent()
		if errI != nil || errJ != nil {
			return false
		}
		return cpuI > cpuJ
	})
}

// sortByMemory sorts processes by memory usage in descending order
func sortByMemory(processes []Process) {
	sort.Slice(processes, func(i, j int) bool {
		memI, errI := processes[i].MemoryInfo()
		memJ, errJ := processes[j].MemoryInfo()
		if errI != nil || errJ != nil {
			return false
		}
		return memI.RSS > memJ.RSS
	})
}
