package repositories

import (
	"time"

	"gorm.io/gorm"
	"raks.com/kamal/practice_crud/src/models"
)

// RoomRepository interface
type RoomRepository interface {
	Add(room *models.Room) error
	Detail(id int) (*models.Room, error)
	DetailAll() (*[]models.Room, error)
	Update(id int, room *models.Room) error
	Delete(id int) error
}

type roomRepository struct {
	db *gorm.DB
}

// NewRoomRepository func
func NewRoomRepository(db *gorm.DB) RoomRepository {
	return &roomRepository{db: db}
}

// Add func
func (r roomRepository) Add(room *models.Room) error {
	room.CreatedAt = time.Now()
	room.UpdatedAt = time.Now()
	err := r.db.Create(&room).Error
	return err
}

// Detail func
func (r roomRepository) Detail(id int) (*models.Room, error) {
	result := &models.Room{}
	err := r.db.Where("id = ?", id).First(result).Error
	return result, err
}

// DetailAll ...
func (r roomRepository) DetailAll() (*[]models.Room, error) {
	result := &[]models.Room{}
	err := r.db.Find(&result).Error
	return result, err
}

// Update func
func (r roomRepository) Update(id int, room *models.Room) error {
	err := r.db.Where("id = ?", id).Updates(room).Error
	return err
}

// Delete func
func (r roomRepository) Delete(id int) error {
	result := &models.Room{}
	err := r.db.Where("id = ?", id).Delete(result).Error
	return err
}
