package main

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"

	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/handler"
	"github.com/FCTL3314/FinSight-transactions/internal/repository"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := gorm.Open(postgres.Open(cfg.Database.GetDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect DB: %v", err)
	}

	repo := repository.NewTransactionRepo(db)
	h := handler.NewHandler(repo)
	log.Printf("starting server on %s", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(cfg.Server.Port, h.Routes()))
}
