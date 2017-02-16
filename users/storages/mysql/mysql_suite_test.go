package mysql_test

import (
	log "github.com/Sirupsen/logrus"
	"github.com/course-extended-golang/users/storages"
	"github.com/course-extended-golang/users/storages/mysql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"time"
)

var mysqlStorage storages.Storage

func TestMysql(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mysql Suite")
}

var _ = BeforeSuite(func() {
	mysqlStorage = mysql.New()
})

var _ = AfterSuite(func() {
	log.Debugf("Closing connections with database")
	if err := mysqlStorage.Close(); err != nil {
		log.Error(err)
	}
})

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, TimestampFormat: time.RFC3339})
}
