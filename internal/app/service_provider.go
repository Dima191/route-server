package app

import (
	"context"
	"github.com/Dima191/route-server/internal/repository"
	postgresrepository "github.com/Dima191/route-server/internal/repository/postgres"
	"github.com/Dima191/route-server/internal/service"
	serviceimpl "github.com/Dima191/route-server/internal/service/implementation"
	"log/slog"
)

type serviceProvider struct {
	rep  repository.Repository
	serv service.Service

	logger *slog.Logger
}

func (sp *serviceProvider) Repository(ctx context.Context, connectionStr string) (repository.Repository, error) {
	if sp.rep == nil {
		rep, err := postgresrepository.New(ctx, connectionStr, sp.logger)
		if err != nil {
			return nil, err
		}
		sp.rep = rep
	}

	return sp.rep, nil
}

func (sp *serviceProvider) Service(ctx context.Context, connectionStr string) (service.Service, error) {
	if sp.serv == nil {
		rep, err := sp.Repository(ctx, connectionStr)
		if err != nil {
			return nil, err
		}
		sp.serv = serviceimpl.New(rep, sp.logger)
	}

	return sp.serv, nil
}

func newServiceProvider(logger *slog.Logger) *serviceProvider {
	return &serviceProvider{
		logger: logger,
	}
}
