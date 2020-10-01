package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"raks.com/kamal/practice_crud/src/controllers"
	"raks.com/kamal/practice_crud/src/repositories"
	"raks.com/kamal/practice_crud/src/services"
)

// RouteRoomList ...
func RouteRoomList(r *gin.Engine, db *gorm.DB) {
	repo := repositories.NewRoomRepository(db)
	service := services.NewRoomService(repo)
	controller := controllers.NewRoomController(service)

	r.GET("/room/:id", controller.SHOW)
	r.POST("/room", controller.STORE)
	r.PUT("/room/:id", controller.UPDATE)
}
