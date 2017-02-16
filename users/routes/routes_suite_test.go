package routes_test

import (
	log "github.com/Sirupsen/logrus"
	"github.com/course-extended-golang/users/storages"
	"github.com/course-extended-golang/users/storages/mysql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"time"
)

func TestRoutes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Routes Suite")
}

var mockStorage storages.Storage

var _ = BeforeSuite(func() {
	mockStorage = mysql.NewMock()
})

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, TimestampFormat: time.RFC3339})
}
