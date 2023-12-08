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
	"mime/multipart"
)

// UserService 管理用户服务
type UserService struct {
	NickName string `form:"nick_name" json:"nick_name"`
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
	Key      string `form:"key" json:"key"`
}

func (service *UserService) Register(c context.Context) serializer.Response {
	var user *model2.User
	code := e.SUCCESS
	if service.Key == "" || len(service.Key) != 12 {
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

// Post 更新用户头像
func (service *UserService) Post(c context.Context, uId uint, file *multipart.FileHeader) serializer.Response {
	code := e.SUCCESS
	var user *model2.User
	var err error
	userDao := dao.NewUserDao(c)
	user, err = userDao.GetUserById(uId)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	path, err := utils.OssUpload(file)
	if err != nil {
		code = e.ErrorUploadFile
		return serializer.Response{
			Status: code,
			Data:   e.GetMsg(code),
			Error:  path,
		}
	}
	user.Avatar = path
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUser(user),
		Msg:    e.GetMsg(code),
	}
}

// Update 用户修改信息
func (service *UserService) Update(c context.Context, uId uint) serializer.Response {
	var user *model2.User
	var err error
	code := e.SUCCESS
	// 找到用户
	userDao := dao.NewUserDao(c)
	user, err = userDao.GetUserById(uId)
	if err != nil {
		return ErrorResponse(err)
	}
	if service.NickName != "" {
		user.NickName = service.NickName
	}
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		return ErrorResponse(err)
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUser(user),
		Msg:    e.GetMsg(code),
	}
}

// ErrorResponse 错误响应
func ErrorResponse(err error) serializer.Response {
	logging.Info(err)
	code := e.ErrorDatabase
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Error:  err.Error(),
	}
}
