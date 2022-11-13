package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Five-Series/questions/infra/environment"
)

const ZeroDBUnix = int64(-62135596800)

type DataBase struct {
	Env *environment.DB
}

// New Constructor for DataBase struct
func New(env *environment.DB) *DataBase {
	return &DataBase{Env: env}

}

func (d *DataBase) buildConnString() string {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", d.Env.DBhost, d.Env.DBPort, d.Env.DbUser, d.Env.DBPass)
	return connStr

}

// Connect Responsible for connect to Postgree
func (d *DataBase) Connect() *sql.DB {

	db, err := sql.Open("postgres", d.buildConnString())
	if err != nil {
		msg := fmt.Sprintf("Error Connecting to database "+time.Now().String(), err)
		log.Fatal(msg)
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db
}
