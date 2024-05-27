package service

import (
	"context"
	"github.com/Dima191/route-server/internal/models"
)

type Service interface {
	AllRoutes(context.Context) (chan models.Route, chan error)
}
