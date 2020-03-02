package handler

import (
	proto "github.com/xtech-cloud/omo-msp-account/proto/account"
)

func takeUUID(_accessToken string) string {
	strategy := proto.Strategy_NONE
	if proto.Strategy_JWT == strategy {
		return _accessToken
	}
	return _accessToken
}
