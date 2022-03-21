package main

import (
	"bufio"
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {
	connectionString := queryConnectionString()
	config, err := pgx.ParseConfig(os.ExpandEnv(connectionString))
	if err != nil {
		log.Fatal("error configuration the database: ", err)
	}
	connection , err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	defer connection.Close(context.Background())
	log.Println("Hey! You successfully connected to your Cockroach DB cluster")
}

func queryConnectionString() (string) {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	scanner := bufio.NewScanner(os.Stdin)
	log.Println("Enter a connection string: ")
	scanner.Scan()
	return scanner.Text()
}