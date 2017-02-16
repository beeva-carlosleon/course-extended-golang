package etcd

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/coreos/etcd/clientv3"
	"github.com/course-extended-golang/users"
	"github.com/course-extended-golang/users/storages"
	"golang.org/x/net/context"
	"time"
)

type EtcdDB struct {
	config clientv3.Config
}

func New(endpoint string) storages.Storage {
	etcd := new(EtcdDB)
	etcd.config = clientv3.Config{
		Endpoints:   []string{endpoint},
		DialTimeout: 5 * time.Second,
	}
	return etcd
}

func (e *EtcdDB) Create(entity users.User) error {
	client, err := clientv3.New(e.config)
	if err != nil {
		return err
	}
	defer client.Close()
	result, err := client.Get(context.TODO(), fmt.Sprintf("/users/%d", entity.Id))
	if result.Count > 0 {
		log.Errorf("User '%+v' already exist", entity)
		return errors.New(fmt.Sprintf("User '%+v' already exist", entity))
	}
	userBytes, err := json.Marshal(entity)
	if err != nil {
		return err
	}
	_, err = client.Put(context.TODO(), fmt.Sprintf("/users/%d", entity.Id), string(userBytes))
	if err != nil {
		return err
	}
	return err
}
func (e *EtcdDB) Delete(entity users.User) error {
	client, err := clientv3.New(e.config)
	if err != nil {
		return err
	}
	defer client.Close()
	result, err := client.Delete(context.TODO(), fmt.Sprintf("/users/%d", entity.Id))
	if err != nil {
		return err
	}
	log.Debugf("Inserted %d", result.Deleted)
	return nil
}
func (e *EtcdDB) Close() error {
	return nil
}
