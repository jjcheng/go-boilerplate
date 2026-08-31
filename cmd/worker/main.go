package main

import (
	"log"
	"os"
	"time"

	"github.com/jjcheng/go-boilerplate/internal/cfg"
)

func main() {
	// set timezone to utc so no need to call .UTC() everytime
	time.Local = time.UTC
	// set log to stdout as error will be handled by loggerService
	log.SetOutput(os.Stdout)
	log.Println("starting worker")
	log.Printf("environment: %s\n", cfg.Default().Site.Environment)
	// setup database
	// logger := service.NewLogger()
	// unitOfWork, err := setup.SetupDatabase(cfg.Default().Database.DSN(), logger)
	// if err != nil {
	// 	log.Fatalf("failed to setup database: %v", err)
	// }
	// dependencies := setup.SetupServices(unitOfWork, logger)
	// ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	// defer stop()
	log.Println("worker started")
	// start worker
	log.Println("worker stopped")
	log.Println("server exiting")
}
