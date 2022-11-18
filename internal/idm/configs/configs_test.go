package configs_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/borisbbtest/GoMon/internal/idm/configs"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
)

type IdmConfigSuite struct {
	suite.Suite
	config *configs.AppConfig
}

func (suite *IdmConfigSuite) SetupSuite() {
	zerolog.SetGlobalLevel(zerolog.Disabled)
	err := os.Setenv("DATABASE_DSN", "dbtest")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Setenv("ADDRESS_GRPC", "rpctest")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Setenv("SESSION_TIME", "5m")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Setenv("REINIT", "false")
	if err != nil {
		log.Fatal(err)
	}
	flags := []string{"-d", "flagtest", "-g", "flagtest", "-r", "false", "-t", "10m"}
	os.Args = append(os.Args[:1], flags...)
}

func (suite *IdmConfigSuite) SetupTest() {
	suite.config = &configs.AppConfig{}
}

func (suite *IdmConfigSuite) TestYamlReadCorrect() {
	err := suite.config.YamlRead("test_data/idm1.yaml")
	suite.NoError(err)
	suite.Equal("127.0.0.1:8080", suite.config.ServerAddressGRPC)
	suite.Equal("postgres://pi:toor@192.168.1.69:5432/yandex", suite.config.DBDSN)
	suite.Equal(false, suite.config.ReInit)
	suite.Equal(300*time.Second, suite.config.SessionTimeExpired)
}

func (suite *IdmConfigSuite) TestYamlReadWrongDuration() {
	err := suite.config.YamlRead("test_data/idm2.yaml")
	suite.Error(err)
}

func (suite *IdmConfigSuite) TestYamlReadWrongBool() {
	err := suite.config.YamlRead("test_data/idm3.yaml")
	suite.Error(err)
}

func (suite *IdmConfigSuite) TestEnvReadCorrect() {
	err := suite.config.EnvRead()
	suite.NoError(err)
	suite.Equal("rpctest", suite.config.ServerAddressGRPC)
	suite.Equal("dbtest", suite.config.DBDSN)
	suite.Equal(false, suite.config.ReInit)
	suite.Equal(300*time.Second, suite.config.SessionTimeExpired)
}

func (suite *IdmConfigSuite) TestEnvReadWrongDuration() {
	err := os.Setenv("SESSION_TIME", "300")
	if err != nil {
		log.Fatal(err)
	}
	err = suite.config.EnvRead()
	suite.Error(err)
}

func (suite *IdmConfigSuite) TestEnvReadWrongBool() {
	err := os.Setenv("REINIT", "test")
	if err != nil {
		log.Fatal(err)
	}
	err = suite.config.EnvRead()
	suite.Error(err)
}

func (suite *IdmConfigSuite) TestLoadAppConfig() {
	err := os.Setenv("DATABASE_DSN", "common")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Setenv("ADDRESS_GRPC", "rpctest")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Setenv("SESSION_TIME", "")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Setenv("REINIT", "false")
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := configs.LoadAppConfig("test_data/idm1.yaml")
	suite.NoError(err)
	suite.Equal("common", cfg.DBDSN)
	suite.Equal(600*time.Second, cfg.SessionTimeExpired)
}

func TestReadConfigIdm(t *testing.T) {
	suite.Run(t, new(IdmConfigSuite))
}
