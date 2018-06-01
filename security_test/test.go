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
	//dbutils.Test2("bill")

	db, err := sql.Open("sqlite3", conf.DbPath)
	dbutils.CheckErr(err)
	defer db.Close()

	var users []security.User

	// var user security.User
	// user = security.User{Username: "bill", Name: "Bill", IsActive: true, IsAdmin: false}
	// user.Create(db)
	// user = security.User{Username: "bob", Name: "Bob", IsActive: true, IsAdmin: false}
	// user.Create(db)
	// user = security.User{Username: "tim", Name: "Tim", IsActive: true, IsAdmin: false}
	// user.Create(db)
	// user = security.User{Username: "Ted", Name: "Ted", IsActive: true, IsAdmin: false}
	// user.Create(db)
	users = security.GetAllUsers(db)

	for _, u := range users {
		fmt.Printf("User %d %s %s\n", u.ID, u.Username, u.Name)
	}

	user := security.User{Username: "bill"}
	user.GetByUsername(db)
	interaction := user.ForgotPassword(db)

	fmt.Printf("%s\n", interaction.Key)

}
