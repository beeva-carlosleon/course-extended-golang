package mysql_test

import (
	log "github.com/Sirupsen/logrus"
	"github.com/course-extended-golang/users"
	"github.com/course-extended-golang/users/storages"
	"github.com/course-extended-golang/users/storages/mysql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User checks", func() {
	var (
		user1      users.User
		err        error
		storage    storages.Storage
		deleteMark bool
	)
	BeforeEach(func() {
		user1 = users.User{Id: 1, Name: "Carlos", SurName: "León", Age: 38}
		storage = mysql.NewMock()
		deleteMark = false
	})
	AfterEach(func() {
		if deleteMark {
			err = mysqlStorage.Delete(user1)
			Expect(err).To(BeNil())
		}
	})
	Context("Insert user tests mock", func() {
		It("Insert user", func() {
			err = storage.Create(user1)
			Expect(err).NotTo(HaveOccurred())
		})
		It("Insert user error", func() {
			err = storage.Create(user1)
			Expect(err).NotTo(HaveOccurred())
			err = storage.Create(user1)
			log.Debugf("%v", err)
			Expect(err).To(HaveOccurred())
			storage.Close()
		})

	})
	Context("Insert user tests", func() {
		It("Insert user", func() {
			err = mysqlStorage.Create(user1)
			Expect(err).NotTo(HaveOccurred())
			deleteMark = true
		})
		It("Insert user error", func() {
			err = mysqlStorage.Create(user1)
			err = mysqlStorage.Create(user1)
			Expect(err).To(HaveOccurred())
			deleteMark = true
		})
	})
	Context("Delete user tests mock", func() {
		It("Delete user", func() {
			err = storage.Create(user1)
			Expect(err).NotTo(HaveOccurred())
			err = storage.Delete(user1)
			Expect(err).NotTo(HaveOccurred())
		})
		It("Delete user with no existent user", func() {
			err = storage.Delete(user1)
			log.Debugf("%v", err)
			Expect(err).NotTo(HaveOccurred())
		})
	})
	Context("Delete user tests", func() {
		It("Delete user", func() {
			err = mysqlStorage.Create(user1)
			Expect(err).NotTo(HaveOccurred())
			err = mysqlStorage.Delete(user1)
			Expect(err).NotTo(HaveOccurred())
		})
		It("Delete with no existent user", func() {
			user1.Id = 2
			err = storage.Create(user1)
			Expect(err).NotTo(HaveOccurred())
			user1.Id = 1
			err = storage.Delete(user1)
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
