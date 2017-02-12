package users_test

import (
	log "github.com/Sirupsen/logrus"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/course-extended-golang/users"
)

var _ = Describe("User checks", func() {
	var (
		user1 users.User
		user2 users.User
	)
	BeforeEach(func() {
		log.Debugf("This is runned before IT")
		user1 = users.User{}
		user2 = users.User{Id: 1, Name: "Carlos", SurName: "León", Age: 38}
	})
	Context("Empty fields", func() {
		It("All fields are empty", func() {
			Expect(user1).To(Equal(users.User{}))
			log.Debugf("IT - All fields are empty")
		})
		It("All fields are fulfilled", func() {
			Expect(user2.Id).NotTo(Equal(int64(0)))
			Expect(user2.Name).NotTo(Equal(""))
			Expect(user2.SurName).NotTo(Equal(""))
			Expect(user2.Age).NotTo(Equal(byte(0)))
			log.Debugf("IT - All fields are fulfilled")
		})
	})
	Context("Fields validation", func() {
		It("All fields are correct", func() {
			Expect(user2.Id).To(Equal(int64(1)))
			Expect(user2.Name).To(Equal("Carlos"))
			Expect(user2.SurName).To(Equal("León"))
			Expect(user2.Age).To(Equal(byte(38)))
			log.Debugf("IT - All fields are correct")
		})
	})
	Context("Swagger config", func() {
		It("All fields", func() {
			Expect(user2.SwaggerDoc()).NotTo(BeNil())

		})
	})
})
