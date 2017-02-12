package etcd_test

import (
	log "github.com/Sirupsen/logrus"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/course-extended-golang/users"
	"github.com/course-extended-golang/users/storages"
	"github.com/course-extended-golang/users/storages/etcd"
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
		storage = etcd.NewMock()
		deleteMark = false
	})
	AfterEach(func() {
		if deleteMark {
			err = etcdStorage.Delete(user1)
			Expect(err).To(BeNil())
		}
	})
	Context("Insert user tests mock", func() {
		It("Insert user", func() {
			err = storage.Create(user1)
			Expect(err).To(BeNil())
		})
		It("Insert user error", func() {
			err = storage.Create(user1)
			Expect(err).To(BeNil())
			err = storage.Create(user1)
			log.Debugf("%v", err)
			Expect(err).NotTo(BeNil())
			storage.Close()
		})

	})
	Context("Insert user tests", func() {
		It("Insert user", func() {
			err = etcdStorage.Create(user1)
			Expect(err).To(BeNil())
			deleteMark = true
		})
		It("Insert user error", func() {
			err = etcdStorage.Create(user1)
			err = etcdStorage.Create(user1)
			Expect(err).NotTo(BeNil())
			deleteMark = true
		})
	})
	Context("Delete user tests mock", func() {
		It("Delete user", func() {
			err = storage.Create(user1)
			Expect(err).To(BeNil())
			err = storage.Delete(user1)
			Expect(err).To(BeNil())
		})
		It("Delete user with no existent user", func() {
			err = storage.Delete(user1)
			log.Debugf("%v", err)
			Expect(err).To(BeNil())
		})
	})
	Context("Delete user tests", func() {
		It("Delete user", func() {
			err = etcdStorage.Create(user1)
			Expect(err).To(BeNil())
			err = etcdStorage.Delete(user1)
			Expect(err).To(BeNil())
		})
		It("Delete with no existent user", func() {
			user1.Id = 2
			Expect(err).To(BeNil())
			err = storage.Create(user1)
			user1.Id = 1
			err = storage.Delete(user1)
			Expect(err).To(BeNil())
		})
	})
})
