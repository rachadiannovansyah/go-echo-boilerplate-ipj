package handlers

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/rachadiannovansyah/go-echo-boilerplate-ipj/models"
	"github.com/rachadiannovansyah/go-echo-boilerplate-ipj/repositories"
	"github.com/rachadiannovansyah/go-echo-boilerplate-ipj/requests"
	"github.com/rachadiannovansyah/go-echo-boilerplate-ipj/responses"
	s "github.com/rachadiannovansyah/go-echo-boilerplate-ipj/server"
	categoryService "github.com/rachadiannovansyah/go-echo-boilerplate-ipj/services/category"
	"github.com/rachadiannovansyah/go-echo-boilerplate-ipj/services/token"
)

type CategoryHandlers struct {
	server *s.Server
}

func NewCategoryHandlers(server *s.Server) *CategoryHandlers {
	return &CategoryHandlers{server: server}
}

func (ch *CategoryHandlers) GetCategories(c echo.Context) error {
	var categories []models.Category

	categoryRepository := repositories.NewCategoryRepository(ch.server.DB)
	categoryRepository.GetCategories(&categories)

	for i := 0; i < len(categories); i++ {
		ch.server.DB.Model(&categories[i])
	}

	response := responses.NewCategoryResponse(categories)
	return responses.Response(c, http.StatusOK, response)
}

func (ch *CategoryHandlers) CreateCategory(c echo.Context) error {
	createCategoryRequest := new(requests.CreateCategoryRequest)

	if err := c.Bind(createCategoryRequest); err != nil {
		return err
	}

	if err := createCategoryRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty")
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*token.JwtCustomClaims)
	id := claims.ID

	category := models.Category{
		Name:        createCategoryRequest.Name,
		Description: createCategoryRequest.Description,
		UserID:      id,
	}

	categoryService := categoryService.NewCategoryService(ch.server.DB)
	categoryService.Create(&category)

	return responses.MessageResponse(c, http.StatusCreated, "Category has been successfully created!")
}

func (ch *CategoryHandlers) UpdateCategory(c echo.Context) error {
	updateCategoryRequest := new(requests.UpdateCategoryRequest)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(updateCategoryRequest); err != nil {
		return err
	}

	if err := updateCategoryRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty")
	}

	category := models.Category{}

	categoryRepository := repositories.NewCategoryRepository(ch.server.DB)
	categoryRepository.GetCategory(&category, id)

	if category.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Category not found")
	}

	categoryService := categoryService.NewCategoryService(ch.server.DB)
	categoryService.Update(&category, updateCategoryRequest)

	return responses.MessageResponse(c, http.StatusOK, "Category has been successfully updated!")
}

func (ch *CategoryHandlers) DeleteCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	category := models.Category{}

	categoryRepository := repositories.NewCategoryRepository(ch.server.DB)
	categoryRepository.GetCategory(&category, id)

	if category.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Category not found")
	}

	categoryService := categoryService.NewCategoryService(ch.server.DB)
	categoryService.Delete(&category)

	return responses.MessageResponse(c, http.StatusOK, "Category has been successfully deleted!")
}
