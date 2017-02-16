package routes

import (
	log "github.com/Sirupsen/logrus"
	"github.com/emicklei/go-restful"
)

type Filter interface {
	LogFilter(request *restful.Request, response *restful.Response, chain *restful.FilterChain)
}

type LogFilter struct {
}

func NewFilter() Filter {
	return new(LogFilter)
}

func (h LogFilter) LogFilter(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
	log.Printf("[webservice-filter (logger)] %s,%s\n", request.Request.Method, request.Request.URL)
	chain.ProcessFilter(request, response)

}
