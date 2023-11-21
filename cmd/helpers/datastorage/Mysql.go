package datastorage

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

func GetConnection(l *zerolog.Logger) *sqlx.DB {
	mysqlConnStr := os.Getenv("DB_URL")
	// connStr = "mysql://bw:rootroot@tcp(192.168.28.81:3306)/wc_atlas_bkp_2023_11_07"
	// connStr = "go:rootroot@(localhost:8101)/billsplit"
	connStrArr := strings.SplitN(mysqlConnStr, "mysql://", 1)
	connStr := strings.TrimPrefix(mysqlConnStr, "mysql://")
	fmt.Println(connStrArr)
	fmt.Println(connStr)
	db, err := sqlx.Connect("mysql", connStr)
	if err != nil {
		l.Panic().Stack().Msgf("Cannot Connect Database: %s", err.Error())
		panic(err)
	}
	re := regexp.MustCompile(`[()]`)

	// Use the regular expression to split the string
	parts := re.Split(connStr, -1)

	// fmt.Println(parts)
	l.Info().Msgf("Database Connected!: %s", parts[1])

	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
