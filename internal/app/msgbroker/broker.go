package msgbroker

import "github.com/nats-io/stan.go"

type Broker struct {
	brokerRepository *BrokerRepository
	stan             stan.Conn
}

func (b *Broker) Msg() BrokerRepositoryInterface {
	if b.brokerRepository != nil {
		return b.brokerRepository
	}

	b.brokerRepository = &BrokerRepository{
		broker: b,
	}

	return b.brokerRepository
}

func New(conn *stan.Conn) *Broker {
	return &Broker{
		stan: *conn,
	}
}
