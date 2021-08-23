package main

import (
	"TTD-golang-gin-test/controller"
	"github.com/gin-gonic/gin"
)

//https://medium.com/@BranLim/basic-rules-for-effective-onion-architecture-a32af1f3b469
func main() {
	r := gin.Default()
	controller.RegisterRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}