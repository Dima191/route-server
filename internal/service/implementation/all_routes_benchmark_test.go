package serviceimpl

import (
	"context"
	"github.com/Dima191/route-server/internal/models"
	"github.com/Dima191/route-server/internal/repository"
	"go.uber.org/mock/gomock"
	"log/slog"
	"os"
	"testing"
)

func BenchmarkAllRoutes(b *testing.B) {
	options := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, options))

	ctrl := gomock.NewController(b)
	mockedRep := repository.NewMockRepository(ctrl)

	serv := New(mockedRep, logger)

	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		b.StopTimer()
		routeCh := make(chan models.Route, 1)
		routeCh <- models.Route{
			Domain: "docs.test.ru",
			Host:   "127.0.0.1",
			Port:   "1234",
		}
		close(routeCh)

		mockedRep.EXPECT().AllRoutes(context.Background()).Return(routeCh, nil)
		b.StartTimer()

		_, _ = serv.AllRoutes(context.Background())
	}
}
