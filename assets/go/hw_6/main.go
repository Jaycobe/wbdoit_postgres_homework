package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host               = "localhost"
	portWithIndexes    = 5000
	portWithoutIndexes = 5001
	user               = "postgres"
	password           = "postgres"
	dbname             = "thai"
)

func getDatabaseConn(port int) (*sql.DB, error) {
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
	retries := []int{10}
	query := `
			WITH all_place AS (
				SELECT count(s.id) as all_place, s.fkbus as fkbus
				FROM book.seat s
				group by s.fkbus
			),
				 order_place AS (
					 SELECT count(t.id) as order_place, t.fkride
					 FROM book.tickets t
					 group by t.fkride
				 )
			SELECT r.id, r.startdate as depart_date, bs.city || ', ' || bs.name as busstation,
				   t.order_place, st.all_place
			FROM book.ride r
					 JOIN book.schedule as s
						  on r.fkschedule = s.id
					 JOIN book.busroute br
						  on s.fkroute = br.id
					 JOIN book.busstation bs
						  on br.fkbusstationfrom = bs.id
					 JOIN order_place t
						  on t.fkride = r.id
					 JOIN all_place st
						  on r.fkbus = st.fkbus
			GROUP BY r.id, r.startdate, bs.city || ', ' || bs.name, t.order_place,st.all_place
			ORDER BY r.startdate
			limit 10;`

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
	dbWithIndexes, err := getDatabaseConn(portWithIndexes)
	if err != nil {
		fmt.Printf("could not obtain db connection on port %d: %s", portWithIndexes, err)
		return
	}

	dbWithoutIndexes, err := getDatabaseConn(portWithoutIndexes)
	if err != nil {
		fmt.Printf("could not obtain db connection on port %d: %s", portWithoutIndexes, err)
		return
	}
	defer dbWithoutIndexes.Close()
	defer dbWithIndexes.Close()

	testDbPerfomanceRead(dbWithIndexes)
	testDbPerfomanceRead(dbWithoutIndexes)
}
