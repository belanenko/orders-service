package msgbroker

import (
	"testing"

	"github.com/nats-io/stan.go"
)

func TestBrokerRepository_Subscribe(t *testing.T) {
	conn := NewConn(t, "test-cluster", "cl")
	b := New(conn)

	data := "data"
	b.Msg().Subscribe("ttt", func(msg *stan.Msg) {
		if string(msg.Data) != data {
			t.Fatal("Data is missing")
		}
	})

	b.stan.Publish("ttt", []byte(data))
}
