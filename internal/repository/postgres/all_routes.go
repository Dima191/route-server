package postgresrepository

import (
	"context"
	"errors"
	"github.com/Dima191/route-server/internal/models"
	"github.com/Dima191/route-server/internal/repository"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"sync"
)

const chanBuffer = 5

func (r *rep) AllRoutes(ctx context.Context) (chan models.Route, chan error) {
	routesCh := make(chan models.Route, chanBuffer)
	errCh := make(chan error, 1)

	query := "select domain, host, port from route"

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		r.logger.Error("failed to make a query to database", slog.String("error", err.Error()))
		errCh <- err
		close(errCh)
		close(routesCh)

		return nil, errCh
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		for rows.Next() {
			route := models.Route{}
			if err = rows.Scan(&route.Domain, &route.Host, &route.Port); err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					errCh <- repository.ErrNoRows
					return
				}
				r.logger.Error("failed to scan row", slog.String("error", err.Error()))
				errCh <- err
				return
			}
			routesCh <- route
		}
	}()

	go func() {
		defer close(routesCh)
		defer close(errCh)

		wg.Wait()
		rows.Close()

		if err = rows.Err(); err != nil {
			r.logger.Error("failed to scan rows", slog.String("error", err.Error()))
			errCh <- err
			return
		}
	}()

	return routesCh, errCh
}
