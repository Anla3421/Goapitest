package Sql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	// ID int `json:"id"`
	// uid            string `json:"uid"`
	Name string `json:"name"`
	// password       string `json:"password"`
	// gender         string `json:"gender"`
	// level          string `json:"level"`
	// position       string `json:"position"`
	// remember_check string `json:"remember_check"`
	// status         string `json:"status"`
	// cellphone      string `json:"cellphone"`
	// remember_token string `json:"remember_token"`
	// api_token      string `json:"api_token"`
}

// var mysqlConn *sql.DB

// func Init() {
// 	fmt.Println("MySQL initial")
// 	CreateConn()
// }

// func CreateConn() {
// 	db, err := sql.Open("mysql", "root:adminstrator@tcp(127.0.0.1:3306)/jared")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	mysqlConn = db
// 	fmt.Println("success connected to MySQL")
// }

func Users() User {
	db, err := sql.Open("mysql", "root:adminstrator@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("success connected to MySQL")

	//insert
	// insert, err := db.Query("select users from users")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	panic(err.Error())
	// }
	// fmt.Println("insert success")

	results, err := db.Query("select name from users")
	if err != nil {
		panic(err.Error())
	}
	var user User
	for results.Next() {
		err = results.Scan(&user.Name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.Name)
	}
	return user

	// defer insert.Close()

}

// func Users() {
// 	var tag Response
// 	err := mysqlConn.QueryRow("Select * FROM users where ID = ?", 2).Scan(&tag.ID)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		// return Response{}
// 	}
// 	// return Response{}
// 	fmt.Println(tag.ID)

// }
