package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB
var err error
var connectionTimeout int = 30

func main() {

	query := url.Values{}
	query.Add("connection timeout", fmt.Sprintf("%d", connectionTimeout))

	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword("sa", "udanjo2011@"),
		Host:   fmt.Sprintf("%s:%d", "localhost", 5434),
		// Path:  instance, // if connecting to an instance instead of a port
		RawQuery: query.Encode(),
	}

	connStr := u.String()

	// Create connection
	db, err = sql.Open("mssql", connStr)
	if err != nil {
		fmt.Println("Error creating connection pool: ", err.Error())
	} else {
		fmt.Println("Conetado com sucesso!")
	}

	SimpleSelect(db)
	fmt.Println(" ")
	MultiplesSelect(db)

	// Close database connection
	defer db.Close()
}

func SimpleSelect(db *sql.DB) {
	var name string

	// Select
	if err := db.QueryRow("SELECT TOP 1 asset_ds FROM ProjectSend.dbo.WalleyAsset").Scan(&name); err != nil {
		fmt.Println("Erro gerado ", err)
	}

	fmt.Println(name)
}

func MultiplesSelect(db *sql.DB) {
	ctx := context.Background()

	// Select Multiples Rows
	row, _ := db.QueryContext(ctx, "SELECT asset_ds FROM ProjectSend.dbo.WalleyAsset")
	count := 0

	for row.Next() {
		var name string
		row.Scan(&name)

		fmt.Println(name)
		count++
	}

	fmt.Println("\nTotal de ", count)
}
