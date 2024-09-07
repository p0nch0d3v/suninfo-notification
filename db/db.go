package db

import (
	"database/sql"
	"fmt"
	"log"
	"suninfo-notification/models"

	_ "github.com/mattn/go-sqlite3"
)

const fileName string = "data/log.db"

var db *sql.DB

func getConnection() *sql.DB {
	if db != nil {
		return db
	}

	var err error
	db, err = sql.Open("sqlite3", fileName)

	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	return db
}

func Init() {
	log.Println("DB - Init")
	db = getConnection()
	var err error

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Log (Id NVARCHAR(10) PRIMARY KEY, Sunset NVARCHAR(11), TwilightEnd NVARCHAR(11), Message NVARCHAR(50));")

	if err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}
}

func IsDateAlreadyAdded(date string) bool {
	row := db.QueryRow(`
    SELECT Id, Sunset, TwilightEnd, Message 
    FROM Log 
    WHERE Id=?`, date)

	if row != nil && row.Err() == nil {
		i := models.LogItem{}
		err := row.Scan(&i.Id, &i.Sunset, &i.TwilightEnd, &i.Message)
		return err == nil || len(i.Id) > 0
	}
	return false
}

func AddSunInfo(date string, sunset string, twilightEnd string, message string) bool {
	result, err := db.Exec(fmt.Sprintf(`INSERT INTO log (Id, Sunset, TwilightEnd, Message) VALUES ('%s', '%s', '%s', '%s');`, date, sunset, twilightEnd, message))
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
	rowsAffected, _ := result.RowsAffected()
	lastInsertId, _ := result.LastInsertId()
	return rowsAffected > 0 && lastInsertId > 0
}

func GetAllLog() []models.LogItem {
	rows, err := db.Query("SELECT * FROM log;")

	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	defer rows.Close()
	data := []models.LogItem{}
	for rows.Next() {
		i := models.LogItem{}
		err = rows.Scan(&i.Id, &i.Sunset, &i.TwilightEnd, &i.Message)
		if err != nil {
			log.Fatalln(err)
			panic(err)
		}
		data = append(data, i)
	}

	return data
}
