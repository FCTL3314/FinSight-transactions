package controller

import "github.com/gin-gonic/gin"

type GetController interface {
	Get(c *gin.Context)
}

type ListController interface {
	List(c *gin.Context)
}

type CreateController interface {
	Create(c *gin.Context)
}

type UpdateController interface {
	Update(c *gin.Context)
}

type DeleteController interface {
	Delete(c *gin.Context)
}

type Controller interface {
	GetController
	ListController
	CreateController
	UpdateController
	DeleteController
}
