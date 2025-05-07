// uart_error_detector.go
// A simple Go program for Raspberry Pi that reads incoming UART lines
// and logs any line containing the keyword "error".

package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/tarm/serial"
)

var ErrorWords = []string{"error", "panic", "fatal"}

func containsError(s string) bool {
	for _, word := range ErrorWords {
		if strings.Contains(strings.ToLower(s), word) {
			return true
		}
	}
	return false
}

func main() {
	// Command-line flags for serial configuration
	portName := flag.String("port", "/dev/serial0", "Serial port device (e.g., /dev/serial0 or /dev/ttyAMA0)")
	baudRate := flag.Int("baud", 115200, "Baud rate for serial communication")
	logFilePath := flag.String("logfile", "uart_errors.log", "Path to error log file")
	flag.Parse()

	// Open the serial port
	cfg := &serial.Config{Name: *portName, Baud: *baudRate}
	port, err := serial.OpenPort(cfg)
	if err != nil {
		log.Fatalf("Failed to open serial port %s: %v", *portName, err)
	}
	defer port.Close()

	fmt.Printf("Listening on %s at %d baud... (press Ctrl+C to exit)\n", *portName, *baudRate)

	// Open log file
	logFile, err := os.OpenFile(*logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	// Setup signal handling for Ctrl+C
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Read lines from UART
	scanner := bufio.NewScanner(port)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			fmt.Println("\nReceived interrupt signal. Exiting.")
			return
		default:
			line := scanner.Text()
			fmt.Println(line)

			if containsError(line) {
				_, err := logFile.WriteString(fmt.Sprintf("Error detected: %s\n", line))
				if err != nil {
					log.Printf("Failed to write to log file: %v", err)
				}
			}
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading from serial port: %v", err)
	}
}
