// uart_error_detector.go
// A simple Go program for Raspberry Pi that reads incoming UART lines
// and logs any line containing the keyword "error".

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/tarm/serial"
)

var ErrorWords = []string{"error", "panic", "fatal"}

func containsError(s string) bool {
	for _, word := range ErrorWords {
		if strings.Contains(s, word) {
			return true
		}
	}
	return false
}

func main() {
	// Command-line flags for serial configuration
	portName := flag.String("port", "/dev/serial0", "Serial port device (e.g., /dev/serial0 or /dev/ttyAMA0)")
	baudRate := flag.Int("baud", 115200, "Baud rate for serial communication")
	flag.Parse()

	// Open the serial port
	cfg := &serial.Config{Name: *portName, Baud: *baudRate}
	port, err := serial.OpenPort(cfg)
	if err != nil {
		log.Fatalf("Failed to open serial port %s: %v", *portName, err)
	}
	defer port.Close()

	log.Printf("Listening on %s at %d baud...", *portName, *baudRate)

	// Read lines from UART
	scanner := bufio.NewScanner(port)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		// Detect the error keywords (case-insensitive)
		if containsError(line) {
			log.Printf("Error detected: %s", line)
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading from serial port: %v", err)
	}
}
