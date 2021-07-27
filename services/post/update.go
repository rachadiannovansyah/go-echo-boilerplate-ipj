package post

import (
	"github.com/rachadiannovansyah/go-echo-boilerplate-ipj/models"

	"github.com/rachadiannovansyah/go-echo-boilerplate-ipj/requests"
)

func (postService *Service) Update(post *models.Post, updatePostRequest *requests.UpdatePostRequest) {
	post.Content = updatePostRequest.Content
	post.Title = updatePostRequest.Title
	postService.DB.Save(post)
}
