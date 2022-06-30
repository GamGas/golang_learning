// routes.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes() {

	// Handle the index route
	router.GET("/", showIndexPage)

	// Handle GET requests at /article/view/some_article_id
	router.GET("/article/view/:article_id", getArticle)

}

func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}
