package main

import (
	"go-node-audit/config"
	"go-node-audit/internal/audit"

	golog "github.com/ipfs/go-log"
)

var log = golog.Logger("Main")

func main() {
	log.Info("Starting audit job")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	auditService := audit.New(cfg)
	if err := auditService.Start(); err != nil {
		log.Fatalf("Audit failed: %v", err)
	}
}
