package app

import (
	"context"
	"fmt"
	"github.com/Dima191/route-server/internal/api"
	"github.com/Dima191/route-server/internal/config"
	pb "github.com/Dima191/route-server/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
)

type App struct {
	cfg *config.Config

	grpc *grpc.Server

	sp   *serviceProvider
	impl *api.Implementation

	logger *slog.Logger
}

func (a *App) initConfig(configPath string) error {
	cfg, err := config.New(configPath)
	if err != nil {
		return err
	}

	a.cfg = cfg
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.sp = newServiceProvider(a.logger)
	return nil
}

func (a *App) initImplementation(ctx context.Context) error {
	serv, err := a.sp.Service(ctx, a.cfg.ConnectionStr)
	if err != nil {
		return err
	}
	a.impl = api.New(serv)

	return nil
}

func (a *App) initGrpcServer(_ context.Context) error {
	a.grpc = grpc.NewServer()

	reflection.Register(a.grpc)

	a.logger.Info("registration")
	pb.RegisterRouteServer(a.grpc, a.impl)

	return nil
}

func (a *App) initDeps(ctx context.Context, configPath string) error {
	//CONFIG INITIALIZATION
	if err := a.initConfig(configPath); err != nil {
		return err
	}

	deps := []func(context.Context) error{
		a.initServiceProvider,
		a.initImplementation,
		a.initGrpcServer,
	}

	for _, f := range deps {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) Run() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.cfg.AppPort))
	if err != nil {
		a.logger.Error("failed to create listener", slog.String("error", err.Error()))
		return err
	}
	a.logger.Info("route server started", slog.Int("port", a.cfg.AppPort))

	if err = a.grpc.Serve(listener); err != nil {
		a.logger.Error("failed to serve", slog.String("error", err.Error()))
		return err
	}

	return nil
}

func (a *App) Stop() {
	if a.sp.rep != nil {
		a.sp.rep.Close()
	}

	a.grpc.GracefulStop()
}

func New(ctx context.Context, configPath string, logger *slog.Logger) (*App, error) {
	a := &App{}
	a.logger = logger

	a.logger.Info("initializing dependencies")
	if err := a.initDeps(ctx, configPath); err != nil {
		a.logger.Error("failed to initialize dependencies", slog.String("error", err.Error()))
		return nil, err
	}
	a.logger.Info("all dependencies initialized")

	return a, nil
}
