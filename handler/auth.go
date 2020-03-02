package handler

import (
	"context"

	"omo-msa-account/model"

	"github.com/micro/go-micro/v2/logger"

	proto "github.com/xtech-cloud/omo-msp-account/proto/account"
)

type Auth struct{}

func (this *Auth) Signup(_ctx context.Context, _req *proto.SignupRequest, _rsp *proto.SignupResponse) error {
	logger.Infof("Received Auth.Signup, username is %v", _req.Username)
	_rsp.Status = &proto.Status{}
	dao := model.NewAccountDAO()

	// 账号存在检测
	exists, err := dao.Exists(_req.Username)
	// 数据库错误
	if nil != err {
		return err
	}

	if exists {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "account exists"
		return nil
	}

	uuid := model.NewUUID()
	account := model.Account{
		UUID:     uuid,
		Username: _req.Username,
		Password: dao.StrengthenPassword(_req.Password, _req.Username),
		Profile:  "",
	}
	err = dao.Insert(account)
	if nil != err {
		return err
	}

	// 无错误
	_rsp.Uuid = uuid
	return nil
}

func (this *Auth) Signin(_ctx context.Context, _req *proto.SigninRequest, _rsp *proto.SigninResponse) error {
	logger.Infof("Received Auth.Signin, username is %v", _req.Username)
	_rsp.Status = &proto.Status{}
	dao := model.NewAccountDAO()

	username := _req.Username
	password := dao.StrengthenPassword(_req.Password, _req.Username)

	account, err := dao.WhereUsername(username)
	if nil != err {
		return err
	}

	if account.UUID == "" {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "account not found"
		return nil
	}

	if account.Password != password {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "password not matched"
		return nil
	}
	if proto.Strategy_JWT == _req.Strategy {
		_rsp.AccessToken = account.UUID
	} else {
		_rsp.AccessToken = account.UUID
	}
	return nil
}

func (this *Auth) Signout(_ctx context.Context, _req *proto.SignoutRequest, _rsp *proto.SignoutResponse) error {
	logger.Infof("Received Auth.Signout, accessToken is %v", _req.AccessToken)
	_rsp.Status = &proto.Status{}
	return nil
}

func (this *Auth) ResetPasswd(_ctx context.Context, _req *proto.ResetPasswdRequest, _rsp *proto.ResetPasswdResponse) error {
	logger.Infof("Received Auth.ResetPasswd, accessToken is %v", _req.AccessToken)
	_rsp.Status = &proto.Status{}
	dao := model.NewAccountDAO()

	uuid := takeUUID(_req.AccessToken)

	account, err := dao.Find(uuid)
	if nil != err {
		return err
	}
	if account.UUID == "" {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "account not found"
		return nil
	}

	password := dao.StrengthenPassword(_req.Password, account.Username)
	return dao.UpdatePassword(uuid, password)
}
