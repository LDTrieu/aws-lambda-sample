package mysqlcontroller

import (
	"context"
	"database/sql"
	"fmt"
	"lambda-sample/pkg/sercfg"
	"lambda-sample/pkg/wUtil"
	"log"
	"time"
)

var (
	dbIns, readOnlyDbIns *sql.DB
)

func initDB() (db *sql.DB, err error) {

	ctx := context.Background()
	dbURL := sercfg.Get(ctx, "mysql_db")
	userName := sercfg.Get(ctx, "mysql_user")
	pwd := sercfg.Get(ctx, "mysql_pwd")
	schema := sercfg.Get(ctx, "mysql_db_schema")

	datasource := fmt.Sprintf("%v:%v@tcp(%v)/%v?parseTime=true", userName, pwd, dbURL, schema)
	db, err = sql.Open("mysql", datasource)
	if err != nil {
		err = wUtil.NewError(err)
		return
	}
	// defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	if err = db.Ping(); err != nil {
		err = wUtil.NewError(err)
	}

	return
}

func initReadOnlyDb() (db *sql.DB, err error) {
	ctx := context.Background()
	dbURL := sercfg.Get(ctx, "read_only_mysql_db")
	userName := sercfg.Get(ctx, "mysql_user")
	pwd := sercfg.Get(ctx, "mysql_pwd")
	schema := sercfg.Get(ctx, "mysql_db_schema")

	datasource := fmt.Sprintf("%v:%v@tcp(%v)/%v?parseTime=true", userName, pwd, dbURL, schema)

	log.Println(wUtil.StrLog(dbURL))
	db, err = sql.Open("mysql", datasource)
	if err != nil {
		err = wUtil.NewError(err)
		return
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	if err = db.Ping(); err != nil {
		err = wUtil.NewError(err)
	}
	log.Println(wUtil.StrLog("ping read only with err:", err))
	return
}

func runReadQuery(act func(db *sql.DB) error) (err error) {
	if readOnlyDbIns == nil {
		readOnlyDbIns, err = initReadOnlyDb()
		if err != nil {
			return
		}
	}
	return act(readOnlyDbIns)
}

func runQuery(act func(*sql.DB) error) (err error) {
	if dbIns == nil {
		dbIns, err = initDB()
		if err != nil {
			return
		}
	}
	err = act(dbIns)
	return
}

func RunReadOnly(act func(db *sql.DB) error) error {
	return runReadQuery(act)
}

func RunQuery(act func(db *sql.DB) error) error {
	return runQuery(act)
}
