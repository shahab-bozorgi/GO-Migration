package db

import (
	"fmt"
	"log"
)


func (db *Database) DetectTableChanges() {
	rows, err := db.DB.Query(`SELECT table_name FROM information_schema.tables WHERE table_schema = 'public';`)
	if err != nil {
		log.Fatalf("Failed to fetch table information: %v", err)
	}
	defer rows.Close()

	log.Println("Tables in the database:")
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}
		log.Println("- ", tableName)
	}
}

func FetchTables(database *Database) ([]string, error) {
	query := `
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema = 'public';
	`

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, err
		}
		tables = append(tables, tableName)
	}

	return tables, nil
}

func (db *Database) CheckDatabaseStatus() {
	err := db.DB.Ping()
	if err != nil {
		log.Fatal("Database connection issue: ", err)
	} else {
		fmt.Println("Database connection is healthy")
	}
}
