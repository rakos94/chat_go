package setup

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"raks.com/kamal/practice_crud/src/utils"
)

var username string
var password string
var host string
var port string
var dbName string
var db *gorm.DB

func initVar() {
	username = utils.GetDotEnvVariable(utils.OptionalValue{Value: "DB_USERNAME"})
	password = utils.GetDotEnvVariable(utils.OptionalValue{Value: "DB_PASSWORD"})
	host = utils.GetDotEnvVariable(utils.OptionalValue{Value: "DB_HOST"})
	port = utils.GetDotEnvVariable(utils.OptionalValue{Value: "DB_PORT"})
	dbName = utils.GetDotEnvVariable(utils.OptionalValue{Value: "DB_NAME"})
}

func initConnection() error {
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		err = errors.New("Cannot connect to DB")
	}
	db = _db
	return err
}

// InitDatabase initializing db
func InitDatabase() (*gorm.DB, error) {
	initVar()
	err := initConnection()
	return db, err
}

// GetDB var
func GetDB() *gorm.DB {
	return db
}
