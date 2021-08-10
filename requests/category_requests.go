package requests

import validation "github.com/go-ozzo/ozzo-validation"

type CategoryValidation struct {
	Name        string `json:"name" validate:"required" example:"Categories"`
	Description string `json:"description" validate:"required" example:"Description Category"`
}

func (cv CategoryValidation) Validate() error {
	return validation.ValidateStruct(&cv,
		validation.Field(&cv.Name, validation.Required),
		validation.Field(&cv.Description, validation.Required),
	)
}

type CreateCategoryRequest struct {
	CategoryValidation
}

type UpdateCategoryRequest struct {
	CategoryValidation
}
