package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ridhohrnf/task-5-vix-btpns-Muhammad-Ridho-Heranof/helpers"
	"github.com/ridhohrnf/task-5-vix-btpns-Muhammad-Ridho-Heranof/models"
)

type PhotoController struct {
	db *gorm.DB
}

func NewPhotoController(db *gorm.DB) *PhotoController {
	return &PhotoController{db}
}

func (pc *PhotoController) CreatePhoto(c *gin.Context) {
	// Handle photo creation logic
}

func (pc *PhotoController)
