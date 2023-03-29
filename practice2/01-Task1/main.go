package main

import (
	"github.com/Arakyn/Course/practice2/01-Task1/Init"
	"github.com/Arakyn/Course/practice2/01-Task1/controllers1"
	"github.com/gin-gonic/gin"
)

func init() {

	Init.LoadEnvVariables()
	Init.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/state", controllers1.StateCreate)
	r.GET("/state", controllers1.StateIndex)
	r.GET("/state/:id", controllers1.StateSingle)
	r.PUT("/state/:id", controllers1.StatesUpdate)
	r.DELETE("/state/:id", controllers1.DeleteState)
	r.Run()
}
