package routes

import (
	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/emicklei/go-restful"
	"net/http"
	"github.com/course-extended-golang/users/storages"
	"github.com/course-extended-golang/users"
)

type Handler interface {
	Create(request *restful.Request, response *restful.Response)
}

type DefaultHandler struct {
	storage storages.Storage
}

func New() Handler {
	return new(DefaultHandler)
}

func NewForStorage(storage storages.Storage) Handler {
	handler := new(DefaultHandler)
	handler.storage = storage
	return handler
}

func (h DefaultHandler) Create(request *restful.Request, response *restful.Response) {
	user := new(users.User)
	log.Debugf("Request: %+v", request)
	err := request.ReadEntity(user)
	if err == nil {
		if (users.User{}) == *user {
			log.Errorf("Invalid Entity")
			response.WriteError(http.StatusBadRequest, errors.New("Invalid Entity"))
		} else {
			log.Debugf("handler-User: %+v", user)
			if err := h.storage.Create(*user); err != nil {
				log.Errorf("Error: %+v", err)
				response.WriteError(http.StatusInternalServerError, err)
			} else {
				log.Debugf("Handler-Created User: %+v", user)
				response.PrettyPrint(false)
				response.WriteHeaderAndJson(http.StatusCreated, user, restful.MIME_JSON)
				//response.WriteEntity()
			}
		}
	} else {
		log.Errorf("Error: %+v", err)
		response.WriteError(http.StatusInternalServerError, err)
	}
}
