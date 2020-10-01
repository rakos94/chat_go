package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"raks.com/kamal/practice_crud/src/models"
	"raks.com/kamal/practice_crud/src/requests"
	"raks.com/kamal/practice_crud/src/services"
	"raks.com/kamal/practice_crud/src/utils"
)

// RoomController ...
type RoomController interface {
	SHOW(c *gin.Context)
	STORE(c *gin.Context)
	UPDATE(c *gin.Context)
}

type roomController struct {
	service services.RoomService
	convert utils.IMConvert
}

// NewRoomController ...
func NewRoomController(service services.RoomService) RoomController {
	return &roomController{
		service: service,
		convert: utils.NewIMConvert(),
	}
}

// SHOW ...
func (r roomController) SHOW(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	response, err := r.service.Detail(id)
	if err != nil {
		utils.ResponseBad(c, err)
		return
	}
	utils.ResponseOK(c, response)
}

// STORE ...
func (r roomController) STORE(c *gin.Context) {
	var req requests.RoomAddRequest
	err := c.ShouldBind(&req)
	if err != nil {
		utils.ResponseBad(c, err)
		return
	}

	room := convertToRoom(&req)

	response, err := r.service.Add(room)
	if err != nil {
		utils.ResponseBad(c, err)
		return
	}
	utils.ResponseOK(c, response)
}

// UPDATE ...
func (r roomController) UPDATE(c *gin.Context) {
	var req requests.RoomAddRequest
	err := c.ShouldBind(&req)
	if err != nil {
		utils.ResponseBad(c, err)
		return
	}

	room := convertToRoom(&req)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ResponseBad(c, err)
		return
	}

	response, err := r.service.Update(id, room)
	if err != nil {
		utils.ResponseBad(c, err)
		return
	}
	utils.ResponseOK(c, response)
}

func convertToRoom(req *requests.RoomAddRequest) *models.Room {
	return &models.Room{
		Name: req.Name,
	}
}
