package models_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/borisbbtest/GoMon/internal/idm/configs"
	mock_database "github.com/borisbbtest/GoMon/internal/idm/database/mock"
	"github.com/borisbbtest/GoMon/internal/idm/models"
	"github.com/borisbbtest/GoMon/internal/idm/service"
	"github.com/borisbbtest/GoMon/internal/models/idm"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

var ctrl *gomock.Controller

type ModelsTestSuite struct {
	suite.Suite
	Cfg *models.ConfigWrapper
}

func (suite *ModelsTestSuite) SetupSuite() {
	ctrl = gomock.NewController(suite.T())
	db := mock_database.NewMockStorager(ctrl)
	user := &idm.User{
		Login:    "test",
		Password: "$2a$08$mZv6cNuIXaD3tpXkkBTntelwFMUBAZtDZfnO9HnR876JMbEyO012m",
	}
	user1 := &idm.User{
		Login:     "test1",
		Firstname: "testfirstname",
		Lastname:  "testlastname",
		Password:  "testpassword",
		Source:    "testsource",
	}
	session1 := &idm.Session{
		Id:    "test_id1",
		Login: "test_login1",
	}
	session2 := &idm.Session{
		Id:    "test_id2",
		Login: "test_login2",
	}
	// user2 := &idm.User{
	// 	Login:     "test2",
	// 	Firstname: "testfirstname",
	// 	Lastname:  "testlastname",
	// 	Password:  "testpassword",
	// 	Source:    "testsource",
	// }
	firstCallGetSession := db.EXPECT().GetSession(context.Background(), &configs.AppConfig{}, "test_login1", "test_id1").Return(session1, nil)
	db.EXPECT().GetSession(context.Background(), &configs.AppConfig{}, "test_login1", "test_id2").Return(nil, service.ErrEmptySQLResult).After(firstCallGetSession)
	firstCallGetAllSessions := db.EXPECT().GetAllSessions(context.Background(), &configs.AppConfig{}).Return([]*idm.Session{session1, session2}, nil)
	db.EXPECT().GetAllSessions(context.Background(), &configs.AppConfig{}).Return(nil, errors.New("test")).After(firstCallGetAllSessions)
	db.EXPECT().CreateSession(context.Background(), &configs.AppConfig{}, gomock.Any()).Return(nil)
	db.EXPECT().DeleteSession(context.Background(), &configs.AppConfig{}, "test_login1", "test_id1").Return(nil)
	firstCallCreationSession := db.EXPECT().CreateUser(context.Background(), &configs.AppConfig{}, user1).Return(nil)
	db.EXPECT().CreateUser(context.Background(), &configs.AppConfig{}, user1).Return(service.ErrUserExists).After(firstCallCreationSession)
	db.EXPECT().DeleteUser(context.Background(), &configs.AppConfig{}, "test1").Return(nil)
	firstCallGetUser := db.EXPECT().GetUser(context.Background(), &configs.AppConfig{}, "test1").Return(user1, nil)
	db.EXPECT().GetUser(context.Background(), &configs.AppConfig{}, "test2").Return(nil, service.ErrEmptySQLResult).After(firstCallGetUser)
	db.EXPECT().GetAllUsers(context.Background(), &configs.AppConfig{}).Return([]*idm.User{{Login: "test1"}, {Login: "test2"}}, nil)
	firstCallAuthorization := db.EXPECT().GetUser(context.Background(), &configs.AppConfig{}, "test").Return(user, nil)
	secondCallAuthorization := db.EXPECT().CreateSession(context.Background(), &configs.AppConfig{}, gomock.Any()).Return(nil).After(firstCallAuthorization)
	db.EXPECT().GetUser(context.Background(), &configs.AppConfig{}, "test1").Return(user1, nil).After(secondCallAuthorization)
	firstCallRegistration := db.EXPECT().CreateUser(context.Background(), &configs.AppConfig{SessionTimeExpired: 10 * time.Second}, &idm.User{Login: "test1", Password: "testpassword"}).Return(nil)
	secondCallRegistration := db.EXPECT().CreateSession(context.Background(), &configs.AppConfig{SessionTimeExpired: 10 * time.Second}, gomock.Any()).Return(nil).After(firstCallRegistration)
	db.EXPECT().CreateUser(context.Background(), gomock.Any(), gomock.Any()).Return(service.ErrUserExists).After(secondCallRegistration)

	cw := &models.ConfigWrapper{
		Cfg:  &configs.AppConfig{},
		Repo: db,
	}
	suite.Cfg = cw
}

func (suite *ModelsTestSuite) TestGetSession() {
	ctx := context.Background()
	session, err := suite.Cfg.GetSession(ctx, "test_login1", "test_id1")
	suite.NoError(err)
	suite.Equal(&idm.Session{
		Id:    "test_id1",
		Login: "test_login1",
	}, session)
	session, err = suite.Cfg.GetSession(ctx, "test_login1", "test_id2")
	suite.Nil(session)
	suite.ErrorIs(err, service.ErrEmptySQLResult)
}

func (suite *ModelsTestSuite) TestGetAllSession() {
	ctx := context.Background()
	sessions, err := suite.Cfg.GetAllSessions(ctx)
	suite.NoError(err)
	suite.Equal([]*idm.Session{{
		Id:    "test_id1",
		Login: "test_login1",
	}, {
		Id:    "test_id2",
		Login: "test_login2",
	}}, sessions)
	sessions, err = suite.Cfg.GetAllSessions(ctx)
	suite.Error(err)
	suite.Nil(sessions)
}

func (suite *ModelsTestSuite) TestCreateSession() {
	ctx := context.Background()
	session, err := suite.Cfg.CreateSession(ctx, &idm.User{
		Login:     "test1",
		Firstname: "testfirstname",
		Lastname:  "testlastname",
		Password:  "testpassword",
		Source:    "testsource",
	})
	suite.NoError(err)
	suite.Equal("test1", session.Login)
	suite.Less(time.Now().Add(-1*time.Second), session.Created.AsTime())
}

func (suite *ModelsTestSuite) TestDeleteSession() {
	ctx := context.Background()
	err := suite.Cfg.DeleteSession(ctx, "test_login1", "test_id1")
	suite.NoError(err)
}

func (suite *ModelsTestSuite) TestCreationUser() {
	ctx := context.Background()
	err := suite.Cfg.CreateUser(ctx, &idm.User{
		Login:     "test1",
		Firstname: "testfirstname",
		Lastname:  "testlastname",
		Password:  "testpassword",
		Source:    "testsource",
	})
	suite.NoError(err)
	err = suite.Cfg.CreateUser(ctx, &idm.User{
		Login:     "test1",
		Firstname: "testfirstname",
		Lastname:  "testlastname",
		Password:  "testpassword",
		Source:    "testsource",
	})
	suite.ErrorIs(err, service.ErrUserExists)
}

func (suite *ModelsTestSuite) TestDeleteUser() {
	ctx := context.Background()
	err := suite.Cfg.DeleteUser(ctx, "test1")
	suite.NoError(err)
}

func (suite *ModelsTestSuite) TestGetUser() {
	ctx := context.Background()
	user, err := suite.Cfg.GetUser(ctx, &idm.User{Login: "test1"})
	suite.NoError(err)
	suite.Equal(&idm.User{
		Login:     "test1",
		Firstname: "testfirstname",
		Lastname:  "testlastname",
		Password:  "testpassword",
		Source:    "testsource",
	}, user)
	user, err = suite.Cfg.GetUser(ctx, &idm.User{Login: "test2"})
	suite.ErrorIs(err, service.ErrEmptySQLResult)
	suite.Nil(user)
}

func (suite *ModelsTestSuite) TestGetAllUsers() {
	ctx := context.Background()
	users, err := suite.Cfg.GetAllUsers(ctx)
	suite.NoError(err)
	suite.Equal([]*idm.User{{
		Login: "test1",
	}, {
		Login: "test2",
	}}, users)
}

func (suite *ModelsTestSuite) TestAuthorization() {
	ctx := context.Background()
	session, err := suite.Cfg.AuthorizeUser(ctx, "test", "testpassword")
	suite.NoError(err)
	suite.Equal("test", session.Login)
	session, err = suite.Cfg.AuthorizeUser(ctx, "test1", "test")
	suite.ErrorIs(err, service.ErrWrongPassword)
	suite.Nil(session)
}

func (suite *ModelsTestSuite) TestRegistration() {
	ctx := context.Background()
	suite.Cfg.Cfg.SessionTimeExpired = 10 * time.Second
	session, err := suite.Cfg.RegisterUser(ctx, &idm.User{Login: "test1", Password: "testpassword"})
	suite.NoError(err)
	suite.Equal("test1", session.Login)
	suite.Less(session.Duration.AsTime(), time.Now().Add(10*time.Second))
	session, err = suite.Cfg.RegisterUser(ctx, &idm.User{Login: "test1", Password: "testpassword"})
	suite.ErrorIs(err, service.ErrUserExists)
	suite.Nil(session)
}

func (suite *ModelsTestSuite) TearDownSuite() {
	ctrl.Finish()
}

func TestModels(t *testing.T) {
	suite.Run(t, new(ModelsTestSuite))
}
