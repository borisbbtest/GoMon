package models_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"

	"github.com/borisbbtest/GoMon/internal/cmdb/configs"
	mock_database "github.com/borisbbtest/GoMon/internal/cmdb/database/mock"
	"github.com/borisbbtest/GoMon/internal/cmdb/models"
	"github.com/borisbbtest/GoMon/internal/cmdb/service"
	pb "github.com/borisbbtest/GoMon/internal/models/cmdb"
)

var (
	ctrl *gomock.Controller
	ci1  = &pb.Ci{
		Name:        "ci1",
		Description: "ci1 for test",
		CreatedBy:   "suite",
		Type:        "virtual_node",
	}
	ci2 = &pb.Ci{
		Name:        "ci2",
		Description: "ci2 for test",
		CreatedBy:   "suite",
		Type:        "virtual_node",
	}
	ci3 = &pb.Ci{
		Name:        "ci3",
		Description: "ci3 for test",
		CreatedBy:   "suite",
		Type:        "virtual_node",
	}
)

type ModelsCmdbTestSuite struct {
	suite.Suite
	Cfg *models.ConfigWrapper
}

func (suite *ModelsCmdbTestSuite) SetupSuite() {
	zerolog.SetGlobalLevel(zerolog.Disabled)
	ctrl = gomock.NewController(suite.T())
	db := mock_database.NewMockStorager(ctrl)
	db.EXPECT().GetObject(gomock.Any(), gomock.Any(), "ci1").Return(ci1, nil).AnyTimes()
	db.EXPECT().GetObject(gomock.Any(), gomock.Any(), "ci2").Return(ci2, nil).AnyTimes()
	db.EXPECT().GetObject(gomock.Any(), gomock.Any(), "ci3").Return(nil, service.ErrEmptySQLResult).AnyTimes()
	firstCallCreate := db.EXPECT().CreateObject(gomock.Any(), gomock.Any(), ci1).Return(nil)
	db.EXPECT().CreateObject(gomock.Any(), gomock.Any(), ci1).Return(service.ErrObjectExists).After(firstCallCreate)
	db.EXPECT().CreateObject(gomock.Any(), gomock.Any(), ci2).Return(nil).AnyTimes()
	firstcallCreateCi3 := db.EXPECT().CreateObject(gomock.Any(), gomock.Any(), ci3).Return(nil)
	db.EXPECT().CreateObject(gomock.Any(), gomock.Any(), ci3).Return(service.ErrObjectExists).After(firstcallCreateCi3)
	firstCallDelete := db.EXPECT().DeleteObject(gomock.Any(), gomock.Any(), "ci1").Return(nil)
	db.EXPECT().DeleteObject(gomock.Any(), gomock.Any(), "ci1").Return(errors.New("test internal error")).After(firstCallDelete)
	db.EXPECT().DeleteObject(gomock.Any(), gomock.Any(), "ci2").Return(nil).AnyTimes()
	firstcallDeleteCi3 := db.EXPECT().DeleteObject(gomock.Any(), gomock.Any(), "ci3").Return(nil)
	db.EXPECT().DeleteObject(gomock.Any(), gomock.Any(), "ci3").Return(errors.New("test internal error")).After(firstcallDeleteCi3)

	cw := &models.ConfigWrapper{
		Cfg:  &configs.AppConfig{},
		Repo: db,
	}
	suite.Cfg = cw
}

func (suite *ModelsCmdbTestSuite) TestGetObject() {
	ctx := context.Background()
	ci, err := suite.Cfg.GetObject(ctx, "ci1")
	suite.NoError(err)
	suite.Equal(ci1, ci)
	ci, err = suite.Cfg.GetObject(ctx, "ci3")
	suite.Nil(ci)
	suite.ErrorIs(err, service.ErrEmptySQLResult)
}

func (suite *ModelsCmdbTestSuite) TestGetBatchObjects() {
	ctx := context.Background()
	cis, err := suite.Cfg.GetBatchObjects(ctx, []string{"ci1", "ci2"})
	suite.NoError(err)
	suite.Equal(2, len(cis))
	cis, err = suite.Cfg.GetBatchObjects(ctx, []string{"ci1", "ci3"})
	suite.ErrorIs(err, service.ErrSelectObjects)
	suite.Equal([]*pb.Ci{ci1}, cis)
}

func (suite *ModelsCmdbTestSuite) TestCreateObject() {
	ctx := context.Background()
	err := suite.Cfg.CreateObject(ctx, ci1)
	suite.NoError(err)
	err = suite.Cfg.CreateObject(ctx, ci1)
	suite.ErrorIs(err, service.ErrObjectExists)
}

func (suite *ModelsCmdbTestSuite) TestCreateBatchObjects() {
	ctx := context.Background()
	err := suite.Cfg.CreateBatchObjects(ctx, []*pb.Ci{ci2, ci3})
	suite.NoError(err)
	err = suite.Cfg.CreateBatchObjects(ctx, []*pb.Ci{ci2, ci3})
	suite.ErrorIs(err, service.ErrInsertObjects)
}

func (suite *ModelsCmdbTestSuite) TestDeleteObject() {
	ctx := context.Background()
	err := suite.Cfg.DeleteObject(ctx, "ci1")
	suite.NoError(err)
	err = suite.Cfg.DeleteObject(ctx, "ci1")
	suite.Error(err)
}

func (suite *ModelsCmdbTestSuite) TestDeleteBatchObject() {
	ctx := context.Background()
	err := suite.Cfg.DeleteBatchObject(ctx, []string{"ci2", "ci3"})
	suite.NoError(err)
	err = suite.Cfg.DeleteBatchObject(ctx, []string{"ci2", "ci3"})
	suite.ErrorIs(err, service.ErrDeleteObjects)
}

func (suite *ModelsCmdbTestSuite) TearDownSuite() {
	ctrl.Finish()
}

func TestModels(t *testing.T) {
	suite.Run(t, new(ModelsCmdbTestSuite))
}
