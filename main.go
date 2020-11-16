package main

import (
	"fakeme/Services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/rest/v1")
	{
		v1.GET("/users", Services.GetUsers)
		v1.GET("/passwords", Services.GetPasswords)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
