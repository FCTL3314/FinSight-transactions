package main

import (
	"github.com/FCTL3314/FinSight-transactions/internal/bootstrap"
	_ "github.com/lib/pq"
)

func main() {
	app := bootstrap.NewApplication()
	app.Run()
}
