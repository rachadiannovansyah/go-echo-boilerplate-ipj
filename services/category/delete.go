package category

import "github.com/rachadiannovansyah/go-echo-boilerplate-ipj/models"

func (categoryService *Service) Delete(category *models.Category) {
	categoryService.DB.Delete(category)
}
