// uart_error_detector.go
// A simple Go program for Raspberry Pi that reads incoming UART lines
// and logs any line containing the keyword "error".

package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/tarm/serial"
	"log"
	"os"
	"strings"
)

var ErrorWords = []string{"error", "panic", "fatal", "fail"}

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
	portName := flag.String("port", "/dev/serial0", "Serial port device")
	baudRate := flag.Int("baud", 115200, "Baud rate for serial communication")
	logFilePath := flag.String("logfile", "errors.log", "Path to error log file")
	flag.Parse()

	// Open the serial port
	cfg := &serial.Config{Name: *portName, Baud: *baudRate}
	port, err := serial.OpenPort(cfg)
	if err != nil {
		log.Fatalf(">>> Failed to open serial port %s: %v", *portName, err)
	}
	defer port.Close()

	fmt.Printf(">>> Listening on %s at %d baud... (press Ctrl+C to exit)\n", *portName, *baudRate)

	// Open log file
	logFile, err := os.OpenFile(*logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf(">>> Failed to open log file: %v", err)
	}
	defer logFile.Close()

	go httpServer()
	fmt.Println(">>> HTTP server started...")

	// Read lines from UART
	scanner := bufio.NewScanner(port)
	fmt.Println(">>> Scanner started...")

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if containsError(line) {
			fmt.Println(">>> Error detected. try to appending it to the file...")
			_, err := logFile.WriteString(line + "\n")
			if err != nil {
				fmt.Printf(">>> Failed to write to log file: %v\n", err)
			}
		}
		if strings.Contains(strings.ToLower(line), "Booting Linux on physical CPU 0x0000000000") {
			fmt.Println(">>> Reboot detected. try to appending it to the file...")
			_, err := logFile.WriteString(line + "\n")
			if err != nil {
				fmt.Printf(">>> Failed to write to log file: %v\n", err)
			}
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading from serial port: %v", err)
	}
}
