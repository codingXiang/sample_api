package delivery

import "github.com/gin-gonic/gin"

type HttpHandler interface {
	GetUserList(c *gin.Context) error
	GetUser(c *gin.Context) error
	CreateUser(c *gin.Context) error
	ModifyUser(c *gin.Context) error
	DeleteUser(c *gin.Context) error
}
