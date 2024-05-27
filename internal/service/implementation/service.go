package serviceimpl

import (
	"github.com/Dima191/route-server/internal/repository"
	"github.com/Dima191/route-server/internal/service"
	"log/slog"
)

type impl struct {
	rep repository.Repository

	logger *slog.Logger
}

func New(rep repository.Repository, logger *slog.Logger) service.Service {
	return &impl{
		rep:    rep,
		logger: logger,
	}
}
