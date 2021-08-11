package responses

import "github.com/rachadiannovansyah/go-echo-boilerplate-ipj/models"

type CategoryResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   string `json:"username"`
	ID          uint   `json:"id"`
}

func NewCategoryResponse(categories []models.Category) *[]CategoryResponse {
	categoryResponse := make([]CategoryResponse, 0)

	for i := range categories {
		categoryResponse = append(categoryResponse, CategoryResponse{
			Name:        categories[i].Name,
			Description: categories[i].Description,
		})
	}

	return &categoryResponse
}

func NewCategoryReponseOne(category models.Category) *CategoryResponse {
	categoryReponse := CategoryResponse{
		Name:        category.Name,
		Description: category.Description,
	}

	return &categoryReponse
}
