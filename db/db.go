package db

import (
	"context"
	"kit-clean-app/ent"
	"log"
)

func New() *ent.Client {
	client, err := ent.Open("mysql", "root@tcp(127.0.0.1:13306)/kit-clean-app?charset=utf8mb4&parseTime=True")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
