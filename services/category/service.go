package category

import (
	"github.com/jinzhu/gorm"
	"github.com/rachadiannovansyah/go-echo-boilerplate-ipj/models"
	"github.com/rachadiannovansyah/go-echo-boilerplate-ipj/requests"
)

type ServiceWrapper interface {
	Create(category *models.Category)
	Update(category *models.Category, updatePostRequest *requests.UpdateCategoryRequest)
	Delete(category *models.Category)
}

type Service struct {
	DB *gorm.DB
}

func NewCategoryService(db *gorm.DB) *Service {
	return &Service{DB: db}
}
