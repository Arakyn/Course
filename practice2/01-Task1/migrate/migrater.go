package main

import (
	"github.com/Arakyn/Course/practice2/01-Task1/Init"
	"github.com/Arakyn/Course/practice2/01-Task1/mod"
)

func init() {
	Init.LoadEnvVariables()
	Init.ConnectToDB()
}

func main() {
	Init.DB.AutoMigrate(&mod.State{})

}
