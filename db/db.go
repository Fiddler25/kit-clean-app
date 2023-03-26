package db

import (
	"clean-architecture-sample/ent"
	"context"
	"log"
)

func New() (context.Context, *ent.Client) {
	client, err := ent.Open("mysql", "root@tcp(127.0.0.1:13306)/clean-architecture-sample?charset=utf8mb4&parseTime=True")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return ctx, client
}
