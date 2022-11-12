package models_test

import (
	"context"
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
	firstCallCreation := db.EXPECT().CreateUser(context.Background(), &configs.AppConfig{}, user1).Return(nil)
	db.EXPECT().CreateUser(context.Background(), &configs.AppConfig{}, user1).Return(service.ErrUserExists).After(firstCallCreation)
	db.EXPECT().GetSession(context.Background(), &configs.AppConfig{}, "test_login1", "test_id1").Return(session1, nil)
	db.EXPECT().GetAllSessions(context.Background(), &configs.AppConfig{}).Return([]*idm.Session{session1, session2}, nil)
	db.EXPECT().CreateSession(context.Background(), &configs.AppConfig{}, gomock.Any()).Return(nil)
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

func (suite *ModelsTestSuite) TearDownSuite() {
	ctrl.Finish()
}

func TestModels(t *testing.T) {
	suite.Run(t, new(ModelsTestSuite))
}
