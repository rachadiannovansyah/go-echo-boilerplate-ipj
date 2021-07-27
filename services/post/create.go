package post

import "github.com/rachadiannovansyah/go-echo-boilerplate-ipj/models"

func (postService *Service) Create(post *models.Post) {
	postService.DB.Create(post)
}
