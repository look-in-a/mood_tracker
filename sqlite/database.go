package sqlite

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

var database *sql.DB

func init() {
	database, _ = sql.Open("sqlite3", "./etc/mood.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS mood_history (time INTEGER, mood TEXT, depth INTEGER)")
	statement.Exec()
}

func InsertMoods(moods map[string]int) error {
	now := time.Now().Unix()
	statement, err := database.Prepare("INSERT INTO mood_history (time, mood, depth) VALUES (?, ?, ?)")
	if err != nil {
		return errors.Wrap(err, "failed to prepare insert request")
	}
	for mood, depth := range moods { // TODO add transaction?
		_, err := statement.Exec(now, mood, depth)
		if err != nil {
			return errors.Wrap(err, "failed to exec insert request")
		}
	}
	return nil
}

func PrintMoods() {
	rows, _ := database.Query("SELECT time, mood, depth FROM mood_history")
	var time, depth int
	var mood string
	for rows.Next() {
		rows.Scan(&time, &mood, &depth)
		fmt.Printf("%v: %v(%v)\n", time, mood, depth)
	}
}
