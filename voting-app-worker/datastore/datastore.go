package datastore

import (
	"fmt"
	"sync"
	"voting-app/voting-app-worker/config"
	types "voting-app/voting-app-worker/types/datastore"
	"voting-app/voting-app-worker/utils/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var dbLogger = logger.GetLogger("db")

type PgDB struct {
	*sqlx.DB
	mutex *sync.RWMutex
}

var PgDBInstance *PgDB

func NewPgDB(pgConfig config.PgConfig) *PgDB {
	var postgresqlConnectionString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", pgConfig.User, pgConfig.Password, pgConfig.Host, pgConfig.Port, pgConfig.DBName)
	sqlxdb, err := sqlx.Connect("postgres", postgresqlConnectionString)
	if err != nil {
		dbLogger.Fatalln("Failed to connect to database:", err)
	}

	if err := sqlxdb.Ping(); err != nil {
		dbLogger.Fatal(err)
		return nil
	}

	pgDB := PgDB{DB: sqlxdb, mutex: &sync.RWMutex{}}
	return &pgDB
}

func (db *PgDB) GetVoteResults() []types.VoteResult {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	voteResults := []types.VoteResult{}
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

	votes := []types.Vote{}
	err := db.Select(&votes, "SELECT voter_id, vote FROM votes")
	if err != nil {
		dbLogger.Errorf("%#v", err)
		return ""
	}
	dbLogger.Infof("%#v", votes)

	return "value"
}

func (db *PgDB) InsertVote(vote types.Vote) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if _, err := db.NamedExec(`INSERT INTO votes (voter_id, vote) VALUES (:voter_id, :vote)`, vote); err != nil {
		dbLogger.WithFields(logrus.Fields{
			"Flow": "datastore",
			"func": "insert vote",
		}).Warn(err)
		return err
	}

	return nil
}
