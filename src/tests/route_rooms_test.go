package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"raks.com/kamal/practice_crud/src/setup"
	"raks.com/kamal/practice_crud/src/utils"
)

func TestGetRooms(t *testing.T) {
	asserts := assert.New(t)
	setup.LoadEnvTest()
	db, _ := setup.InitDatabase()
	r := setup.RouterNoSetMode(false, db)
	w := utils.DoURLTest(r, "GET", "/room/1")

	asserts.Equal(200, w.Code, "Response Code must same")
	asserts.NotEmpty(w.Body)
}
