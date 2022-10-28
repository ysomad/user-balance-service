package app

import (
	"github.com/ysomad/avito-internship-task/internal/config"

	"github.com/ysomad/avito-internship-task/logger"
	"github.com/ysomad/avito-internship-task/pgclient"
)

func Run(conf *config.Config) {
	log := logger.New(
		conf.App.Ver,
		logger.WithLevel(conf.Log.Level),
		logger.WithLocation(conf.Log.TimeLoc),
		logger.WithSkipFrameCount(1),
	)

	pg, err := pgclient.New(conf.PG.URL, pgclient.WithMaxConns(conf.PG.MaxConns))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer pg.Close()

}
