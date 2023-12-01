package service

import (
	"context"
	logging "github.com/sirupsen/logrus"
	"gmall/conf"
	"gmall/consts"
	"gmall/pkg/e"
	"gmall/pkg/utils"
	"gmall/repository/db/dao"
	model2 "gmall/repository/db/model"
	"gmall/serializer"
)

// UserService 管理用户服务
type UserService struct {
	NickName string `form:"nick_name" json:"nick-name"`
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
	Key      string `form:"key" json:"key"`
}

func (service UserService) Register(c context.Context) serializer.Response {
	var user *model2.User
	code := e.SUCCESS
	if service.Key == "" || len(service.Key) != 16 {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密钥长度不足",
		}
	}
	utils.Encrypt.SetKey(service.Key)
	userDao := dao.NewUserDao(c)
	_, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if exist {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user = &model2.User{
		NickName: service.NickName,
		UserName: service.UserName,
		Status:   model2.Active,
		Money:    utils.Encrypt.AesEncoding("100"), // 初始金额
	}
	// 加密密码
	if err = user.SetPassword(service.Password); err != nil {
		logging.Info(err)
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if conf.UploadModel == consts.UploadModelOss {
		user.Avatar = "\"http://q1.qlogo.cn/g?b=qq&nk=294350394&s=640"
	} else {
		user.Avatar = "avatar.JPG"
	}
	// 创建用户
	err = userDao.CreateUser(user)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// Login 用户登陆
func (service *UserService) Login(c context.Context) serializer.Response {
	var user *model2.User
	code := e.SUCCESS
	userDao := dao.NewUserDao(c)
	user, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if !exist { // 如果查不到，返回相应的错误
		logging.Info(err)
		code = e.ErrorUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}

	}
	if user.CheckPassword(service.Password) == false {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	token, err := utils.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		logging.Info(err)
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    e.GetMsg(code),
	}
}
