package main

import (
	"TTD-golang-gin-test/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}