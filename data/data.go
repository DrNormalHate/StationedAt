package data

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func CreateDBEngine() (*sql.DB, error) {
	db, err := sql.Open("mysql", "developer:SGTams1212$$@tcp(146.148.41.233:3306)/StationedAt")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}

func PreparedStatement(statement string, args ...interface{}) error {
	createStatement := `
		DESCRIBE Users
	`
	if data, err := CreateDBEngine(); err == nil {
		defer data.Close()
		if tx, err := data.Begin(); err == nil {
			defer tx.Rollback()
			if stmt, err := tx.Prepare(createStatement); err == nil {
				defer stmt.Close()
				if _, err := stmt.Exec(); err == nil {
					if err := tx.Commit(); err == nil {

						return nil
					} else {
						return err
					}
				} else {
					return err
				}
			} else {
				return err
			}
		} else {
			return err
		}
	} else {
		return err
	}
}
