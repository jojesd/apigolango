package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webap-with-go/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		fleets := main.Group("Fleets")
		{
			fleets.GET("/:id", controllers.ShowFleets)
			fleets.GET("/", controllers.ShowFleet)
			fleets.POST("/", controllers.CreateFleet)
			fleets.PUT("/", controllers.UpdateFleet)
			fleets.DELETE("/:id", controllers.DeleteFleets)
		}

		vehicle := main.Group("Vehicles")
		{
			vehicle.GET("/:id", controllers.ShowVehicle)
			vehicle.GET("/", controllers.ShowVehicles)
			vehicle.POST("/", controllers.CreateVehicle)
			vehicle.PUT("/", controllers.UpdateVehicle)
			vehicle.DELETE("/:id", controllers.DeleteVehicle)
		}

		vehicle_position := main.Group("Vehicles_Position")
		{
			vehicle_position.GET("/:id", controllers.Showvehicle_position)
			vehicle_position.GET("/", controllers.Showvehicle_position)
			vehicle_position.POST("/", controllers.Createvehicle_position)
			vehicle_position.PUT("/", controllers.Updatevehicle_position)
			vehicle_position.DELETE("/:id", controllers.Deletevehicle_position)
		}

	}
	return router
}
