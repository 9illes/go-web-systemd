package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// The path matches the name defined in the .service file.
// File written in this directory will persist across service restart.
const logFile = "/var/log/goweb/log.txt"

func main() {

	// setup signal catching
	sigs := make(chan os.Signal, 1)

	// catch all signals since not explicitly listing
	signal.Notify(sigs)

	// method invoked upon seeing signal
	go func() {
		for sig := range sigs {
			log.Printf("RECEIVED SIGNAL: %s", sig)

			switch sig {
			case syscall.SIGURG:
				log.Printf("ignoring sigurg")
			default:
				// AppCleanup()
				os.Exit(1)
			}
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		log.Printf("/ => Hello, you've requested: %s\n", r.URL.Path)

		// systemd should restart this service if it crashes
		if r.URL.Path == "/panic" {
			log.Fatal("forced app crash")
			panic("crashed !")
		}

		defer timer("request handler")()
	})

	fmt.Println("Listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// timer returns a function that can be used as a timer to measure the execution time of a given task.
func timer(name string) func() {
	start := time.Now()
	return func() {
		writeCustomLog(fmt.Sprintf("%s %s took %v", time.Now().Format("2006-01-02 15:04:05"), name, time.Since(start)))
	}
}

func writeCustomLog(msg string) {
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(msg + "\n"); err != nil {
		log.Println(err)
	}
}
