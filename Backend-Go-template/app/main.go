package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/Nxwbtk/Mono-repo-template/Backend-Go-template/config"
	docs "github.com/Nxwbtk/Mono-repo-template/Backend-Go-template/docs"
	"github.com/Nxwbtk/Mono-repo-template/Backend-Go-template/security"
	cat "github.com/Nxwbtk/Mono-repo-template/Backend-Go-template/services/Cat"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// @BasePath /api/v1

const (
	ServerReadHeaderTimeout = 5 * time.Second
	ServerReadTimeout       = 5 * time.Second
	ServerWriteTimeout      = 10 * time.Second
	handlerTimeout          = ServerWriteTimeout - (time.Millisecond * 100)
)

const (
	gracefulShutdownDuration = 10 * time.Second
)

// @title           Template API documentation
// @version         1.0
// @description     This is a backend server for the template project
// @termsOfService  http://swagger.io/terms/
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	zap, _ := zap.NewProduction()
	defer zap.Sync()
	r, stop := router()
	defer stop()

	srv := &http.Server{
		Addr:              fmt.Sprint(":", "2324"),
		Handler:           r,
		ReadHeaderTimeout: ServerReadHeaderTimeout,
		ReadTimeout:       ServerReadTimeout,
		WriteTimeout:      ServerWriteTimeout,
		MaxHeaderBytes:    1 << 20,
	}

	go gracefully(srv, zap, gracefulShutdownDuration)

	zap.Info("run at :" + "2324")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed to start server: %v", err)
	}

	log.Println("server exited gracefully")
}

func gracefully(srv *http.Server, log *zap.Logger, shutdownTimeout time.Duration) {
	{
		ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer cancel()
		<-ctx.Done()
	}

	log.Info("initiating graceful shutdown", zap.Duration("timeout(s)", shutdownTimeout))
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Info("HTTP server Shutdown: " + err.Error())
	}
}

func router() (*gin.Engine, func()) {
	db, instance, err := config.CreateClientDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	r := gin.Default()

	// config.Migrate(db)
	docs.SwaggerInfo.BasePath = "/api/v1"
	catRepository := cat.NewCatRepo(db)
	catService := cat.NewCatService(catRepository)
	catHandler := cat.NewCatHandler(catService)
	api := r.Group("/api/v1/")
	{
		catGroup := api.Group("/cat").Use(security.Middleware())
		{
			catHandler.RegisterCatrouters(catGroup)
		}
	}
	if config.NewConfig().ENV == "dev" || config.NewConfig().ENV == "uat" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	return r, func() {
		instance.Close()
	}
}
