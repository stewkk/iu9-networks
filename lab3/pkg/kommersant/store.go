package kommersant

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DbUsername string
	DbPassword string
	DbAddress  string
	DbName     string
)

func init() {
	DbUsername = os.Getenv("DB_USER")
	DbPassword = os.Getenv("DB_PASS")
	DbAddress = os.Getenv("DB_ADDR")
	DbName = os.Getenv("DB_NAME")
}

func Store(news []NewsEntry) error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", DbUsername, DbPassword, DbAddress, DbName))
	if err != nil {
		return err
	}

	db.Exec(`DROP TABLE ihatemysql;`)

	_, err = db.Exec(`
CREATE TABLE IF NOT EXISTS ihatemysql(
id varchar(300) PRIMARY KEY,
link varchar(300),
category varchar(50) CHARACTER SET utf8mb4,
title varchar(150) CHARACTER SET utf8mb4,
date timestamp,
description text CHARACTER SET utf8mb4
);`)
	if err != nil {
		return err
	}

	insStmt, err := db.Prepare(`INSERT INTO ihatemysql VALUES(?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE id=id`)
	if err != nil {
		return err
	}
	defer insStmt.Close()

	for _, entry := range news {
		_, err := insStmt.Exec(entry.Id, entry.Link, entry.Category, entry.Title, entry.PubDate, entry.Description)
		if err != nil {
			return err
		}
	}

	return nil
}

func Clear() error {
	_, err := db.Exec(`TRUNCATE TABLE ihatemysql`)
	return err
}
