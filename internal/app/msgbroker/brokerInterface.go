package msgbroker

type BrokerInterface interface {
	Msg() BrokerRepositoryInterface
}
