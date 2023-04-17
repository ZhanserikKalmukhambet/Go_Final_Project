package main

import (
	"fmt"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/initializers"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	fmt.Println("This is migrate function")
	initializers.DB.AutoMigrate(&models.Car{})
	initializers.DB.AutoMigrate(&models.CarImage{})
}
