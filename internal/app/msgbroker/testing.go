package msgbroker

import (
	"testing"

	"github.com/nats-io/stan.go"
)

func NewConn(t *testing.T, stanClusterID, clientID string) *stan.Conn {
	t.Helper()
	conn, err := stan.Connect(stanClusterID, clientID)
	if err != nil {
		t.Fatal(err)
	}

	return &conn
}
