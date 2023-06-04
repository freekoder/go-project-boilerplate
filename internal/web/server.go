package web

import (
	"context"
	"fmt"
	"github.com/freekoder/go-project-boilerplate/internal/config"
	"github.com/freekoder/go-project-boilerplate/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type Server struct {
	Service *service.Service
	ctx     context.Context
	cfg     config.WebConfig
	log     *zap.Logger
	e       *echo.Echo
}

func New(ctx context.Context, cfg config.WebConfig, log *zap.Logger, service *service.Service) *Server {
	e := configureEcho(log)
	configureRoutes(e, service)
	return &Server{ctx: ctx, e: e, cfg: cfg, log: log, Service: service}
}

func configureRoutes(e *echo.Echo, service *service.Service) {
	handler := &Handler{Service: service}
	g := e.Group("/api/v1")
	g.GET("/", handler.Index)
}

func (s *Server) Run() error {
	return s.e.Start(fmt.Sprintf(":%d", s.cfg.Port))
}

func (s *Server) Shutdown() error {
	ctx := context.Background()
	return s.e.Shutdown(ctx)
}

func configureEcho(log *zap.Logger) *echo.Echo {
	e := echo.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info("request",
				zap.String("uri", v.URI),
				zap.Int("status", v.Status),
			)
			return nil
		},
	}))
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	return e
}
