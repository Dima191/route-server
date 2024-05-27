package api

import (
	"github.com/Dima191/route-server/internal/service"
	pb "github.com/Dima191/route-server/pkg"
)

type Implementation struct {
	pb.UnimplementedRouteServer
	s service.Service
}

func New(s service.Service) *Implementation {
	return &Implementation{
		s: s,
	}
}
