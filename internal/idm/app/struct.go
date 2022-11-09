package app

import (
	"github.com/borisbbtest/GoMon/internal/idm/models"
	pb "github.com/borisbbtest/GoMon/internal/idm/proto/idm"
)

type GRPC struct {
	pb.UnimplementedEventsServer
	App *models.ConfigWrapper
}
