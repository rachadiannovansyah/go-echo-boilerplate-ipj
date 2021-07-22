package post

import (
	"github.com/khihadysucahyo/go-echo-boilerplate/models"
	"github.com/khihadysucahyo/go-echo-boilerplate/requests"

	"github.com/jinzhu/gorm"
)

type ServiceWrapper interface {
	Create(post *models.Post)
	Delete(post *models.Post)
	Update(post *models.Post, updatePostRequest *requests.UpdatePostRequest)
}

type Service struct {
	DB *gorm.DB
}

func NewPostService(db *gorm.DB) *Service {
	return &Service{DB: db}
}
