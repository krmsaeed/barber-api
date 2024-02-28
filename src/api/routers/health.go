package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/krmsaeed/barber-api/api/handlers"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()

	r.GET("/", handler.Health)
}
