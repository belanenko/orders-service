package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/belanenko/orders-service/internal/app/apiserver"
	"github.com/belanenko/orders-service/internal/app/msgbroker"
	"github.com/belanenko/orders-service/internal/app/store"
	"github.com/belanenko/orders-service/internal/app/store/localstore"
	"github.com/belanenko/orders-service/internal/app/store/sqlstore"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/nats-io/stan.go"

	_ "github.com/lib/pq" // ...
)

func main() {
	godotenv.Load()

	localstore := localstore.New()

	sqlConfig := sqlstore.NewConfig()
	env.Parse(sqlConfig)
	if sqlConfig.DatabaseUrl == "" {
		log.Fatal("connection string not found")
	}

	sqlConn, err := sql.Open("postgres", sqlConfig.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	if err := sqlConn.Ping(); err != nil {
		log.Fatal(err)
	}

	sqlstore := sqlstore.New(sqlConn)
	storage := store.New(localstore, sqlstore)

	stanConn, err := stan.Connect(os.Getenv("STAN_CLUSTER_ID"), os.Getenv("STAN_CLIENT_ID"))
	if err != nil {
		log.Fatal(err)
	}
	broker := msgbroker.New(&stanConn)

	config := apiserver.NewConfig()
	env.Parse(config)

	apiserver := apiserver.New(storage, broker, config)
	if err := apiserver.Start(); err != nil {
		log.Fatal(err)
	}
}
