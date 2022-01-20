package main

import (
	"inventory-app/ent"

	"context"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:db?cache=shared&_fk=1")
	if err != nil {
		log.Fatal(err)
	}

	err = client.Schema.Create(context.Background())
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
