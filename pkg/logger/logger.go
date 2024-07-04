package logger

import (
	"io"
	"log"
	"os"
)

var (
	logInitialized bool
)

func init() {
	if !logInitialized {
		if err := initLog(); err != nil {
			log.Printf("failed started logger: %v", err)
		}
		logInitialized = true
		log.Print("logger started with success")
	}
}

func initLog() error {
	logFile, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	log.SetOutput(io.MultiWriter(os.Stdout, logFile))
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	return nil
}
