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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ClockWsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket Upgrade error: %v\n", err)
		return
	}
	defer conn.Close()

	log.Println("WebSocket connection established.")

	var buf bytes.Buffer

	// Create a done channel to signal when to stop sending messages
	done := make(chan struct{})
	go func() {
		// Loop to write messages
		for {
			// Writing a message to the client
			buf.Reset() // Reset the buffer for each message
			components.Clock(time.Now().Format("Monday, January 2, 2006 15:04:05")).Render(context.Background(), &buf)

			// Attempt to write to the WebSocket
			select {
			case <-done:
				return // Exit if done signal is received
			default:
				err := conn.WriteMessage(websocket.TextMessage, buf.Bytes())
				if err != nil {
					log.Printf("Error writing message: %v\n", err)
					close(done) // Close the done channel to stop the loop
					return
				}
			}
			time.Sleep(1 * time.Second)
		}
	}()

	// Wait for the connection to close or an error to occur
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Printf("WebSocket read error: %v\n", err)
			close(done) // Signal to stop sending messages
			return
		}
	}
}
func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket Upgrade error: %v\n", err)
		return
	}
	defer conn.Close()

	log.Println("WebSocket connection established.")

	var buf bytes.Buffer
	for {
		// Writing a message to the client
		components.Clock((time.Now().Format("Monday, January 2, 2006 15:04:05"))).Render(context.Background(), &buf)
		err := conn.WriteMessage(websocket.TextMessage, buf.Bytes())
		if err != nil {
			log.Printf("Error writing message: %v\n", err)
			return
		}
		time.Sleep(1 * time.Second)

	}
}

func MetricsWsHandler(provider metrics.StatProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		mu := sync.Mutex{}
		if err != nil {
			log.Printf("WebSocket Upgrade error: %v\n", err)
			return
		}
		defer conn.Close()
		log.Println("Metrics WebSocket connection established.")
		go func() {

			for {

				// Clear the buffer before each use
				var buf bytes.Buffer

				// Gather metrics
				data, err := metrics.GatherAllMetrics(provider)
				if err != nil {
					log.Printf("Error gathering metrics: %v\n", err)
					return
				}

				// Render the system info component
				err = components.DiskInfo(data.DiskInfo.Usage).Render(r.Context(), &buf)
				if err != nil {
					log.Printf("Error rendering component: %v\n", err)
					return
				}

				// Write the message to the client
				mu.Lock()

				err = conn.WriteMessage(websocket.TextMessage, buf.Bytes())
				mu.Unlock()
				if err != nil {
					if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
						log.Printf("Unexpected WebSocket closure: %v\n", err)
					} else {
						log.Printf("WebSocket closed: %v\n", err)
					}
					break // Exit the loop on write error
				}

			}
		}()
		go func() {

			for {

				// Clear the buffer before each use
				var buf bytes.Buffer

				// Gather metrics
				data, err := metrics.GatherAllMetrics(provider)
				if err != nil {
					log.Printf("Error gathering metrics: %v\n", err)
					return
				}

				// Render the system info component
				err = components.SystemInfo(data.SystemInfo.Info).Render(context.Background(), &buf)
				if err != nil {
					log.Printf("Error rendering component: %v\n", err)
					return
				}

				// Write the message to the client
				mu.Lock()
				err = conn.WriteMessage(websocket.TextMessage, buf.Bytes())
				mu.Unlock()
				if err != nil {
					if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
						log.Printf("Unexpected WebSocket closure: %v\n", err)
					} else {
						log.Printf("WebSocket closed: %v\n", err)
					}
					break // Exit the loop on write error
				}

			}
		}()
		select {}
	}
}
