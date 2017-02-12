package mongo

import (
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/course-extended-golang/users/storages"
	"github.com/course-extended-golang/users"
)

type MongoDB struct {
	mongoSession *mgo.Session
}

func New() storages.Storage {
	mongodb := new(MongoDB)
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
	if err = session.Ping(); err != nil {
		log.Fatal(err)
	}
	mongodb.mongoSession = session
	return mongodb
}

func (m *MongoDB) Create(entity users.User) error {
	c := m.mongoSession.DB("curso").C("curso")
	count, err := c.Find(bson.M{"id": entity.Id}).Count()
	if err != nil && err.Error() != "not found" {
		log.Error(err)
		return err
	} else if count > 0 {
		log.Errorf("User '%+v' already exist", entity)
		return errors.New(fmt.Sprintf("User '%+v' already exist", entity))
	}
	err = c.Insert(entity)
	if err != nil {
		log.Error(err)
	}
	return err
}
func (m *MongoDB) Delete(entity users.User) error {
	c := m.mongoSession.DB("curso").C("curso")
	if err := c.Remove(bson.M{"id": entity.Id}); err != nil {
		log.Error(err)
		return err
	}
	return nil
}
func (m *MongoDB) Close() error {
	m.mongoSession.Close()
	return nil
}
