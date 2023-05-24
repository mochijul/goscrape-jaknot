package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var exec string

func DBConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/")

	if err != nil {
		fmt.Println(err.Error())
	}

	return db
}

func RunExec(exec string, description string) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	db := DBConnect()
	res, err := db.ExecContext(ctx, exec)
	if err != nil {
		log.Printf("Error %s when %s \n", err, description)
		return
	}
	no, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
		return
	}
	log.Printf("rows affected %d , %s \n", no, description)
	db.Close()
}

func RunQuery(query string, description string) *sql.Rows {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	db := DBConnect()
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when %s \n", err, description)
	}
	defer rows.Close()
	// for rows.Next() {
	// 	var name string
	// 	var id int
	// 	err = rows.Scan(&id, &name)
	// 	if err != nil {
	// 		log.Printf("Error %s when fetching data", err)
	// 		return
	// 	}
	// 	log.Printf("ID: %d, Name: %s\n", id, name)
	// }
	db.Close()
	return rows
}
