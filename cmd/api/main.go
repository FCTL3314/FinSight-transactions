package main

import (
	"github.com/FCTL3314/FinSight-transactions/internal/bootstrap"
	"github.com/FCTL3314/FinSight-transactions/internal/logging"
)

func main() {
	app := bootstrap.NewApplication()

	if err := app.Run(); err != nil {
		app.LoggerGroup.General.Error(
			"The application ended with an error ...",
			logging.WithField("error", err),
		)
	}
}
