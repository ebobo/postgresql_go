package store

import (
	_ "embed" // support embedding files in variables
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

//go:embed schema.sql
var schema string

func RunSqlx() {
	fmt.Println("run sql package")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		db.Close()
		panic(err)
	}

	fmt.Println("Successfully connected!")

	dropTable(db)

	err = createSchema(db)
	if err != nil {
		db.Close()
		panic(err)
	}

	addLegoSet(db, 10278, "Police Station", "Creator")
	addLegoSet(db, 10220, "Volkswagen T1 Camper Van", "Creator")
	addLegoSet(db, 10252, "Folkevognboble", "Creator")
	addLegoSet(db, 21309, "NASA Apollo Saturn V", "Ideas")

	displayLegoTableContent(db)

	addLegoSet(db, 21321, "International Space Station", "Ideas")

	log.Println("Database Close")

}

func createSchema(db *sqlx.DB) error {
	for n, statement := range strings.Split(schema, ";") {
		_, err := db.Exec(statement)
		if err != nil {
			return fmt.Errorf("statement %d failed: \"%s\" : %w", n+1, statement, err)
		}
	}
	return nil
}

func dropTable(db *sqlx.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS lego")
	if err == nil {
		fmt.Println("Lego table dropped")
	}
	return err
}

// We are passing db reference connection from main to our method with other parameters
func addLegoSet(db *sqlx.DB, id int, catalog string, name string) {
	fmt.Println("Inserting lego info to table ...")
	// use "ON CONFLICT DO NOTHING" to avoid insert duplication
	insertLegoSQL := `INSERT INTO lego (model_id, catalog, name) VALUES ($1, $2, $3) ON CONFLICT (model_id) DO NOTHING`
	statement, err := db.Prepare(insertLegoSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = statement.Exec(id, catalog, name)
	if err != nil {
		fmt.Println("err 2")
		fmt.Println(err.Error())
	}
}

func displayLegoTableContent(db *sqlx.DB) {
	fmt.Println("display lego table")
	row, err := db.Query("SELECT * FROM lego ORDER BY catalog")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int64
		var model_id int
		var catalog string
		var name string
		row.Scan(&id, &model_id, &catalog, &name)
		log.Println("Set: ", id, " ", model_id, " ", catalog, " ", name)
	}
}
