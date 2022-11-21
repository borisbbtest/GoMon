package configs_test

import (
	"log"
	"os"
	"testing"

	"github.com/borisbbtest/GoMon/internal/fanin/configs"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
)

type FaninConfigSuite struct {
	suite.Suite
	config *configs.AppConfig
}

func (suite *FaninConfigSuite) SetupSuite() {
	zerolog.SetGlobalLevel(zerolog.Disabled)
	err := os.Setenv("HTTP_ADDRESS", ":7443")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Setenv("IDM_ADDRESS", ":8071")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Setenv("CMDB_ADDRESS", ":8072")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Setenv("METRICS_ADDRESS", ":8073")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Setenv("EVENTS_ADDRESS", ":8074")
	if err != nil {
		log.Fatal(err)
	}
	flags := []string{"-s", ":443", "-i", ":8081", "-m", ":8083", "-e", ":8084"}
	os.Args = append(os.Args[:1], flags...)
}

func (suite *FaninConfigSuite) SetupTest() {
	suite.config = &configs.AppConfig{}
}

func (suite *FaninConfigSuite) TestYamlReadCorrect() {
	err := suite.config.YamlRead("test_data/fanin1.yaml")
	suite.NoError(err)
	suite.Equal(":8443", suite.config.HTTPServerAddress)
	suite.Equal(":8091", suite.config.IdmAddress)
	suite.Equal(":8092", suite.config.CmdbAddress)
	suite.Equal(":8093", suite.config.MetricsAddress)
	suite.Equal(":8094", suite.config.EventsAddress)
}

func (suite *FaninConfigSuite) TestEnvReadCorrect() {
	err := suite.config.EnvRead()
	suite.NoError(err)
	suite.Equal(":7443", suite.config.HTTPServerAddress)
	suite.Equal(":8071", suite.config.IdmAddress)
	suite.Equal(":8072", suite.config.CmdbAddress)
	suite.Equal(":8073", suite.config.MetricsAddress)
	suite.Equal(":8074", suite.config.EventsAddress)
}

func (suite *FaninConfigSuite) TestLoadAppConfig() {
	err := os.Setenv("HTTP_ADDRESS", ":2443")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Setenv("IDM_ADDRESS", "")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Setenv("CMDB_ADDRESS", "")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Setenv("METRICS_ADDRESS", "")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Setenv("EVENTS_ADDRESS", ":2084")
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := configs.LoadAppConfig("test_data/fanin2.yaml")
	suite.config = cfg
	suite.NoError(err)
	suite.Equal(":2443", suite.config.HTTPServerAddress)
	suite.Equal(":8081", suite.config.IdmAddress)
	suite.Equal(":1082", suite.config.CmdbAddress)
	suite.Equal(":8083", suite.config.MetricsAddress)
	suite.Equal(":2084", suite.config.EventsAddress)
}

func TestReadConfigIdm(t *testing.T) {
	suite.Run(t, new(FaninConfigSuite))
}
