package model

import "apiserver/pkg/constvar"

type UserModel struct{
    BaseModel
    Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
    Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

func (c *UserModel) TableName()string  {
	return "tb_users"
}

func (u *UserModel) Create()error  {
	return DB.Self.Create(&u).Error
}

func DeleteUser(id uint64)error  {
	user:=UserModel{}
	user.BaseModel.Id = id
	return DB.Self.Delete(&user).Error
}

func (u *UserModel) Update()error  {
	return DB.Self.Save(&u).Error
}

func GetUser(username string)(*UserModel,error)  {
	u:=UserModel{}
	d:=DB.Self.Where("username = ?",username).First(&u)
	return u,d.Error
}

func ListUser(username string,offset,limit int)([]*UserModel,uint64,error)  {
	if limit == 0{
		limit=constvar.DefaultLimit
	}
}