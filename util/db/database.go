package db

import (
	"database/sql"
	"fmt"
	"golang-blog-journey/util/log"
)

type DBT struct {
	DB *sql.DB
}

var dbT *DBT

func InitDatabase(user, password, ip string) {
	dbT = new(DBT)
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/", user, password, ip)
	db, err := sql.Open("mysql", dataSourceName)
	log.Infof("dataSource", dataSourceName)
	if err != nil {
		log.Errorw("InitDatabase", err)
		db.Close()
	} else {
		dbT.DB = db
	}
}

func ExecSQL(str string) (int64, error) {
	res, err := dbT.DB.Exec(str)
	if err != nil {
		log.Errorw("ExecSQL", err)
		return 0, err
	}
	return res.RowsAffected()
}

func QuerySQL(str string) (*sql.Rows, error) {
	res, err := dbT.DB.Query(str)
	if err != nil {
		log.Errorw("QuerySQL", err)
		return nil, err
	}
	return res, nil
}
