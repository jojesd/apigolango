package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webap-with-go/database"
	"github.com/hyperyuri/webap-with-go/models"
)

func ShowVehicle(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return

	}

	db := database.GetDatabase()

	var Vehicle models.Vehicle
	err = db.First(&Vehicle, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find vehicle: " + err.Error(),
		})

		return
	}

	c.JSON(200, Vehicle)
}

func CreateVehicle(c *gin.Context) {
	db := database.GetDatabase()

	var vehicle models.Vehicle

	err := c.ShouldBindJSON(&vehicle)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})

		return
	}

	err = db.Create(&vehicle).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create vehicle: " + err.Error(),
		})

		return
	}

	c.JSON(200, vehicle)
}

func ShowVehicles(c *gin.Context) {
	db := database.GetDatabase()

	var vehicle []models.Vehicle
	err := db.Find(&vehicle).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list vehicles: " + err.Error(),
		})

		return
	}
	c.JSON(200, vehicle)
}

func UpdateVehicle(c *gin.Context) {
	db := database.GetDatabase()

	var vehicle models.Vehicle

	err := c.ShouldBindJSON(&vehicle)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})

		return
	}

	err = db.Save(&vehicle).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot UPDATE vehicle: " + err.Error(),
		})

		return
	}

	c.JSON(200, vehicle)
}

func DeleteVehicle(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return

	}

	db := database.GetDatabase()

	err = db.Delete(&models.Vehicle{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete vehicle" + err.Error(),
		})
		return

	}

	c.Status(204)

}
