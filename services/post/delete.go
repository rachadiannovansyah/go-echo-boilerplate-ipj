package post

import "github.com/khihadysucahyo/go-echo-boilerplate/models"

func (postService *Service) Delete(post *models.Post) {
	postService.DB.Delete(post)
}
