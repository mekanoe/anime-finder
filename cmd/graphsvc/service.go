package main

import (
	"context"
	"encoding/json"
	"log"

	dgraph "github.com/dgraph-io/dgraph/client"
	"github.com/imdario/mergo"
	"github.com/kayteh/anime-finder/types"
	nats "github.com/nats-io/go-nats"
)

type ctxStr string

var (
	CtxSvc = ctxStr("service")
	CtxMsg = ctxStr("message")
)

type ClientHandler func(ctx context.Context)

type Service struct {
	n  *nats.Conn
	dg *dgraph.Dgraph
}

type SvcMsg struct {
	Type    string
	Service string
	Data    map[string]interface{}
}

func (s SvcMsg) Scan(v interface{}) error {
	return mergo.MapWithOverwrite(v, s.Data)
}

func (s *Service) Mount() {
	s.n.QueueSubscribe("ingress:anime", "graphsvc", s.wrap(s.animeHandler))
	s.n.QueueSubscribe("ingress:users", "graphsvc", s.wrap(s.usersHandler))
	s.n.Subscribe("graphsvc:req:*", s.requestHandler)
}

func (s *Service) wrap(h ClientHandler) nats.MsgHandler {
	return func(msg *nats.Msg) {
		var cm SvcMsg
		err := json.Unmarshal(msg.Data, &cm)
		if err != nil {
			log.Println("wrapper err: json unmarshal failed:", err)
			return
		}

		ctx := context.WithValue(context.Background(), CtxMsg, cm)

		h(ctx)
	}
}

func (s *Service) animeHandler(ctx context.Context) {
	msg := ctx.Value(CtxMsg).(SvcMsg)

	as := make([]types.Anime)
}

func (s *Service) usersHandler(ctx context.Context) {

}

func (s *Service) requestHandler(msg *nats.Msg) {
}
