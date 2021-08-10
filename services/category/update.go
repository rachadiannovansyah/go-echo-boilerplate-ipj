package category

import (
	"github.com/rachadiannovansyah/go-echo-boilerplate-ipj/models"
	"github.com/rachadiannovansyah/go-echo-boilerplate-ipj/requests"
)

func (categoryService *Service) Update(category *models.Category, UpdateCategoryRequest *requests.UpdateCategoryRequest) {
	category.Name = UpdateCategoryRequest.Name
	category.Description = UpdateCategoryRequest.Description
	categoryService.DB.Save(category)
}
