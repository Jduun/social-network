package handlers

import (
	"github.com/gin-gonic/gin"
)

type AuthHandlers interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	GetMe(c *gin.Context)
}

type UserHandlers interface {
	GetUser(c *gin.Context)
	GetUsers(c *gin.Context)
	GetUserFollowers(c *gin.Context)
	GetFollowingUsers(c *gin.Context)
}

type PostHandlers interface {
	GetPost(c *gin.Context)
	GetPosts(c *gin.Context)
	CreatePost(c *gin.Context)
	UpdatePost(c *gin.Context)
	DeletePost(c *gin.Context)
	GetFeed(c *gin.Context)
	CreatePostReaction(c *gin.Context)
	DeletePostReaction(c *gin.Context)
}

type FollowHandlers interface {
	Follow(c *gin.Context)
	Unfollow(c *gin.Context)
}
