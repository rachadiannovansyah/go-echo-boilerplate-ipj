package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/rachadiannovansyah/go-echo-boilerplate-ipj/models"
)

type CategoryRepositoryQ interface {
	GetCategories(categories *[]models.Category)
	GetCategory(category *models.Category, id int)
}

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (categoryRepository *CategoryRepository) GetCategories(categories *[]models.Category) {
	categoryRepository.DB.Find(categories)
}

func (categoryRepository CategoryRepository) GetCategory(category *models.Category, id int) {
	categoryRepository.DB.Where("id = ?", id).Find(category)
}
