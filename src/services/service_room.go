package services

import (
	"raks.com/kamal/practice_crud/src/models"
	"raks.com/kamal/practice_crud/src/repositories"
	"raks.com/kamal/practice_crud/src/responses"
	"raks.com/kamal/practice_crud/src/utils"
)

// RoomService interface
type RoomService interface {
	Add(room *models.Room) (*responses.RoomDetailResponse, error)
	Detail(id int) (*responses.RoomDetailResponse, error)
	Update(id int, room *models.Room) (*responses.RoomDetailResponse, error)
	// Delete(id string) error
}

type roomService struct {
	roomRepo repositories.RoomRepository
	convert  utils.IMConvert
	result   *responses.RoomDetailResponse
	err      error
	res      interface{}
}

// NewRoomService func
func NewRoomService(roomRepo repositories.RoomRepository) RoomService {
	return &roomService{
		roomRepo: roomRepo,
		convert:  utils.NewIMConvert(),
		result:   &responses.RoomDetailResponse{},
	}
}

// Add func
func (r roomService) Add(room *models.Room) (*responses.RoomDetailResponse, error) {
	r.err = r.roomRepo.Add(room)
	if r.err != nil {
		return r.result, r.err
	}
	r.res = room

	return standartResult(r)
}

// Detail func
func (r roomService) Detail(id int) (*responses.RoomDetailResponse, error) {
	r.res, r.err = r.roomRepo.Detail(id)
	if r.err != nil {
		return r.result, r.err
	}

	return standartResult(r)
}

// Update func
func (r roomService) Update(id int, room *models.Room) (*responses.RoomDetailResponse, error) {
	r.err = r.roomRepo.Update(id, room)
	r.res, _ = r.roomRepo.Detail(id)

	return standartResult(r)
}

func standartResult(r roomService) (*responses.RoomDetailResponse, error) {
	r.convert.Object(r.res, &r.result)
	return r.result, r.err
}
