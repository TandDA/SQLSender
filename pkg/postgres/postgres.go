package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	connStr = "user=postgres password=123 dbname=postgres sslmode=disable"
)

type Postgres struct {
	db *sql.DB
}

func New() *Postgres {
	postgres := Postgres{}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Print("Failed to open db connection:", err.Error())
	}
	postgres.db = db
	return &postgres
}

func (psg *Postgres) Execute(sql string) [][]string {
	rows, err := psg.db.Query(sql)
	if err != nil {
		fmt.Print("Failed to execute sql:", err.Error())
	}
	defer rows.Close()
	fmt.Println(sql)
	cols, _ := rows.Columns()
	result := make([][]string, 0)
	result = append(result, cols)
	for rows.Next() {
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		rows.Scan(columnPointers...)
		result = append(result, columns)
	}
	return result
}
