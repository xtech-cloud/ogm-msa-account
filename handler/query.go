package handler

import (
	"context"
	"omo-msa-account/model"

	"github.com/micro/go-micro/v2/logger"
	proto "github.com/xtech-cloud/omo-msp-account/proto/account"
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

	dao := model.NewAccountDAO()
	accounts, err := dao.List(offset, count)
	if nil != err {
		return err
	}
	logger.Infof("has %v accounts", len(accounts))

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
