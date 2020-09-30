package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("template/**/*")

	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "post/index.html", gin.H{
			"title": "Posts",
		})
	})
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "a2/index.html", gin.H{
			"title": "Users",
		})
	})

	//router.LoadHTMLGlob("static/*")
	//router.Static("/static", "static")
	//
	//router.GET("/list.html", api.Front.SearchByCon)
	//router.GET("/list-con.html",api.Front.SearchDetail)
	//router.GET("/search.html",api.Front.SearchByTags)

	router.Run(":8081")
}
