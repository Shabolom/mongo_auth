package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"test_hh_1/config"
	"test_hh_1/init/routes"
	"test_hh_1/tools"
)

func main() {
	config.CheckFlagEnv()

	err := tools.InitLogger()
	if err != nil {
		fmt.Println("ошибка при инициализаии логера")
	}

	err = config.ConnectionDB()
	if err != nil {
		log.WithField("component", "DB").Panic(err)
		panic(err)
	}

	r := routes.SetupRouter()

	// запуск сервера
	if err = r.Run(config.Env.Host + ":" + config.Env.Port); err != nil {
		log.WithField("component", "run").Panic(err)
	}
}
