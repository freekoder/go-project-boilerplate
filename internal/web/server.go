package web

import (
	"context"
	"crypto/subtle"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	ctx context.Context
	e   *echo.Echo
}

func New(ctx context.Context) *Server {
	e := configureEcho()
	return &Server{ctx: ctx, e: e}
}

func (s *Server) Run() error {
	// TODO: fix server port constant with config value
	// TODO: select http/https version of server based on config value
	return s.e.Start(fmt.Sprintf(":%d", 9090))
}

func (s *Server) Shutdown() error {
	ctx := context.Background()
	return s.e.Shutdown(ctx)
}

func configureEcho() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// TODO: fix test/test auth with config value
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("test")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("test")) == 1 {
			return true, nil
		}
		return false, nil
	}))

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	return e
}
