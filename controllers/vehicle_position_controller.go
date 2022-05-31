package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webap-with-go/database"
	"github.com/hyperyuri/webap-with-go/models"
)

func Showvehicle_position(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return

	}

	db := database.GetDatabase()

	var vehicle_position models.Vehicle_Position
	err = db.First(&vehicle_position, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find vehicle_position: " + err.Error(),
		})

		return
	}

	c.JSON(200, vehicle_position)
}

func Createvehicle_position(c *gin.Context) {
	db := database.GetDatabase()

	var vehicle_position models.Vehicle_Position

	err := c.ShouldBindJSON(&vehicle_position)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})

		return
	}

	err = db.Create(&vehicle_position).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create vehicle_position: " + err.Error(),
		})

		return
	}

	c.JSON(200, vehicle_position)
}

func Showvehicle_positions(c *gin.Context) {
	db := database.GetDatabase()

	var vehicle_position []models.Vehicle_Position
	err := db.Find(&vehicle_position).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list vehicle_positions: " + err.Error(),
		})

		return
	}
	c.JSON(200, vehicle_position)
}

func Updatevehicle_position(c *gin.Context) {
	db := database.GetDatabase()

	var vehicle_position models.Vehicle_Position

	err := c.ShouldBindJSON(&vehicle_position)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})

		return
	}

	err = db.Save(&vehicle_position).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot UPDATE vehicle_position: " + err.Error(),
		})

		return
	}

	c.JSON(200, vehicle_position)
}

func Deletevehicle_position(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return

	}

	db := database.GetDatabase()

	err = db.Delete(&models.Vehicle_Position{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete vehicle_position" + err.Error(),
		})
		return

	}

	c.Status(204)

}
