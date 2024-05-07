package main

import (
	"github.com/gin-gonic/gin"
	"xb1de-bdat-query/config"
	"xb1de-bdat-query/controllers"
)

func main() {
	r := gin.Default()

	r.POST("translations", controllers.Translations)

	r.Run(config.GetConfig().Addr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
