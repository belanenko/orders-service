package main

import (
	"database/sql"
	"log"

	"github.com/belanenko/orders-service/internal/app/apiserver"
	"github.com/belanenko/orders-service/internal/app/msgbroker"
	"github.com/belanenko/orders-service/internal/app/store"
	"github.com/belanenko/orders-service/internal/app/store/localstore"
	"github.com/belanenko/orders-service/internal/app/store/sqlstore"
	"github.com/nats-io/stan.go"

	_ "github.com/lib/pq" // ...
)

func main() {
	localstore := localstore.New()

	sqlConn, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=admin password=123 dbname=dev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := sqlConn.Ping(); err != nil {
		log.Fatal(err)
	}

	sqlstore := sqlstore.New(sqlConn)
	storage := store.New(localstore, sqlstore)

	stanConn, err := stan.Connect("test-cluster", "me")
	if err != nil {
		log.Fatal(err)
	}
	broker := msgbroker.New(&stanConn)

	config := apiserver.NewConfig()
	config.BindAddr = ":8081"
	apiserver := apiserver.New(storage, broker, config)
	if err := apiserver.Start(); err != nil {
		log.Fatal(err)
	}
}
