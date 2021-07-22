package post

import (
	"github.com/khihadysucahyo/go-echo-boilerplate/models"

	"github.com/khihadysucahyo/go-echo-boilerplate/requests"
)

func (postService *Service) Update(post *models.Post, updatePostRequest *requests.UpdatePostRequest) {
	post.Content = updatePostRequest.Content
	post.Title = updatePostRequest.Title
	postService.DB.Save(post)
}
