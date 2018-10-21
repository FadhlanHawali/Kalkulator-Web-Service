package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"kalkulatorWebService/controller"
)

func main() {
	router := gin.Default()
	router.GET("/",controller.Welcome)
	router.POST("/hitung",controller.Hitung)
	router.GET("/api/hitung",controller.HitungQuery)

	router.Run(":8080")
}