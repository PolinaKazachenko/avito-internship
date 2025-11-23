package main

import (
	"context"
	"log"

	"avito-internship/internal/app/handlers"
	pull_requests2 "avito-internship/internal/app/handlers/pull_requests"
	teams2 "avito-internship/internal/app/handlers/teams"
	users2 "avito-internship/internal/app/handlers/users"
	"avito-internship/internal/app/middlewares"
	"avito-internship/internal/app/repository/tx_facade/pull_requests"
	"avito-internship/internal/app/repository/tx_facade/teams"
	"avito-internship/internal/app/repository/tx_facade/users"
	"avito-internship/internal/app/services"
	"avito-internship/internal/config"
	"avito-internship/internal/logger"
	"avito-internship/internal/postgres"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	var err error
	var sugarLogger *zap.SugaredLogger

	ctx, sugarLogger, err = createContextLogger(ctx)
	if err != nil {
		log.Fatalf("can't create logger: %s", err)
	}

	var cfg *config.Postgres
	cfg, err = config.FromEnv()
	if err != nil {
		logger.Fatalf(ctx, "can't get config from env: %s", err.Error())
	}

	var conn *pgxpool.Pool
	conn, err = postgres.NewClient(ctx, cfg)
	if err != nil {
		logger.Fatalf(ctx, "can't get Postgres client: %s", err.Error())
	}

	userFacade := users.NewTxFacade(conn)
	teamFacade := teams.NewTxFacade(conn)
	pullRequestFacade := pull_requests.NewTxFacade(conn)

	userService := services.NewUserService(userFacade, pullRequestFacade)
	teamService := services.NewTeamService(teamFacade, userFacade)
	pullRequestService := services.NewPullRequestService(pullRequestFacade, userFacade)

	loggerMiddleware := middlewares.NewLogger(sugarLogger)
	router := gin.Default()
	router.Use(loggerMiddleware.LoggerMiddleware())
	handlers.RegisterHandlers(
		router,
		users2.NewHandler(userService),
		teams2.NewHandler(teamService),
		pull_requests2.NewHandler(pullRequestService),
	)
	if err = router.Run(); err != nil {
		logger.Fatalf(ctx, "can't start server: %s", err.Error())
	}
}

func createContextLogger(ctx context.Context) (context.Context, *zap.SugaredLogger, error) {
	l, err := logger.NewProductionLogger()
	if err != nil {
		return nil, nil, err
	}
	return logger.WithLogger(ctx, l), l, nil
}
