package msgbroker

import stan "github.com/nats-io/stan.go"

type BrokerRepository struct {
	broker *Broker
}

func (r *BrokerRepository) Subscribe(q string, handler func(msg *stan.Msg)) {
	r.broker.stan.Subscribe(q, handler)
}
