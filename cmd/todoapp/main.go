package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/wander1ch/to_do_list_app/internal/core/logger"
	core_postgres_pool "github.com/wander1ch/to_do_list_app/internal/core/repository/postgres/pool"
	core_http_middleware "github.com/wander1ch/to_do_list_app/internal/core/transport/http/middleware"
	core_http_server "github.com/wander1ch/to_do_list_app/internal/core/transport/http/server"
	users_postgres_repository "github.com/wander1ch/to_do_list_app/internal/features/users/repository/postgres"
	users_service "github.com/wander1ch/to_do_list_app/internal/features/users/service"
	users_transport_http "github.com/wander1ch/to_do_list_app/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	defer cancel()

	fmt.Println("Starting To-Do List App...")

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Printf("failed to initialize logger: %v", err)
		os.Exit(1)
	}
	defer logger.Close()


	pool, err := core_postgres_pool.NewConnectionPool(
		ctx,
		core_postgres_pool.NewConfigMust(),
	)
	if err != nil {
		logger.Fatal("failed to initialize Postgres pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("Logger initialized successfully", zap.String("environment", "development"))
	usersRepo := users_postgres_repository.NewUserRepository(pool)
	usersService := users_service.NewUserService(usersRepo)
	usersTransportHTTP := users_transport_http.NewUserHTTPHandler(usersService)

	logger.Debug("User HTTP handler initialized successfully")

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Panic(),
		core_http_middleware.Tracing(),
	)

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.APIVersionV1)
	apiVersionRouter.RegisterRoutes(usersTransportHTTP.Routes())

	httpServer.RegisterAPIRouters(apiVersionRouter)
	
	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server encountered an error", zap.Error(err))
	}
}
		
	