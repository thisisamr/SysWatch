package metrics_test

import (
	"testing"

	"github.com/thisisamr/SysWatch/internal/metrics"
	"github.com/thisisamr/SysWatch/internal/mocks"
)

func TestGetTopProcesses(t *testing.T) {
	// Using the mock process provider

	processes, err := metrics.GetTopProcesses(&mocks.MockProvider{}, 10)
	if err != nil {
		t.Fatalf("Failed to get top processes: %v", err)
	}

	if len(processes) != 2 {
		t.Errorf("Expected 2 processes, got %d", len(processes))
	}

	if name, _ := processes[0].Name(); name != "Process1" {
		t.Errorf("Expected Process1, got %v", name)
	}

	if cpu, _ := processes[0].CPUPercent(); cpu != 30.0 {
		t.Errorf("Expected Process1 with 30%% CPU, got %v%%", cpu)
	}
}
