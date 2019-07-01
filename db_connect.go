package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/josuegiron/log"
)

var db *sql.DB

func dbConnect() bool {

	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", config.DataBase.User, config.DataBase.Password, config.DataBase.Server, config.DataBase.Port, config.DataBase.DataBase)

	var err error

	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Error(err)
		return false
	}

	if err = db.Ping(); err != nil {
		log.Error(err)
		return false
	}

	ctx := context.Background()

	// Ping database to see if it's still alive.
	// Important for handling network issues and long queries.
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	err = mysqlVersion()
	if err != nil {
		log.Error(err)
		return false
	}

	return true

}

func mysqlVersion() error {
	query := "select version()"

	row := db.QueryRow(query)

	var version string

	err := row.Scan(&version)
	if err != nil {
		return err
	}

	log.Info(version)

	return nil
}

// Generate string sql query
func getQueryString(query string, params ...interface{}) (string, error) {
	valid := regexp.MustCompile(` @(([a-z]|[A-Z]|[0-9])+)( |$)`)

	for _, param := range params {
		sqlparam := param.(sql.NamedArg)
		value := getValue(sqlparam.Value)
		prm := regexp.MustCompile(` @` + sqlparam.Name + `( |$)`)
		for prm.MatchString(query) {
			query = prm.ReplaceAllLiteralString(query, fmt.Sprintf(" %s ", value))
		}
	}

	if valid.MatchString(query) {
		return query, errors.New("Existen parametros vacíos")
	}

	return query, nil
}

func getValue(v interface{}) string {
	switch v.(type) {
	case int:
		return fmt.Sprintf("%v", v.(int))
	case int64:
		return fmt.Sprintf("%v", v.(int64))
	case float64:
		return fmt.Sprintf("%v", v.(float64))
	case string:
		return fmt.Sprintf("'%v'", v.(string))
	case bool:
		return fmt.Sprintf("%v", v.(bool))
	case time.Time:
		return fmt.Sprintf("'%v'", v.(time.Time).Format("2006-01-02 15:04:05-06:00")) // for Guatemala UTC
	//... etc
	default:
		return "unknown"
	}
}
