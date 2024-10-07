package routes

import (
	"Backend-Go/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.POST("/vehicle/entry", controller.VehicleEntry)
    r.PUT("/vehicle/exit/:id", controller.VehicleExit)

    return r
}
