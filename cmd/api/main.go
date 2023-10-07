package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	fmt.Println("Little Jira API using Golang...")
	app := &Application{}
	if err := Server(app); err != nil {
		log.Println(err.Error())
	}
}
