package user

import (
	"apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	username := c.Param("username")
	user, err := model.GetUser(username)
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
	}
	handler.SendResponse(c, nil, user)
}
