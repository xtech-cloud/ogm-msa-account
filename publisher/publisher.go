package publisher

import (
	"context"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	proto "github.com/xtech-cloud/omo-msp-account/proto/account"
)

var DefaultPublisher micro.Event

func Publish(_notification *proto.Notification, _ctx context.Context) {
	err := DefaultPublisher.Publish(_ctx, _notification)
	if nil != err {
		logger.Error(err)
	}
}