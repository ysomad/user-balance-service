package main

import (
	"github.com/ysomad/avito-internship-task/internal/app"
	"github.com/ysomad/avito-internship-task/internal/config"
)

func main() {
	var conf config.Config
	app.Run(&conf)
}
