package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerRoutes() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*.html")
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", nil)
	})

	r.GET("/employees/:id/vacation", func(context *gin.Context) {
		id := context.Param("id")
		timesOff, ok := TimesOff[id]
		if !ok {
			context.String(http.StatusNotFound, "404 - Page Not Found")
			return
		}
		context.HTML(http.StatusOK, "vacation-overview.html", map[string]interface{}{
			"TimesOff": timesOff,
		})
	})

	admin := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))
	admin.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "admin-overview.html", nil)
	})

	r.Static("/public", "./public")
	return r
}