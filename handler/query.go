package handler

import (
	"context"
	"ogm-msa-account/model"

	"github.com/asim/go-micro/v3/logger"
	proto "github.com/xtech-cloud/ogm-msp-account/proto/account"
)

type Query struct{}

func (this *Query) List(_ctx context.Context, _req *proto.QueryListRequest, _rsp *proto.QueryListResponse) error {
	logger.Infof("Received Query.List, req is %v", _req)
	_rsp.Status = &proto.Status{}

	offset := int64(0)
	if _req.Offset > 0 {
		offset = _req.Offset
	}

	count := int64(0)
	if _req.Count > 0 {
		count = _req.Count
	}

	dao := model.NewAccountDAO(nil)
	accounts, err := dao.List(offset, count)
	if nil != err {
		return err
	}

	total, err := dao.Count()
	if nil != err {
		return err
	}
	_rsp.Total = total

	_rsp.Account = make([]*proto.AccountEntity, len(accounts))
	for i := 0; i < len(accounts); i++ {
		_rsp.Account[i] = &proto.AccountEntity{
			Username:  accounts[i].Username,
			Uuid:      accounts[i].UUID,
			CreatedAt: accounts[i].Embedded.CreatedAt.Unix(),
			UpdatedAt: accounts[i].Embedded.UpdatedAt.Unix(),
		}
	}
	return nil
}

func (this *Query) Single(_ctx context.Context, _req *proto.QuerySingleRequest, _rsp *proto.QuerySingleResponse) error {
	logger.Infof("Received Query.Single, req is %v", _req)
	_rsp.Status = &proto.Status{}

	var err error
	var account model.Account
	dao := model.NewAccountDAO(nil)
	if proto.QueryField_QUERY_FIELD_UUID == _req.Field {
		account, err = dao.Find(_req.Value)
	} else if proto.QueryField_QUERY_FIELD_USERNAME == _req.Field {
		account, err = dao.WhereUsername(_req.Value)
	}

	if err != nil {
		return err
	}

	if account.UUID == "" {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "not fount"
		return nil
	}

	_rsp.Account = &proto.AccountEntity{
		Username:  account.Username,
		Uuid:      account.UUID,
		Profile:   account.Profile,
		CreatedAt: account.Embedded.CreatedAt.Unix(),
		UpdatedAt: account.Embedded.UpdatedAt.Unix(),
	}
	return nil
}
