package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "postgres"
	dbname   = "thai"
)

func getDatabaseConn() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func testDbPerfomanceRead(db *sql.DB) {
	retries := []int{10, 50, 100, 500, 1000, 5000}
	query := `select * from book.tickets`

	timeStart := time.Now()
	db.QueryRow(query)
	timeEnd := time.Now()
	fmt.Printf("single query time: %v ms\n", timeEnd.Sub(timeStart).Milliseconds())

	for i := 0; i < len(retries); i++ {
		timeStart := time.Now()
		for j := 0; j < retries[i]; j++ {
			db.QueryRow(query)
		}
		timeEnd := time.Now()

		fmt.Printf("%d queries time: %v ms\n", retries[i], timeEnd.Sub(timeStart).Milliseconds())
	}
}

func testDbPerfomanceWrite(db *sql.DB) {
	retries := []int{10, 50, 100, 500, 1000, 5000}
	query := `INSERT INTO book.tickets (fkRide, fio, contact, fkSeat)
				VALUES (
				ceil(random()*100)
				, (array(SELECT fam FROM book.fam))[ceil(random()*110)]::text || ' ' ||
				(array(SELECT nam FROM book.nam))[ceil(random()*110)]::text
				,('{"phone":"+7' || (1000000000::bigint + floor(random()*9000000000)::bigint)::text || '"}')::jsonb
				, ceil(random()*100))`

	timeStart := time.Now()
	db.QueryRow(query)
	timeEnd := time.Now()
	fmt.Printf("single query time: %v ms\n", timeEnd.Sub(timeStart).Milliseconds())

	for i := 0; i < len(retries); i++ {
		timeStart := time.Now()

		for j := 0; j < retries[i]; j++ {
			db.QueryRow(query)
		}
		timeEnd := time.Now()
		fmt.Printf("%d queries time: %v ms\n", retries[i], timeEnd.Sub(timeStart).Milliseconds())
	}
}

func main() {
	db, err := getDatabaseConn()

	if err != nil {
		fmt.Printf("could not obtain db connection: %s", err)
		return
	}
	defer db.Close()

	testDbPerfomanceRead(db)
	testDbPerfomanceWrite(db)
}
