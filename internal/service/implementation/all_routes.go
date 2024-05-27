package serviceimpl

import (
	"context"
	"github.com/Dima191/route-server/internal/models"
)

func (i *impl) AllRoutes(ctx context.Context) (chan models.Route, chan error) {
	return i.rep.AllRoutes(ctx)
}
