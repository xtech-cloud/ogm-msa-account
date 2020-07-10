package handler

import (
	"context"

	"omo-msa-account/model"
	"omo-msa-account/publisher"

	"github.com/micro/go-micro/v2/logger"

	proto "github.com/xtech-cloud/omo-msp-account/proto/account"
)

type Auth struct{}

func (this *Auth) Signup(_ctx context.Context, _req *proto.SignupRequest, _rsp *proto.SignupResponse) error {
	logger.Infof("Received Auth.Signup, username is %v", _req.Username)

	_rsp.Status = &proto.Status{}

	if "" == _req.Username {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "username is required"
		return nil
	}

	if "" == _req.Password {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "password is required"
		return nil
	}

	dao := model.NewAccountDAO()

	// 账号存在检测
	exists, err := dao.Exists(_req.Username)
	// 数据库错误
	if nil != err {
		return err
	}

	if exists {
		_rsp.Status.Code = 2
		_rsp.Status.Message = "account exists"
		return nil
	}

	uuid := model.NewUUID()
	account := model.Account{
		UUID:     uuid,
		Username: _req.Username,
		Password: dao.GeneratePassword(_req.Password, _req.Username),
		Profile:  "",
	}
	err = dao.Insert(account)
	if nil != err {
		return err
	}

	// 无错误
	_rsp.Uuid = uuid

	// 发布消息
	ctx := buildNotifyContext(_ctx, uuid)
	publisher.Publish(&proto.Notification{
		Action: "/signup",
		Head:   "",
		Body:   uuid,
	}, ctx)
	return nil
}

func (this *Auth) Signin(_ctx context.Context, _req *proto.SigninRequest, _rsp *proto.SigninResponse) error {
	logger.Infof("Received Auth.Signin, username is %v, strategy is %v", _req.Username, _req.Strategy)
	_rsp.Status = &proto.Status{}

	if "" == _req.Username {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "username is required"
		return nil
	}

	if "" == _req.Password {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "password is required"
		return nil
	}

	dao := model.NewAccountDAO()

	username := _req.Username

	account, err := dao.WhereUsername(username)
	if nil != err {
		return err
	}

	if account.UUID == "" {
		_rsp.Status.Code = 2
		_rsp.Status.Message = "account not found"
		return nil
	}

	err = dao.VerifyPassword(_req.Password, _req.Username, account.Password)
	if nil != err {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "password not matched"
		return nil
	}
	if proto.Strategy_JWT == proto.Strategy(_req.Strategy) {
		token, err := tokenFromJWT(account.UUID)
		if nil != err {
			return nil
		}
		_rsp.AccessToken = token
	} else {
		_rsp.AccessToken = account.UUID
	}
	_rsp.Uuid = account.UUID
	// 发布消息
	ctx := buildNotifyContext(_ctx, account.UUID)
	publisher.Publish(&proto.Notification{
		Action: "/signin",
		Head:   _rsp.AccessToken,
		Body:   _rsp.Uuid,
	}, ctx)
	return nil
}

func (this *Auth) Signout(_ctx context.Context, _req *proto.SignoutRequest, _rsp *proto.SignoutResponse) error {
	logger.Infof("Received Auth.Signout, accessToken is %v", _req.AccessToken)
	_rsp.Status = &proto.Status{}

	uuid, err := useridFromToken(_req.AccessToken, _req.Strategy)
	if nil != err {
		return err
	}

	// 发布消息
	ctx := buildNotifyContext(_ctx, uuid)
	publisher.Publish(&proto.Notification{
		Action: "/signout",
		Head:   _req.AccessToken,
		Body:   uuid,
	}, ctx)
	return nil
}

func (this *Auth) ResetPasswd(_ctx context.Context, _req *proto.ResetPasswdRequest, _rsp *proto.ResetPasswdResponse) error {
	logger.Infof("Received Auth.ResetPasswd, accessToken is %v", _req.AccessToken)
	_rsp.Status = &proto.Status{}

	if "" == _req.AccessToken {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "accessToken is required"
		return nil
	}

	if "" == _req.Password {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "password is required"
		return nil
	}

	dao := model.NewAccountDAO()

	uuid, err := useridFromToken(_req.AccessToken, _req.Strategy)
	if nil != err {
		return err
	}

	logger.Infof("uuid is %v", uuid)

	account, err := dao.Find(uuid)
	if nil != err {
		return err
	}
	if account.UUID == "" {
		_rsp.Status.Code = 2
		_rsp.Status.Message = "account not found"
		return nil
	}

	password := dao.GeneratePassword(_req.Password, account.Username)
	err = dao.UpdatePassword(uuid, password)
	if nil != err {
		return err
	}

	// 发布消息
	ctx := buildNotifyContext(_ctx, uuid)
	publisher.Publish(&proto.Notification{
		Action: "/reset/password",
		Head:   _req.AccessToken,
		Body:   uuid,
	}, ctx)
	return nil
}
