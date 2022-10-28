package main

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/ysomad/avito-internship-task/internal/app"
	"github.com/ysomad/avito-internship-task/internal/config"
)

func main() {
	var conf config.Config
	if err := cleanenv.ReadConfig("./configs/local.yml", &conf); err != nil {
		log.Fatalf("config parse error: %s", err)
	}

	app.Run(&conf)
}
