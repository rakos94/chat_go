package setup

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"raks.com/kamal/practice_crud/src/routes"
	"raks.com/kamal/practice_crud/src/utils"
)

// RouterNoSetMode setup routers
func RouterNoSetMode(isServe bool, db *gorm.DB) *gin.Engine {
	mode := NewGinMode()
	r := Router(db)
	mode.SetNoMode()
	if isServe {
		serve(r)
	}

	return r
}

// Router setup routers
func Router(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	useCors(r)
	routes.RouteList(r, db)

	return r
}

func useCors(r *gin.Engine) {
	r.Use(cors.Default())
}

func serve(r *gin.Engine) {
	port := utils.GetDotEnvVariable(utils.OptionalValue{Value: "SERVE_PORT", DefaultValue: "8080"})
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
