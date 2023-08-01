package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func DbConnection() {
	user := "root"
	password := "password"
	dbName := "Golang"

	cfg := mysql.Config{
		User:   user,
		Passwd: password,
		Net:    "tcp",
		Addr:   "db:3306", // Utilisez simplement "db" au lieu de "127.0.0.1" car "db" est le nom du service dans Docker Compose.
		DBName: dbName,
	}

	// Ouvrez la connexion à la base de données MySQL
	var err error
	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println("Cant open db")
		log.Fatal(err)
	}

	// Vérifiez la connexion
	pingErr := DB.Ping()
	if pingErr != nil {
		// Vérifiez si l'erreur est de type *mysql.MySQLError
		fmt.Println("Can't ping DB")
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to the database!")
}
