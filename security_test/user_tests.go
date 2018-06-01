package main

import (
	"database/sql"
	"fmt"
	"os"
	"security"

	"flag"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func testUserCreate() {
	fmt.Println("Starting user create test")
	var user security.User
	user = security.User{Username: "testcreate", Name: "Test Create", IsActive: true, IsAdmin: false}
	var role = security.Role{Description: "Normal"}
	var err error
	var interaction security.Interaction
	err = role.GetByDescription(db)
	if err != nil {
		fmt.Printf("Error creating %s", err)
		os.Exit(1)
	}
	interaction, err = user.Create(db, role)

	if err != nil {
		fmt.Printf("Error creating %s", err)
		os.Exit(1)
	}
	fmt.Printf("Interaction created - key:%s", interaction.Key)
	interaction.Delete(db)
	user.DeleteLogs(db)
	user.DeleteRoles(db)
	user.Delete(db)
	fmt.Println("Finished user create test")
}

func testRoleCreate() {
	fmt.Println("Starting role create test")
	var role security.Role
	role = security.Role{Description: "testcreate", Data: "Test Create"}
	err := role.Create(db)

	if err != nil {
		fmt.Printf("Error creating %s", err)
		os.Exit(1)
	}

	role.Delete(db)
	fmt.Println("Finished role create test")
}

func testForgotPassword() {
	fmt.Println("Starting forgot password test")
	var user security.User
	user = security.User{Username: "testforgot", Name: "Test Forgot", IsActive: true, IsAdmin: false}

	var role = security.Role{Description: "Normal"}
	var err error
	var interaction security.Interaction
	err = role.GetByDescription(db)
	if err != nil {
		fmt.Printf("Error creating %s", err)
		os.Exit(1)
	}

	interaction, err = user.Create(db, role)

	if err != nil {
		fmt.Printf("Error creating %s", err)
		os.Exit(1)
	}

	interaction.Delete(db)

	interaction = user.ForgotPassword(db)
	fmt.Printf("Interaction created - key:%s", interaction.Key)
	interaction.Delete(db)
	user.DeleteLogs(db)
	user.DeleteRoles(db)
	user.Delete(db)

	fmt.Println("Finished forgot password test")
}

func main() {

	dbPath := flag.String("dbpath", "./security.db", "DB Path")
	flag.Parse()

	var err error
	fmt.Printf("db path %s\n", *dbPath)
	db, err = sql.Open("sqlite3", *dbPath)
	if err != nil {
		panic(err)
	}

	//testUserCreate()
	//testRoleCreate()
	testForgotPassword()
}
