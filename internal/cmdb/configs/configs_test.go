package configs_test

import (
	"log"
	"os"
	"testing"

	"github.com/borisbbtest/GoMon/internal/cmdb/configs"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
)

type CmdbConfigSuite struct {
	suite.Suite
	config *configs.AppConfig
}

func (suite *CmdbConfigSuite) SetupSuite() {
	zerolog.SetGlobalLevel(zerolog.Disabled)
	err := os.Setenv("DATABASE_DSN", "dbtest")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Setenv("ADDRESS_GRPC", "rpctest")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Setenv("REINIT", "false")
	if err != nil {
		log.Fatal(err)
	}
	flags := []string{"-d", "flagtest", "-g", "flagtest", "-r", "false"}
	os.Args = append(os.Args[:1], flags...)
}

func (suite *CmdbConfigSuite) SetupTest() {
	suite.config = &configs.AppConfig{}
}

func (suite *CmdbConfigSuite) TestYamlReadCorrect() {
	err := suite.config.YamlRead("test_data/cmdb1.yaml")
	suite.NoError(err)
	suite.Equal("127.0.0.1:8080", suite.config.ServerAddressGRPC)
	suite.Equal("postgres://pi:toor@192.168.1.69:5432/yandex", suite.config.DBDSN)
	suite.Equal(false, suite.config.ReInit)
}

func (suite *CmdbConfigSuite) TestYamlReadWrongBool() {
	err := suite.config.YamlRead("test_data/cmdb3.yaml")
	suite.Error(err)
}

func (suite *CmdbConfigSuite) TestEnvReadCorrect() {
	err := suite.config.EnvRead()
	suite.NoError(err)
	suite.Equal("rpctest", suite.config.ServerAddressGRPC)
	suite.Equal("dbtest", suite.config.DBDSN)
	suite.Equal(false, suite.config.ReInit)
}

func (suite *CmdbConfigSuite) TestEnvReadWrongDuration() {
	err := os.Setenv("SESSION_TIME", "300")
	if err != nil {
		log.Fatal(err)
	}
	err = suite.config.EnvRead()
	suite.Error(err)
}

func (suite *CmdbConfigSuite) TestEnvReadWrongBool() {
	err := os.Setenv("REINIT", "test")
	if err != nil {
		log.Fatal(err)
	}
	err = suite.config.EnvRead()
	suite.Error(err)
}

func (suite *CmdbConfigSuite) TestLoadAppConfig() {
	err := os.Setenv("DATABASE_DSN", "common")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Setenv("ADDRESS_GRPC", "rpctest")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Setenv("REINIT", "false")
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := configs.LoadAppConfig("test_data/cmdb1.yaml")
	suite.NoError(err)
	suite.Equal("common", cfg.DBDSN)
}

func TestReadConfigIdm(t *testing.T) {
	suite.Run(t, new(CmdbConfigSuite))
}
