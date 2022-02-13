package main

import (
	"database/sql"
	"fmt"
	"log"

	// load postgres driver
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "lego_db"
)

func main() {
	fmt.Println("Hello Qi")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	createLegoTable(db)

	addLegoSet(db, 10278, "Police Station", "Creator")
	addLegoSet(db, 10220, "Volkswagen T1 Camper Van", "Creator")

	displayLegoTableContent(db)

	log.Println("Database Close")
}

func createLegoTable(db *sql.DB) {
	createLegoTableSQL := `CREATE TABLE IF NOT EXISTS lego (
		"model_id" integer NOT NULL PRIMARY KEY,		
		"catalog" TEXT NOT NULL,
		"name" TEXT NOT NULL	
	  );` // SQL Statement for Create Table

	statement, err := db.Prepare(createLegoTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	fmt.Println("lego table created")
}

// We are passing db reference connection from main to our method with other parameters
func addLegoSet(db *sql.DB, id int, catalog string, name string) {
	log.Println("Inserting lego info to table ...")
	insertLegoSQL := `INSERT INTO lego (model_id, catalog, name) VALUES ($1, $2, $3)`
	statement, err := db.Prepare(insertLegoSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(id, catalog, name)
	if err != nil {
		fmt.Println("err 2")
		log.Fatalln(err.Error())
	}
}

func displayLegoTableContent(db *sql.DB) {
	fmt.Println("display lego table")
	row, err := db.Query("SELECT * FROM lego ORDER BY catalog")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var catalog string
		var name string
		row.Scan(&id, &catalog, &name)
		log.Println("Set: ", id, " ", catalog, " ", name)
	}
}
