package repository

import (
	"context"
	"github.com/Dima191/route-server/internal/models"
)

type Repository interface {
	AllRoutes(ctx context.Context) (chan models.Route, chan error)
	Close()
}
