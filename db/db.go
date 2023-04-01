package db

import (
	"context"
	"kit-clean-app/ent"
	"log"
)

type DB struct {
	*ent.Client
}

func New() (*DB, context.Context) {
	client, err := ent.Open("mysql", "root@tcp(127.0.0.1:13306)/kit-clean-app?charset=utf8mb4&parseTime=True")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}

	ctx := ent.NewContext(context.Background(), client)

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return &DB{client}, ctx
}

func Client(ctx context.Context) *ent.Client {
	if tx := ent.TxFromContext(ctx); tx != nil {
		return tx.Client()
	}
	return ent.FromContext(ctx)
}
