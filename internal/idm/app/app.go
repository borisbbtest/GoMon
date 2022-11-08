package app

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/borisbbtest/GoMon/internal/idm/configs"
	"github.com/borisbbtest/GoMon/internal/idm/database"
	"github.com/borisbbtest/GoMon/internal/idm/models"
	pb "github.com/borisbbtest/GoMon/internal/idm/proto/idm"
	"github.com/borisbbtest/GoMon/internal/idm/service"
)

var log = zerolog.New(service.LogConfig()).With().Timestamp().Caller().Logger()

func BuildApp() {
	cfg := configs.LoadAppConfig()
	ctx := context.Background()
	fmt.Println(cfg)
	repo, err := database.NewDBStorage(ctx, cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	srv := models.AppWrapper{
		Cfg:  cfg,
		Repo: repo,
	}
	user := &pb.User{
		Login:     "test2",
		Lastname:  "Porubov",
		Password:  "password",
		CreatedAt: timestamppb.Now(),
		Source:    "manual",
	}
	ss, err := srv.RegisterUser(ctx, user)
	if err != nil {
		log.Error().Err(err).Msg("failed register user")
		return
	}
	log.Info().Msg(fmt.Sprint(ss))
	ss, err = srv.AuthorizeUser(ctx, user.Login, user.Password)
	if err != nil {
		log.Error().Err(err).Msg("failed authorize user")
		return
	}
	log.Info().Msg(fmt.Sprint(ss))
	user1 := &pb.User{
		Login: "test2",
	}
	userpb, err := srv.GetUser(ctx, user1)
	if err != nil {
		log.Error().Err(err).Msg("get all users")
	}
	fmt.Println(userpb)

	users, err := srv.GetAllUsers(ctx)
	if err != nil {
		log.Error().Err(err).Msg("get all users")
	}
	fmt.Println(users)
	sessions, err := srv.GetAllSessions(ctx)
	if err != nil {
		log.Error().Err(err).Msg("get all sessions")
	}
	fmt.Println(sessions)

}
