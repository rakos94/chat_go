package requests

// RoomAddRequest room response
type RoomAddRequest struct {
	Name string `form:"name" json:"name" binding:"required" example:"Lich Truong"`
}
