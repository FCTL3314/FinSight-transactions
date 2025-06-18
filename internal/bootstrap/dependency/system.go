package dependency

import (
	"github.com/FCTL3314/FinSight-transactions/internal/api/router"
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/logging"
	"github.com/gin-gonic/gin"
)

type SystemContainer struct {
	Router            router.SystemRouter
	RouterRegistrator router.Registrator
	Logger            logging.Logger
}

func NewSystemContainer(
	baseRouter *gin.RouterGroup,
	cfg *config.Config,
) *SystemContainer {
	var container SystemContainer

	container.Router = router.NewSystemRouter(
		baseRouter,
		cfg,
	)
	container.RouterRegistrator = router.NewSystemRouterRegistrator(
		container.Router,
	)

	return &container
}
