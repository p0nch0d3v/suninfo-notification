package db

import (
	"database/sql"
	"fmt"
	"log"
	"sunrise-sunset-notification/models"

	_ "github.com/mattn/go-sqlite3"
)

const fileName string = "log.db"

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
	db = getConnection()
	var err error

	log.Println("CREATE")

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Log (Id NVARCHAR(10) PRIMARY KEY, Sunset NVARCHAR(11), TwilightEnd NVARCHAR(11));")

	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
}

func IsDateAlreadyAdded(date string) bool {
	row := db.QueryRow(`
    SELECT Id, Sunset, TwilightEnd 
    FROM Log 
    WHERE Id=?`, date)

	if row != nil && row.Err() == nil {
		i := models.LogItem{}
		err := row.Scan(&i.Id, &i.Sunset, &i.TwilightEnd)
		return err == nil || len(i.Id) > 0
	}
	return false
}

func AddSunInfo(date string, sunset string, twilightEnd string) bool {
	result, err := db.Exec(fmt.Sprintf(`INSERT INTO log (Id, Sunset, TwilightEnd) VALUES ('%s', '%s', '%s');`, date, sunset, twilightEnd))
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
		err = rows.Scan(&i.Id, &i.Sunset, &i.TwilightEnd)
		if err != nil {
			log.Fatalln(err)
			panic(err)
		}
		data = append(data, i)
	}

	return data
}
