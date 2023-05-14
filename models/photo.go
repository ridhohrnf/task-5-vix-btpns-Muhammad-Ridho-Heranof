package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Photo struct {
	ID        uint `gorm:"primary_key"`
	Title     string
	Caption   string
	PhotoURL  string
	UserID    uint
	User      User `gorm:"foreignkey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreatePhoto(db *gorm.DB, photo *Photo) error {
	return db.Create(photo).Error
}

func GetPhotoByID(db *gorm.DB, photoID uint) (*Photo, error) {
	var photo Photo
	err := db.Preload("User").Where("id = ?", photoID).First(&photo).Error
	if err != nil {
		return nil, err
	}
	return &photo, nil
}

func GetPhotos(db *gorm.DB) ([]Photo, error) {
	var photos []Photo
	err := db.Preload("User").Find(&photos).Error
	if err != nil {
		return nil, err
	}
	return photos, nil
}

func UpdatePhoto(db *gorm.DB, photo *Photo) error {
	return db.Save(photo).Error
}

func DeletePhoto(db *gorm.DB, photo *Photo) error {
	return db.Delete(photo).Error
}
