package handlers

import (
	"github.com/borisbbtest/GoMon/internal/idm/models"
	pb "github.com/borisbbtest/GoMon/internal/models/idm"
)

type GRPC struct {
	pb.UnimplementedIdmServer
	App *models.ConfigWrapper
}
