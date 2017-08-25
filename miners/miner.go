package miners

import (
	"context"
	"encoding/json"
	"log"

	nats "github.com/nats-io/go-nats"
)

type CtxKey string

var (
	CtxMiner = CtxKey("miner")
	CtxMsg   = CtxKey("msg")
)

var handles = map[string]ClientHandler{}

type Miner struct {
	NATS *nats.Conn
}

type MinerMsg struct {
	Action string
	ID     string
}

type ClientHandler func(ctx context.Context)

// RegisterMinerClient adds support for side-effect miner clients to mount theirselves.
func RegisterMinerClient(topic string, h ClientHandler) {
	handles[topic] = h
}

// NewMiner creates a miner and mounts the various clients that were registered prior to calling this.
func NewMiner(nc *nats.Conn) (*Miner, error) {
	m := &Miner{
		NATS: nc,
	}

	for topic, h := range handles {
		_, err := nc.QueueSubscribe(topic, "dataminer", m.wrap(h))
		if err != nil {
			return nil, err
		}
	}

	return m, nil
}

func (m *Miner) wrap(h ClientHandler) nats.MsgHandler {
	return func(msg *nats.Msg) {
		ctx := context.WithValue(context.Background(), CtxMiner, m)

		var cm MinerMsg
		err := json.Unmarshal(msg.Data, &cm)
		if err != nil {
			log.Println("wrapper err: json unmarshal failed:", err)
			return
		}

		ctx = context.WithValue(ctx, CtxMsg, cm)

		h(ctx)
	}
}
