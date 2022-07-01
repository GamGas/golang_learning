package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default() // *gin.Engine тип, проверено

	router.POST("/employee")       //	CREATE
	router.GET("/employee/:id")    //	READ
	router.PUT("/employee/:id")    //	UPDATE
	router.DELETE("/employee/:id") //	DELETE

	router.Run() // Стандартный вид без аргументов равнозначен .Run(":8080")
}
