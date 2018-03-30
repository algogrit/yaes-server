package api_test

import (
	"os"
	"testing"

	api "github.com/gauravagarwalr/Yet-Another-Expense-Splitter/src/api"
	db "github.com/gauravagarwalr/Yet-Another-Expense-Splitter/src/config/db"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Suite")
}

var _ = BeforeSuite(func() {
	dbUrl := os.Getenv("DATABASE_URL")
	dbName := os.Getenv("DB_NAME")

	db.InitializeDB("test", dbUrl, dbName)
	api.InitializeRouter()
})

var _ = AfterSuite(func() {
	db.Instance().Close()
})
