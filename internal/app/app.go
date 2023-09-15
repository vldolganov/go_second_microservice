package app

import (
	"github.com/sirupsen/logrus"
	"github.com/vldolganov/go_second_microservice/internal/config"
	"github.com/vldolganov/go_second_microservice/internal/server"
)

func Run(){
	log := logrus.New()
	cfg, err := config.MustLoad(log)

	if err != nil {
		return
	}

	server.Start(log, cfg.PORT)
}