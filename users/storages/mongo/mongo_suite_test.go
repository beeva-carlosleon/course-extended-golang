package mongo_test

import (
	log "github.com/Sirupsen/logrus"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"time"
	"github.com/course-extended-golang/users/storages"
	"github.com/course-extended-golang/users/storages/mongo"
)

var mongoStorage storages.Storage

func TestMongoDB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MongoDB Suite")
}

var _ = BeforeSuite(func() {
	mongoStorage = mongo.New()
})

var _ = AfterSuite(func() {
	log.Debugf("Closing connections with database")
	mongoStorage.Close()
})

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, TimestampFormat: time.RFC3339})
}
