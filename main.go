package main

import (
	"TheMostSimpleAPI/app/infrastructure"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := infrastructure.NewRouting()
	_ = r.Run()
}
