package msg

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/nats-io/nats.go"
	_ "github.com/rs/zerolog/log"
)

type NATS struct {
	Conn *nats.Conn
	JS   nats.JetStreamContext
}

func Connect(url string) (*NATS, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	js, err := nc.JetStream()
	if err != nil {
		nc.Close()
		return nil, err
	}
	return &NATS{Conn: nc, JS: js}, nil
}

func (n *NATS) Close() { n.Conn.Close() }

func (n *NATS) PublishEvent(ctx context.Context, subject string, data any) (seq uint64, err error) {

	if n == nil || n.JS == nil {
		return 0, fmt.Errorf("nats: JetStream context is nil (forgot Connect?)")
	}
	body, err := json.Marshal(data)
	if err != nil {
		return 0, fmt.Errorf("marshal: %w", err)
	}
	ack, err := n.JS.Publish(subject, body)
	if err != nil {
		return 0, fmt.Errorf("jetstream publish: %w", err)
	}
	return ack.Sequence, nil
}
