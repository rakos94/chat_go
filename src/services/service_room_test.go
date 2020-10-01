package services

import (
	"testing"

	"raks.com/kamal/practice_crud/src/requests"

	"github.com/stretchr/testify/assert"

	"raks.com/kamal/practice_crud/src/setup"

	"raks.com/kamal/practice_crud/src/repositories"
)

func TestRoomService_Add(t *testing.T) {
	setup.LoadEnvTest()
	db, err := setup.InitDatabase()
	if err != nil {
		print(err)
	}
	repos := repositories.NewRoomRepository(db)
	service := NewRoomService(repos)
	req := &requests.RoomAddRequest{
		Name: "Test Terresa",
	}
	res, err := service.Add(req)
	if err != nil {
		println(err)
	}
	assert.EqualValues(t, req.Name, res.Name)
}

func TestRoomService_Detail(t *testing.T) {
	setup.LoadEnvTest()
	db, err := setup.InitDatabase()
	if err != nil {
		println("error init db")
		// errors.New("error init db")
	}
	repos := repositories.NewRoomRepository(db)
	service := NewRoomService(repos)
	res, err := service.Detail(1)
	if err != nil {
		println("error service function")
		// errors.New("error service function")
	}
	assert.EqualValues(t, 1, res.ID)
}
