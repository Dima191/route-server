package api

import (
	"errors"
	"github.com/Dima191/route-server/internal/repository"
	pb "github.com/Dima191/route-server/pkg"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
	"os"
)

func (i *Implementation) AllRoutes(_ *emptypb.Empty, stream pb.Route_AllRoutesServer) error {
	options := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, options))

	routes, errCh := i.s.AllRoutes(stream.Context())
	select {
	case err := <-errCh:
		if err != nil {
			if errors.Is(err, repository.ErrNoRows) {
				return status.Error(codes.NotFound, err.Error())
			}

			logger.Error("api error", slog.String("error", err.Error()))
			return status.Error(codes.Internal, err.Error())
		}
	default:
	}

	for route := range routes {
		if err := stream.Send(&pb.AllRoutesResponse{
			Domain: route.Domain,
			Host:   route.Host,
			Port:   route.Port,
		}); err != nil {
			logger.Error("api error", slog.String("error", err.Error()))
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}
