package etcd_test

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/course-extended-golang/users/storages"
	"github.com/course-extended-golang/users/storages/etcd"
	"github.com/mwitkow/go-etcd-harness"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"testing"
	"time"
)

var (
	etcdStorage storages.Storage
	etddHarness *etcd_harness.Harness
)

func TestEtcdDB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Etcd Suite")
}

var _ = BeforeSuite(func() {
	if !etcd_harness.LocalEtcdAvailable() {
		Fail("etcd is not available in $PATH, skipping suite")
	}

	harness, err := etcd_harness.New(os.Stderr)
	if err != nil {
		Fail(fmt.Sprintf("failed starting etcd harness: %v", err))
	}
	etddHarness = harness
	//etcdStorage = etcd.New("localhost:2379")
	time.Sleep(1 * time.Second)
	etcdStorage = etcd.New(etddHarness.Endpoint)
})

var _ = AfterSuite(func() {
	log.Debugf("cleaned up etcd harness")
	etddHarness.Stop()
})

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, TimestampFormat: time.RFC3339})
}
