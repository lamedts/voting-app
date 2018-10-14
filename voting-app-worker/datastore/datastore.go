package datastore

import (
	"fmt"
	"sync"

	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PgDB struct {
	*sqlx.DB
	mutex *sync.RWMutex
}

var PgDBInstance *PgDB

func NewPgDB() *PgDB {
	var dbname = "pgdata"
	var user = "testuser1"
	var password = "password123!"
	var host = "localhost"
	var port = "5432"
	var postgresqlConnectionString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	sqlxdb, err := sqlx.Connect("postgres", postgresqlConnectionString)
	if err != nil {
		glog.Fatalln("Failed to connect to zakkaya database:", err)
	}

	if err := sqlxdb.Ping(); err != nil {
		glog.Fatal(err)
		return nil
	}

	pgDB := PgDB{DB: sqlxdb, mutex: &sync.RWMutex{}}
	return &pgDB
}

func (db *PgDB) GetVoteResult(field string) string {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	var value string
	err := db.Get(&value, "SELECT vote, COUNT(id) AS count FROM votes GROUP BY vote", field)
	if err != nil {
		return ""
	}
	return value
}

func (db *PgDB) Vote(field string) string {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	var value string
	err := db.Get(&value, "SELECT value from system_metadata WHERE field=$1", field)
	if err != nil {
		return ""
	}
	return value
}
