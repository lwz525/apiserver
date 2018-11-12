package user

import (
	"github.com/gin-gonic/gin"
	"apiserver/model"
	. "apiserver/handler"
	"apiserver/pkg/auth"
	"apiserver/pkg/token"
)

func Login(c *gin.Context)  {
	var u model.UserModel
	if err:=c.Bind(&u);err!=nil{
		SendResponse(c,err,nil)
		return
	}

	d,err:=model.GetUser(u.Username)
	if err!=nil{
		SendResponse(c,err,nil)
		return
	}
	if err:=auth.Compare(d.Password,u.Password);err!=nil{
		SendResponse(c,err,nil)
		return
	}
	t,err:=token.Sign(c,token.Context{ID:d.Id,Username:d.Username},"")
	if err!=nil{
		SendResponse(c,err,nil)
		return
	}
	SendResponse(c,nil,model.Token{Token:t})
}
