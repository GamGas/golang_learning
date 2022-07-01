package main

import "github.com/gin-gonic/gin"

func main() {
	memoryStorage := NewMemoryStorage()
	handler := NewHandler(memoryStorage)

	router := gin.Default() // *gin.Engine тип, проверено

	router.POST("/employee", handler.CreateEmployee)       //	CREATE
	router.GET("/employee/:id", handler.GetEmployee)       //	READ
	router.GET("/employee", handler.GetAllEmployees)       //	READALL
	router.PUT("/employee/:id", handler.UpdateEmployee)    //	UPDATE
	router.DELETE("/employee/:id", handler.DeleteEmployee) //	DELETE

	router.Run() // Стандартный вид без аргументов равнозначен .Run(":8080")
}
