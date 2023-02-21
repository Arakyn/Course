package controllers1

import (
	"log"

	"github.com/arakyn/Course/practice2/01-Task1/Init"
	"github.com/arakyn/Course/practice2/01-Task1/mod"
	"github.com/gin-gonic/gin"
)

func StateCreate(c *gin.Context) {

	// getting data off the request body

	var place struct {
		City    string
		Country string
	}
	c.Bind(&place)
	// creating a state
	state := mod.State{City: place.City, Country: place.Country}

	result := Init.DB.Create(&state)
	if result.Error != nil {
		c.Status(400)
	}
	c.JSON(200, gin.H{
		"State": state,
	})

}

func StateIndex(c *gin.Context) {
	// getting the index
	var States []mod.State

	result := Init.DB.Find(&States)
	if result.Error != nil {
		log.Fatal("Error at finding stuff")
		c.Status(500)
		return
	}
	c.JSON(200, gin.H{
		"States": States,
	})

	// responding with them

}
func StateSingle(c *gin.Context) {
	// getting the id off the url

	id := c.Param("id")
	var state mod.State

	Init.DB.First(&state, id)

	c.JSON(200, gin.H{
		"state": state,
	})
}

func StatesUpdate(c *gin.Context) {
	// Getting the id off the url
	id := c.Param("id")

	// get the data off the request body
	var updatedState struct {
		City    string
		Country string
	}
	c.Bind(&updatedState)

	// find the post we are updating
	var oldState mod.State
	Init.DB.First(&oldState, id)

	// update it
	Init.DB.Model(&oldState).Updates(mod.State{City: updatedState.City, Country: updatedState.Country})

	// respond with it
	c.JSON(200, gin.H{
		"Updated State": updatedState,
	})

}

// func delete just removes it, so it cant be shown on the website or something but the actual data
// is present in the database for data recovery.
func DeleteState(c *gin.Context) {

	// Stripping the ID of the url
	id := c.Param("id")

	// Deleting the state

	Init.DB.Delete(&mod.State{}, id)

	// Showing it ig?

	c.Status(200)

}
