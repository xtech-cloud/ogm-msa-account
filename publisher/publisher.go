package publisher

import (
	"context"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	proto "github.com/xtech-cloud/omo-msp-notification/proto/notification"
)

var DefaultPublisher micro.Event

func Publish(_ctx context.Context, _action string, _head string, _body string) {
	err := DefaultPublisher.Publish(_ctx, &proto.SimpleMessage{
		Action: _action,
		Head:   _head,
		Body:   _body,
	})
	if nil != err {
		logger.Error(err)
	}
}
