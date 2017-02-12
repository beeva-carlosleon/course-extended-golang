package main

import (
	log "github.com/Sirupsen/logrus"
	userroutes "github.com/course-extended-golang/users/routes"
	"github.com/course-extended-golang/users/storages/mysql"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-swagger12"
	"net/http"
	"time"
)

func startServer() {
	connectionPool := mysql.New()
	defer connectionPool.Close()
	wsContainer := restful.NewContainer()
	wsContainer.Add(userroutes.NewEndPoint(connectionPool))
	// Optionally, you can install the Swagger Service which provides a nice Web UI on your REST API
	// You need to download the Swagger HTML5 assets and change the FilePath location in the config below.
	// Open http://localhost:8080/apidocs and enter http://localhost:8080/apidocs.json in the api input field.
	config := swagger.Config{
		WebServices:    wsContainer.RegisteredWebServices(), // you control what services are visible
		WebServicesUrl: "http://localhost:8080",
		ApiPath:        "/apidocs.json",

		// Optionally, specify where the UI is located
		SwaggerPath:     "/apidocs/",
		SwaggerFilePath: "swagger",
		ApiVersion:      "1.0.0",
		Info: swagger.Info{
			Title:       "Course Golang API",
			Contact:     "contact_mail@example.com",
			Description: "API for learning rest",
		},
	}
	swagger.RegisterSwaggerService(config, wsContainer)

	log.Printf("start listening on localhost:8080")
	server := &http.Server{Addr: ":8080", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}

func main() {
	startServer()
}

func init() {
	log.SetLevel(log.ErrorLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, TimestampFormat: time.RFC3339})
}
