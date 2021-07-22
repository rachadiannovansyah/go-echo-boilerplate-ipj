package post

import "github.com/khihadysucahyo/go-echo-boilerplate/models"

func (postService *Service) Create(post *models.Post) {
	postService.DB.Create(post)
}
