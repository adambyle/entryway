package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

type pronounSet struct {
	subject, object, possAdj, possNoun string
}

var pronouns = []pronounSet{
	{"they", "them", "their", "theirs"},
	{"he", "him", "his", "his"},
	{"she", "her", "her", "hers"},
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")
	r.Static("/static", "./static")

	const entrywayDb = "entryway"
	db, dbOpenErr := sql.Open("sqlite", fmt.Sprintf("./db/%s.db", entrywayDb))
	if dbOpenErr != nil {
		log.Fatal(dbOpenErr)
	}
	defer db.Close()

	if dbPingErr := db.Ping(); dbPingErr != nil {
		log.Fatal(dbPingErr)
	}

	if dbInitErr := initializeDb(db); dbInitErr != nil {
		log.Fatal(dbInitErr)
	}

	r.GET("/login", getLoginPage)

	r.Run(":8080")
}

func initializeDb(db *sql.DB) error {
	_, err := db.Exec(`--sql
		CREATE TABLE IF NOT EXISTS users (
			UserID     INTEGER PRIMARY KEY AUTOINCREMENT,
			Username   TEXT    UNIQUE NOT NULL,
			Password   TEXT    NOT NULL,
			Email      TEXT,
			FirstName  TEXT,
			LastName   TEXT,
			Pronouns   INTEGER DEFAULT(0)
		);
	`)
	return err
}

func getLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "universal.html", gin.H{
		"title":   "Login",
		"content": "loginContent",
	})
}
