package migration

import (
	"database/sql"
	"fmt"
	"go-boilerplate/shared/database"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func Migrate() {
	connString := database.GetConnectionString()
	db, err := sql.Open("mysql", fmt.Sprintf("%s&multiStatements=true", connString))
	if err != nil {
		fmt.Errorf("Error: %s", err.Error())
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		fmt.Errorf("Error: %s", err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./database/migration",
		"mysql",
		driver,
	)

	if err != nil {
		fmt.Errorf("Error: %s", err.Error())
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatalf("An error occurred while syncing the database.. %v", err)
	}

	db.Close()
	driver.Close()
	m.Close()
	fmt.Println("Migrate successfully!")
}
