package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

/* Define User Structure -------------------------------- */

type User struct {
	ID       int
	Email    string
	Password string
}

/* Storage Layer ---------------------------------------- */
var userStorage []User

func main() {
	fmt.Println("Hello, Todo App!")

	//define a command
	command := flag.String("command", "no command", "create a new task")
	flag.Parse() //now we have the info of command

	/* For Loop For Running Command ----------------------- */
	for {
		runCommand(*command)

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("please enter another command or exit:")
		scanner.Scan()
		*command = scanner.Text()
	}

	fmt.Printf("userStorage: %+v\n", userStorage)

}

//how make our code more clear?
// use functions

func runCommand(command string) {
	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser()
	case "login":
		login()

	case "exit":
		os.Exit(0)
	default:
		fmt.Println("Unknown command", command)
	}
}

func createTask() {
	scanner := bufio.NewScanner(os.Stdin)

	var name, duedate, category string

	fmt.Println("Please Enter task name: ")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("Please Enter task category: ")
	scanner.Scan()
	category = scanner.Text()

	fmt.Println("Please Enter task due date: ")
	scanner.Scan()
	duedate = scanner.Text()

	fmt.Println("task: ", name, category, duedate)
}

func createCategory() {

	scanner := bufio.NewScanner(os.Stdin)
	var title, color string

	fmt.Println("Please enter category title: ")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("Please enter category color: ")
	scanner.Scan()
	color = scanner.Text()

	fmt.Println("category: ", title, color)
}

func registerUser() {
	scanner := bufio.NewScanner(os.Stdin)

	var id, email, password string

	fmt.Println("Please enter your email: ")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("Please enter your password: ")
	scanner.Scan()
	password = scanner.Text()

	id = email
	fmt.Println("Your information: ", id, email, password)

	//store user
	user := User{
		ID:       len(userStorage) + 1,
		Email:    email,
		Password: password,
	}

	userStorage = append(userStorage, user) // append new user to userStorage slice

}

func login() {
	scanner := bufio.NewScanner(os.Stdin)

	var id, email, password string

	fmt.Println("Please enter your email: ")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("Please enter your password: ")
	scanner.Scan()
	password = scanner.Text()

	fmt.Println("category: ", id, email, password)
}
