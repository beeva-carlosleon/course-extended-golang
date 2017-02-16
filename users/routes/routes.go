package routes

import (
	"github.com/course-extended-golang/users"
	"github.com/course-extended-golang/users/storages"
	"github.com/emicklei/go-restful"
)

const (
	EndPoint      = "/users"
	PathParameter = "/{user-id}"
)

func NewEndPoint(storage storages.Storage) *restful.WebService {
	handlers := NewForStorage(storage)
	service := new(restful.WebService)
	service.
		Path(EndPoint).
		Filter(NewFilter().LogFilter).
		Doc("Manage Users entities").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	service.Route(service.POST("").
		Doc("create a user").
		Operation("Create").
		To(handlers.Create).
		Reads(users.User{}).
		Writes(users.User{}))
	return service
}
