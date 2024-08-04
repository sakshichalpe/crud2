package db

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/lib/pq"
// )

// var db *sql.DB

// func DBConnection() {

// 	db, err := sql.Open("postgres", "host=pg-cool-dev.postgres.database.azure.com port=5432 user=svc-cool password=no$a4YEc dbname=cool")
// 	if err != nil {
// 		fmt.Print("failed in open sql", err)
// 		return
// 	}
// 	err = db.Ping()
// 	if err != nil {
// 		fmt.Print("ping error::", err)
// 		return
// 	}
// 	fmt.Println("sucessfull")
// }

// // nc SetupDB() *sql.DB {
// // 	// db, err := sql.Open("mysql", "root:king@tcp(127.0.0.1:3306)/learn")
// // 	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=king dbname=postgres sslmode=disable")

// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// // 	return db
// // }
