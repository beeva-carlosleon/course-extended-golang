package etcd_test

import (
	log "github.com/Sirupsen/logrus"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"time"
	"github.com/course-extended-golang/users/storages"
	"github.com/course-extended-golang/users/storages/etcd"
)

var etcdStorage storages.Storage

func TestEtcdDB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Etcd Suite")
}

var _ = BeforeSuite(func() {
	etcdStorage = etcd.New("localhost:2379")
})

var _ = AfterSuite(func() {
	log.Debugf("cleaned up etcd harness")
})

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, TimestampFormat: time.RFC3339})
}
