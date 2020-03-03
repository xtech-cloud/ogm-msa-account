package main

import (
	"omo-msa-account/config"
	"omo-msa-account/handler"
	"omo-msa-account/model"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	proto "github.com/xtech-cloud/omo-msp-account/proto/account"
)

func main() {
	config.Setup()
	model.Setup()
	model.AutoMigrateDatabase()

	// New Service
	service := micro.NewService(
		micro.Name("omo.msa.account"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*time.Duration(config.Schema.Service.TTL)),
		micro.RegisterInterval(time.Second*time.Duration(config.Schema.Service.Interval)),
		micro.Address(config.Schema.Service.Address),
	)

	// Initialise service
	service.Init()

	// Register Handler
	proto.RegisterAuthHandler(service.Server(), new(handler.Auth))
	proto.RegisterProfileHandler(service.Server(), new(handler.Profile))

	// Run service
	if err := service.Run(); err != nil {
		logger.Error(err)
	}
}
