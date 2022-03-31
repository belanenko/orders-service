package msgbroker

import stan "github.com/nats-io/stan.go"

type BrokerRepositoryInterface interface {
	Subscribe(string, func(msg *stan.Msg))
}
