package main

import (
	"database/sql"
	"dbutils"
	"fmt"
	"os"

	"config"
	"security"

	_ "github.com/mattn/go-sqlite3"
)

type configuration struct {
	DbPath string
}

func main() {
	args := os.Args

	conf := configuration{}
	if len(args) > 1 {
		config.Load(args[1], &conf)
	} else {
		fmt.Printf(args[0])
		os.Exit(1)
	}

	fmt.Println(conf.DbPath)
	fmt.Printf("hello, world\n")
	//dbutils.Test1("bob")
	dbutils.Test2("bill")

	db, err := sql.Open("sqlite3", conf.DbPath)
	dbutils.CheckErr(err)
	defer db.Close()

	users := security.GetAllUsers(db)

	for _, u := range users {
		fmt.Printf("User %d %s %s\n", u.ID, u.Username, u.Name)
	}

}
