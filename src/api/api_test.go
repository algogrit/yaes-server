package api_test

import (
	"testing"

	"github.com/caarlos0/env"
	api "algogrit.com/yaes-server/src/api"
	db "algogrit.com/yaes-server/src/config/db"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
	dbcleaner "gopkg.in/khaiql/dbcleaner.v2"
	"gopkg.in/khaiql/dbcleaner.v2/engine"
)

var Cleaner = dbcleaner.New()

func TestAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Suite")
}

func BenchmarkAPI(b *testing.B) {
	RegisterFailHandler(Fail)
	RunSpecs(b, "API Suite")
}

type Config struct {
	DBName string `env:"DB_NAME" envDefault:"yaes-test"`
	DBUrl  string `env:"DATABASE_URL"`
	AppEnv string `env:"GO_APP_ENV" envDefault:"test"`
	Port   string `env:"PORT" envDefault:"12345"`
}

var cfg Config

var _ = BeforeSuite(func() {
	env.Parse(&cfg)
	dbConnectionString := db.GetConnectionString(cfg.DBUrl, cfg.DBName)

	postgres := engine.NewPostgresEngine(dbConnectionString)
	Cleaner.SetEngine(postgres)

	log.Info("Setup DB Cleaner")

	db.InitializeDB("test", cfg.DBUrl, cfg.DBName)
	api.InitializeRouter("test")
})

var _ = AfterSuite(func() {
	db.Instance().Close()
})
