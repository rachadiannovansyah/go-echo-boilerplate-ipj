package post

import "github.com/rachadiannovansyah/go-echo-boilerplate-ipj/models"

func (postService *Service) Delete(post *models.Post) {
	postService.DB.Delete(post)
}
