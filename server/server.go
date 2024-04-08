package server

import (
	"context"
	"fmt"
	"time"

	"github.com/cod3rboy/practice-cqrs/config"
	"github.com/gin-contrib/graceful"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type NewServerParams struct {
	fx.In
	LC     fx.Lifecycle
	Logger *zap.Logger
	Config config.Config
}

type Server struct {
	engine   *gin.Engine
	graceful *graceful.Graceful
}

func NewServer(params NewServerParams) *Server {
	server := &Server{engine: gin.New()}
	// middlewares
	server.engine.Use(ginzap.Ginzap(params.Logger, time.RFC3339, true))
	server.engine.Use(ginzap.RecoveryWithZap(params.Logger, true))

	params.LC.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				server.Run(params.Config.ServerPort)
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			return server.Stop()
		},
	})

	return server
}

func (s *Server) Run(port int) error {
	gracefulServer, err := graceful.New(s.engine, graceful.WithAddr(fmt.Sprintf(":%d", port)))
	if err != nil {
		return err
	}
	s.graceful = gracefulServer
	return s.graceful.Run()
}

func (s *Server) Stop() error {
	if s.graceful == nil {
		return nil
	}
	return s.graceful.Stop()
}

func (s *Server) Router() *gin.Engine {
	return s.engine
}
