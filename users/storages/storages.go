package storages

import "github.com/course-extended-golang/users"

type Storage interface {
	Create(entity users.User) error
	Delete(entity users.User) error
	Close() error
}
