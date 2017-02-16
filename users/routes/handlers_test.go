package routes_test

import (
	"bytes"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/course-extended-golang/users"
	"github.com/course-extended-golang/users/routes"
	"github.com/emicklei/go-restful"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"strings"
)

type test struct {
	NovalidEntity string
}

var _ = Describe("User checks", func() {
	var (
		user            users.User
		userBody        []byte
		errorEntity     test = test{}
		errorEntityBody []byte
		err             error
	)
	BeforeEach(func() {
		user = users.User{Id: 1, Name: "Carlos", SurName: "León", Age: 38}
		userBody, err = json.Marshal(user)
		Expect(err).To(BeNil())
		errorEntityBody, err = json.Marshal(errorEntity)
		Expect(err).To(BeNil())
	})
	AfterEach(func() {
		mockStorage.Delete(user)
	})
	Context("Create user handler tests", func() {
		It("Create new default handler", func() {
			Expect(routes.New()).NotTo(BeNil())
		})
		It("Create user call", func() {
			handler := routes.NewForStorage(mockStorage)
			log.Debugf("JSON to send: '%s'", userBody)
			req := httptest.NewRequest("POST", "/users", bytes.NewReader(userBody))
			req.Header.Set("Content-Type", restful.MIME_JSON)
			w := httptest.NewRecorder()
			handler.Create(restful.NewRequest(req), restful.NewResponse(w))
			Expect(w.Code).To(Equal(http.StatusCreated))
			Expect(strings.TrimSpace(w.Body.String())).To(Equal(string(userBody)))
		})
		It("Create user error call", func() {
			handler := routes.NewForStorage(mockStorage)
			log.Debugf("JSON to send: '%s'", errorEntityBody)
			req := httptest.NewRequest("POST", "/users", bytes.NewReader(errorEntityBody))
			req.Header.Set("Content-Type", restful.MIME_JSON)
			w := httptest.NewRecorder()
			handler.Create(restful.NewRequest(req), restful.NewResponse(w))
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})
		It("Error creating an user twice", func() {
			handler := routes.NewForStorage(mockStorage)
			log.Debugf("JSON to send: '%s'", userBody)
			req := httptest.NewRequest("POST", "/users", bytes.NewReader(userBody))
			req.Header.Set("Content-Type", restful.MIME_JSON)
			w := httptest.NewRecorder()
			handler.Create(restful.NewRequest(req), restful.NewResponse(w))
			Expect(w.Code).To(Equal(http.StatusCreated))
			req = httptest.NewRequest("POST", "/users", bytes.NewReader(userBody))
			req.Header.Set("Content-Type", restful.MIME_JSON)
			w = httptest.NewRecorder()
			handler.Create(restful.NewRequest(req), restful.NewResponse(w))
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
	})
})
