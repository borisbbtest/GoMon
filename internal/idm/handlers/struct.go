package handlers

import (
	"github.com/borisbbtest/GoMon/internal/idm/models"
	pb "github.com/borisbbtest/GoMon/internal/idm/proto/idm"
)

type GRPC struct {
	pb.UnimplementedIdmServer
	App *models.ConfigWrapper
}
