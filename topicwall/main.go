package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Роутер по умолчанию в Gin
	router = gin.Default()

	// Обрабатываем шаблоны вначале,
	// так что их не нужно будет перечитывать снова.
	// Из-за этого вывод HTML-страниц такой быстрый
	router.LoadHTMLGlob("templates/*")

	// Инициализируем роуты
	initializeRoutes()

	// Запускаем приложение
	router.Run()
}
