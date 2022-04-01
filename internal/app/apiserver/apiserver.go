package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/belanenko/orders-service/internal/app/model"
	"github.com/belanenko/orders-service/internal/app/msgbroker"
	"github.com/belanenko/orders-service/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	Store  store.StoreInterface
	Broker msgbroker.BrokerInterface
	Logger *logrus.Logger
	Router *mux.Router
}

func New(store store.StoreInterface, borker msgbroker.BrokerInterface, config *Config) *APIServer {
	return &APIServer{
		config: config,
		Store:  store,
		Broker: borker,
		Router: mux.NewRouter(),
		Logger: logrus.New(),
	}
}

func (s *APIServer) Start() error {
	s.configureRouter()
	s.configureSubscribes()
	s.Logger.Infof("Start to listen %s", s.config.BindAddr)
	return http.ListenAndServe(s.config.BindAddr, s.Router)
}

func (s *APIServer) configureRouter() {
	s.Router.HandleFunc("/api/orders", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Add("ALLOW", "GET")
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		id := r.URL.Query().Get("id")
		if id != "" {
			item, _ := s.Store.Item().Get(id)
			if item == nil {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			w.Write(item.Json())
			return
		}

		items, err := s.Store.Item().GetAll()
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}

		jsonStr, err := json.Marshal(items)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		w.Write(jsonStr)
	})

	fileServer := http.FileServer(http.Dir("./web/static"))
	s.Router.Handle("/", fileServer)
}

func (s *APIServer) configureSubscribes() {
	s.Broker.Msg().Subscribe("items", func(msg *stan.Msg) {
		var item model.Order
		json.Unmarshal(msg.Data, &item)
		if err := s.Store.Item().Set(item.OrderUID, &item); err != nil {
			s.Logger.Infof("Recived msg bad format: %s", err)
		}
	})
}
