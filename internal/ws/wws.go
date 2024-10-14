package ws

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/thisisamr/SysWatch/app/components"
	"github.com/thisisamr/SysWatch/internal/metrics"
)

// Subscriber holds the channels for the various metrics
type Subscriber struct {
	SystemInfoChan  chan *metrics.SystemInfoResult
	DiskInfoChan    chan *metrics.DiskUsageResult
	ProcessInfoChan chan []*metrics.Proc
	CpuInfoChan     chan *metrics.CpuInfoResult
	MemInfoChan     chan *metrics.MemoryStatResult
	conn            *websocket.Conn
	mu              sync.Mutex
	closeCh         chan struct{}
}

// Upgrade HTTP to WebSocket
var up = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Create a new subscriber with channels for each type of data
func NewSubscriber(conn *websocket.Conn) *Subscriber {
	return &Subscriber{
		SystemInfoChan:  make(chan *metrics.SystemInfoResult),
		DiskInfoChan:    make(chan *metrics.DiskUsageResult),
		ProcessInfoChan: make(chan []*metrics.Proc),
		CpuInfoChan:     make(chan *metrics.CpuInfoResult),
		MemInfoChan:     make(chan *metrics.MemoryStatResult),
		conn:            conn,
		closeCh:         make(chan struct{}), // Channel to signal when to stop
	}
}

// WebSocket handler for metrics data
func MetricsHandler(provider metrics.StatProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := up.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("WebSocket Upgrade error: %v\n", err)
			return
		}
		defer conn.Close()

		subscriber := NewSubscriber(conn)

		// Start metrics provider in a goroutine
		go startMetricsProvider(subscriber, provider)

		// Listen on subscriber channels and write messages when available
		go func() {
			for {
				select {
				case sysInfo := <-subscriber.SystemInfoChan:
					if !sendDataToWebSocket(subscriber, sysInfo) {
						return
					}
				case diskInfo := <-subscriber.DiskInfoChan:
					if !sendDataToWebSocket(subscriber, diskInfo) {
						return
					}
				case processInfo := <-subscriber.ProcessInfoChan:
					if !sendDataToWebSocket(subscriber, processInfo) {
						return
					}
				case cpuInfo := <-subscriber.CpuInfoChan:
					if !sendDataToWebSocket(subscriber, cpuInfo) {
						return
					}
				case memInfo := <-subscriber.MemInfoChan:
					if !sendDataToWebSocket(subscriber, memInfo) {
						return
					}
				case <-subscriber.closeCh:
					return // Stop the loop when connection is closed
				}
			}
		}()

		// Handle WebSocket read (just to detect client disconnection)
		go func() {
			for {
				_, _, err := conn.ReadMessage()
				if err != nil {
					log.Printf("WebSocket read error: %v\n", err)
					close(subscriber.closeCh) // Signal to close
					return
				}
			}
		}()

		select {} // Block forever
	}
}

// Send the gathered data to WebSocket
func sendDataToWebSocket(subscriber *Subscriber, data interface{}) bool {
	subscriber.mu.Lock()
	defer subscriber.mu.Unlock()

	var buf bytes.Buffer
	// Render the data to a component or serialize it to JSON
	switch v := data.(type) {
	case *metrics.SystemInfoResult:
		components.SystemInfo(v.Info).Render(context.Background(), &buf)
	case *metrics.DiskUsageResult:
		components.DiskInfo(v.Usage).Render(context.Background(), &buf)
		// Add other cases for different metrics as needed
	case *metrics.CpuInfoResult:
		components.CpuInfo(v).Render(context.Background(), &buf)
	case []*metrics.Proc:
		components.ProcessInfo(v).Render(context.Background(), &buf)
	case *metrics.MemoryStatResult:
		components.MemInfo(v.MemoryStat).Render(context.Background(), &buf)
	}

	// Write to WebSocket
	err := subscriber.conn.WriteMessage(websocket.TextMessage, buf.Bytes())
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			log.Printf("WebSocket closed unexpectedly: %v\n", err)
		}
		close(subscriber.closeCh) // Stop further attempts to send data
		return false
	}
	return true
}

// Simulate metrics provider pushing data to channels
func startMetricsProvider(subscriber *Subscriber, provider metrics.StatProvider) {
	for {
		select {
		case <-time.After(1 * time.Second):
			// Gather the metrics
			data, err := metrics.GatherAllMetrics(provider)
			if err != nil {
				log.Printf("Error gathering metrics: %v\n", err)
				return
			}
			// fmt.Println(data.CPUInfo)
			// Push the data to the appropriate channels
			select {
			case subscriber.SystemInfoChan <- data.SystemInfo:
			case <-subscriber.closeCh: // Stop pushing data if connection is closed
				return
			}
			select {
			case subscriber.DiskInfoChan <- data.DiskInfo:
			case <-subscriber.closeCh: // Stop pushing data if connection is closed
				return
			}
			select {
			case subscriber.ProcessInfoChan <- data.Processes:
			case <-subscriber.closeCh: // Stop pushing data if connection is closed
				return
			}
			select {
			case subscriber.CpuInfoChan <- data.CPUInfo:
			case <-subscriber.closeCh: // Stop pushing data if connection is closed
				return
			}
			select {
			case subscriber.MemInfoChan <- data.MemoryInfo:
			case <-subscriber.closeCh: // Stop pushing data if connection is closed
				return
			}
		case <-subscriber.closeCh: // Stop the metrics provider if connection is closed
			return
		}
	}
}
