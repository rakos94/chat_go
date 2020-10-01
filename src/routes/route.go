package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RouteList ...
func RouteList(r *gin.Engine, db *gorm.DB) {
	MainRoute(r)
	RouteRoomList(r, db)
}
