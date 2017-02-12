package mongo

import (
	"fmt"
	log "github.com/Sirupsen/logrus"

	"github.com/course-extended-golang/users"
	"github.com/course-extended-golang/users/storages"
)

type MongoMock struct {
	database map[string][]users.User
}

func NewMock() storages.Storage {
	mock := new(MongoMock)
	mock.database = make(map[string][]users.User)
	mock.database["users"] = make([]users.User, 0)
	return mock
}

func (m *MongoMock) Create(entity users.User) error {
	exist := false
	for _, v := range m.database["users"] {
		if v.Id == entity.Id {
			exist = true
		}
	}
	if exist {
		log.Errorf("User %+v already exist", entity)
		return fmt.Errorf("User %+v already exist", entity)
	}
	m.database["users"] = append(m.database["users"], entity)
	return nil
}

func (m *MongoMock) Delete(entity users.User) error {
	deleted := false
	newArray := make([]users.User, 0)
	for _, v := range m.database["users"] {
		if v.Id == entity.Id {
			deleted = true
		} else {
			newArray = append(newArray, v)
		}
	}
	if deleted {
		m.database["users"] = newArray
		log.Debugf("User %+v deleted", entity)
	}
	return nil
}
func (m *MongoMock) Close() error {
	return nil
}
