package category

import "github.com/rachadiannovansyah/go-echo-boilerplate-ipj/models"

func (categoryService *Service) Create(category *models.Category) {
	categoryService.DB.Create(category)
}
