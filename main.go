package main

import (
	_ "embed"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

//var schema = `
//CREATE TABLE test (
//   id          INT       NOT NULL AUTO_INCREMENT PRIMARY KEY,
//   data        JSON      NOT NULL,
//   create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
//)
//`

var schema = `
CREATE TABLE test (
  id          INT       NOT NULL AUTO_INCREMENT PRIMARY KEY,
  data        JSON      NOT NULL,
  create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  INDEX(create_time)
)
`

//go:embed data.json
var bigJSON string

type Test struct {
	ID         int64     `db:"id"`
	Data       string    `db:"data"`
	CreateTime time.Time `db:"create_time"`
}

func main() {
	fmt.Println()
	defer fmt.Println()

	db, err := sqlx.Connect("mysql", "root:@(localhost:3306)/test?parseTime=true")
	if err != nil {
		panic(err)
	}

	//db.MustExec(schema)

	//for i := 0; i < 10_000; i++ {
	//	db.MustExec("INSERT INTO test (data) VALUES (?)", bigJSON)
	//}

	//for i := 0; i < 10_000; i++ {
	//	db.MustExec("INSERT INTO test (data) VALUES (?)", "{}")
	//}

	var count int64
	err = db.Get(&count, "SELECT COUNT(*) FROM test")
	if err != nil {
		panic(err)
	}

	fmt.Printf("COUNT(*) = %v\n\n", count)

	//var data []Test
	//err = db.Select(&data, "SELECT * FROM test ORDER BY create_time DESC LIMIT 10 OFFSET 0")
	//if err != nil {
	//	panic(err)
	//}

	timer(func() {
		var data []Test
		err = db.Select(&data, "SELECT * FROM test ORDER BY create_time DESC LIMIT 10 OFFSET 0")
		if err != nil {
			panic(err)
		}
	})

	timer(func() {
		var data []Test
		err = db.Select(&data, "SELECT * FROM test ORDER BY create_time DESC LIMIT 10 OFFSET 1000")
		if err != nil {
			panic(err)
		}
	})

	//timer(func() {
	//	var data []Test
	//	err = db.Select(&data, "SELECT * FROM test ORDER BY create_time DESC LIMIT 10 OFFSET 9000")
	//	if err != nil {
	//		panic(err)
	//	}
	//})

	timer(func() {
		var data []Test
		err = db.Select(&data, "SELECT * FROM test FORCE INDEX (create_time) ORDER BY create_time DESC LIMIT 10 OFFSET 1000")
		if err != nil {
			panic(err)
		}
	})

	//timer(func() {
	//	var data []Test
	//	err = db.Select(&data, "SELECT * FROM test FORCE INDEX (create_time) ORDER BY create_time DESC LIMIT 10 OFFSET 9000")
	//	if err != nil {
	//		panic(err)
	//	}
	//})

	timer(func() {
		var data []Test
		err = db.Select(&data, "SELECT * FROM test JOIN (SELECT id FROM test ORDER BY create_time DESC LIMIT 10 OFFSET 1000) AS t ON test.id = t.id")
		if err != nil {
			panic(err)
		}
	})
}

func timer(f func()) {
	start := time.Now()
	defer func() {
		end := time.Now()
		fmt.Printf("elapsed: %v\n\n", end.Sub(start).String())
	}()

	f()
}
