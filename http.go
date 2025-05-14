package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func httpServer() {
	http.HandleFunc("/events", sseHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.ListenAndServe(":80", nil)
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
	logFile := filepath.Join(os.Getenv("HOME"), "errors.log")
	file, err := os.Open(logFile)
	if err != nil {
		http.Error(w, "Could not open log file", 500)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	// Seek to the end of file
	file.Seek(0, 2)
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == nil {
			fmt.Fprintf(w, "data: %s\n\n", line)
			flusher.Flush()
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}
