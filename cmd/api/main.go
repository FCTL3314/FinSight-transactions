package main

import (
	"github.com/FCTL3314/FinSight-transactions/internal/bootstrap"
)

func main() {
	app := bootstrap.NewApplication()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
