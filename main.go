package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"omo-msa-account/config"
	"omo-msa-account/handler"
	"omo-msa-account/model"
	"omo-msa-account/publisher"
	"os"
	"path/filepath"
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
		micro.Version(BuildVersion),
		micro.RegisterTTL(time.Second*time.Duration(config.Schema.Service.TTL)),
		micro.RegisterInterval(time.Second*time.Duration(config.Schema.Service.Interval)),
		micro.Address(config.Schema.Service.Address),
	)

	// Initialise service
	service.Init()

	// Register publisher
	publisher.DefaultPublisher = micro.NewPublisher("omo.msa.account.notification", service.Client())
	// Register Handler
	proto.RegisterAuthHandler(service.Server(), new(handler.Auth))
	proto.RegisterProfileHandler(service.Server(), new(handler.Profile))
	proto.RegisterQueryHandler(service.Server(), new(handler.Query))

	app, _ := filepath.Abs(os.Args[0])

	logger.Info("-------------------------------------------------------------")
	logger.Info("- Micro Service Agent -> Run")
	logger.Info("-------------------------------------------------------------")
	logger.Infof("- version      : %s", BuildVersion)
	logger.Infof("- application  : %s", app)
	logger.Infof("- md5          : %s", md5hex(app))
	logger.Infof("- build        : %s", BuildTime)
	logger.Infof("- commit       : %s", CommitID)
	logger.Info("-------------------------------------------------------------")
	// Run service
	if err := service.Run(); err != nil {
		logger.Error(err)
	}
}

func md5hex(_file string) string {
	h := md5.New()

	f, err := os.Open(_file)
	if err != nil {
		return ""
	}
	defer f.Close()

	io.Copy(h, f)

	return hex.EncodeToString(h.Sum(nil))
}
