package main

import (
	
	"github.com/gin-gonic/gin"
)


func main() {

	router := gin.Default()
	router.LoadHTMLGlob("bace/bace.html")
	router.GET("/", hendlerINdex)
	router.Run(":8080") 
}

func hendlerINdex (c *gin.Context) {
	c.HTML(200, "bace.html", nil)
	
}

