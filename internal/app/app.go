package app

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/ysomad/pgxatomic"

	"github.com/ysomad/avito-internship-task/internal"
	"github.com/ysomad/avito-internship-task/internal/config"
	v1 "github.com/ysomad/avito-internship-task/internal/handler/v1"
	"github.com/ysomad/avito-internship-task/internal/postgres"
	"github.com/ysomad/avito-internship-task/internal/service"

	"github.com/ysomad/avito-internship-task/internal/pkg/httpserver"
	"github.com/ysomad/avito-internship-task/internal/pkg/logger"
	"github.com/ysomad/avito-internship-task/internal/pkg/pgclient"
)

func Run(conf *config.Config) {
	log := logger.New(
		conf.App.Ver,
		logger.WithLevel(conf.Log.Level),
		logger.WithLocation(conf.Log.TimeLoc),
		logger.WithSkipFrameCount(1),
	)

	// dependencies
	pgClient, err := pgclient.New(conf.PG.URL, pgclient.WithMaxConns(conf.PG.MaxConns))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer pgClient.Close()

	atomicRunner, err := pgxatomic.NewRunner(pgClient.Pool, pgx.TxOptions{})
	if err != nil {
		log.Fatal(err.Error())
	}

	atomicPool := pgxatomic.NewPool(pgClient.Pool)

	// services
	accountRepo := postgres.NewAccountRepo(atomicPool, pgClient.Builder)
	revenueReportRepo := postgres.NewRevenueReportRepo(atomicPool, pgClient.Builder)
	reservationRepo := postgres.NewReservationRepo(atomicPool, pgClient.Builder)
	transactionRepo := postgres.NewTransactionRepo(log, atomicPool, pgClient.Builder)

	accountService := service.NewAccount(log, accountRepo, revenueReportRepo, reservationRepo, transactionRepo)

	// init handlers
	mux := chi.NewMux()
	mux.Use(middleware.Logger, middleware.Recoverer)

	handlerV1 := v1.NewHandler(log, atomicRunner, accountService)
	v1.HandlerFromMuxWithBaseURL(handlerV1, mux, "/v1")

	runHTTPServer(mux, log, conf.HTTP.Port)
}

func runHTTPServer(mux http.Handler, log internal.Logger, port string) {
	log.Infof("starting http server on port %s", port)

	httpServer := httpserver.New(mux, httpserver.WithPort(port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Infof("received signal from httpserver: %s", s.String())
	case err := <-httpServer.Notify():
		log.Infof("got error from http server notify %s", err.Error())
	}

	if err := httpServer.Shutdown(); err != nil {
		log.Infof("got error on http server shutdown %s", err.Error())
	}
}
