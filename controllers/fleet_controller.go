package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webap-with-go/database"
	"github.com/hyperyuri/webap-with-go/models"
)

func ShowFleets(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return

	}

	db := database.GetDatabase()

	var fleets models.Fleets
	err = db.First(&fleets, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find fleets: " + err.Error(),
		})

		return
	}

	c.JSON(200, fleets)
}

func CreateFleet(c *gin.Context) {
	db := database.GetDatabase()

	var fleets models.Fleets

	err := c.ShouldBindJSON(&fleets)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})

		return
	}

	err = db.Create(&fleets).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create Fleets: " + err.Error(),
		})

		return
	}

	c.JSON(200, fleets)
}

func ShowFleet(c *gin.Context) {
	db := database.GetDatabase()

	var fleets []models.Fleets
	err := db.Find(&fleets).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list fleets: " + err.Error(),
		})

		return
	}
	c.JSON(200, fleets)
}

func UpdateFleet(c *gin.Context) {
	db := database.GetDatabase()

	var fleets models.Fleets

	err := c.ShouldBindJSON(&fleets)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})

		return
	}

	err = db.Save(&fleets).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot UPDATE Fleets: " + err.Error(),
		})

		return
	}

	c.JSON(200, fleets)
}

func DeleteFleets(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return

	}

	db := database.GetDatabase()

	err = db.Delete(&models.Fleets{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete fleet:" + err.Error(),
		})
		return

	}

	c.Status(204)

}
