package datastore

import (
	"fmt"
	"sync"
	"voting-app/voting-app-worker/utils/logger"

	"voting-app/voting-app-worker/types/datastore"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var dbLogger *logrus.Entry = logger.GetLogger("db")

type PgDB struct {
	*sqlx.DB
	mutex *sync.RWMutex
}
type Votes struct {
	VoterID int    `db:"voter_id" json:"voter_id"`
	Vote    string `db:"vote" json:"vote"`
}

var PgDBInstance *PgDB

func NewPgDB() *PgDB {
	var dbname = "vote"
	// var dbname = "pgdata"
	var user = "testuser1"
	// var password = "password123!"
	var password = "password123"
	var host = "localhost"
	var port = "5432"
	var postgresqlConnectionString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	sqlxdb, err := sqlx.Connect("postgres", postgresqlConnectionString)
	if err != nil {
		dbLogger.Fatalln("Failed to connect to zakkaya database:", err)
	}

	if err := sqlxdb.Ping(); err != nil {
		dbLogger.Fatal(err)
		return nil
	}

	pgDB := PgDB{DB: sqlxdb, mutex: &sync.RWMutex{}}
	return &pgDB
}

func (db *PgDB) GetVoteResults() []datastore.VoteResult {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	voteResults := []datastore.VoteResult{}
	err2 := db.Select(&voteResults, "SELECT vote, COUNT(id) AS count FROM votes GROUP BY vote")
	if err2 != nil {
		dbLogger.Errorf("%#v", err2)
		return nil
	}
	dbLogger.Infof("%+v", voteResults)
	return voteResults
}

func (db *PgDB) GetAllVotes(field string) string {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	votes := []Votes{}
	err := db.Select(&votes, "SELECT voter_id, vote FROM votes")
	if err != nil {
		dbLogger.Errorf("%#v", err)
		return ""
	}
	dbLogger.Infof("%#v", votes)

	return "value"
}
