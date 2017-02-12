package users_test

import (
	log "github.com/Sirupsen/logrus"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"time"
)

func TestRoutes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Users Suite")
}

var _ = BeforeSuite(func() {
	log.Debugf("This is executed before suite test")
})

var _ = AfterSuite(func() {
	log.Debugf("This is executed after suite test")
})

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, TimestampFormat: time.RFC3339})
}
