package handler

import (
	"context"
	"omo-msa-account/model"

	"github.com/micro/go-micro/v2/logger"
	proto "github.com/xtech-cloud/omo-msp-account/proto/account"
)

type Profile struct{}

func (this *Profile) Query(_ctx context.Context, _req *proto.QueryProfileRequest, _rsp *proto.QueryProfileResponse) error {
	logger.Infof("Received Profile.Query, accessToken is %v", _req.AccessToken)
	_rsp.Status = &proto.Status{}
	if "" == _req.AccessToken {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "accessToken is required"
		return nil
	}

	dao := model.NewAccountDAO()

	uuid, err := useridFromToken(_req.AccessToken, _req.Strategy)
	if nil != err {
		return err
	}
	account, err := dao.Find(uuid)
	if nil != err {
		return err
	}
	if "" == account.UUID {
		_rsp.Status.Code = 2
		_rsp.Status.Message = "account not found"
		return nil
	}
	_rsp.Profile = account.Profile
	return nil
}

func (this *Profile) Update(_ctx context.Context, _req *proto.UpdateProfileRequest, _rsp *proto.UpdateProfileResponse) error {
	logger.Infof("Received Profile.Update, accessToken is %v", _req.AccessToken)
	_rsp.Status = &proto.Status{}
	dao := model.NewAccountDAO()

	if "" == _req.AccessToken {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "accessToken is required"
		return nil
	}

	uuid, err := useridFromToken(_req.AccessToken, _req.Strategy)
	if nil != err {
		return err
	}
	account, err := dao.Find(uuid)
	if nil != err {
		return err
	}
	if "" == account.UUID {
		_rsp.Status.Code = 2
		_rsp.Status.Message = "account not found"
		return nil
	}
	return dao.UpdateProfile(uuid, _req.Profile)
}
