package mysql

import (
	"database/sql"
	log "github.com/Sirupsen/logrus"
	"github.com/course-extended-golang/users"
	"github.com/course-extended-golang/users/storages"
	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	mysqlConnection *sql.DB
}

func New() storages.Storage {
	mysql := new(MySQL)
	db, err := sql.Open("mysql", "root:curso@tcp(localhost:3306)/curso")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	mysql.mysqlConnection = db
	return mysql
}

func (m *MySQL) Create(entity users.User) error {
	tx, err := m.mysqlConnection.Begin()
	if err != nil {
		log.Error(err)
	}

	stmt, err := tx.Prepare("INSERT INTO users (ID,NAME,SURNAME,AGE) VALUES (?,?,?,?)")
	if err != nil {
		log.Error(err)
	}
	defer stmt.Close()
	res, err := stmt.Exec(entity.Id, entity.Name, entity.SurName, entity.Age)
	if err != nil {
		log.Error(err)
		defer tx.Rollback()
	} else {
		lastId, err := res.LastInsertId()
		if err != nil {
			log.Error(err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Error(err)
		}
		log.Debugf("ID = %d, affected = %d", lastId, rowCnt)
		err = tx.Commit()
	}
	return err
}
func (m *MySQL) Delete(entity users.User) error {
	tx, err := m.mysqlConnection.Begin()
	if err != nil {
		log.Error(err)
	}

	stmt, err := tx.Prepare("DELETE FROM users WHERE ID=?")
	if err != nil {
		log.Error(err)
	}
	defer stmt.Close()
	res, err := stmt.Exec(entity.Id)
	if err != nil {
		log.Error(err)
		defer tx.Rollback()
	} else {
		lastId, err := res.LastInsertId()
		if err != nil {
			log.Error(err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Error(err)
		}
		log.Debugf("ID = %d, affected = %d", lastId, rowCnt)
		err = tx.Commit()
	}
	return err
}
func (m *MySQL) Close() error {
	return m.mysqlConnection.Close()
}
