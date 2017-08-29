package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/kayteh/anime-finder/miners/kitsu/client"

	"github.com/imdario/mergo"
	nats "github.com/nats-io/go-nats"
)

type Anime = kitsuclient.Anime

type ctxStr string

var (
	CtxSvc = ctxStr("service")
	CtxMsg = ctxStr("message")
)

type ClientHandler func(ctx context.Context)

type Service struct {
	n *nats.Conn
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
	s.n.QueueSubscribe("ingress:user", "graphsvc", s.wrap(s.userHandler))
	s.n.QueueSubscribe("ingress:userentries", "graphsvc", s.wrap(s.userEntryHandler))
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

func getDGContext() context.Context {
	return context.Background()
}

func (s *Service) animeHandler(ctx context.Context) {
	msg := ctx.Value(CtxMsg).(SvcMsg)

	asTmp := msg.Data["anime"].([]interface{})

	for _, at := range asTmp {

		atm := at.(map[string]interface{})
		var a Anime

		err := mergo.MapWithOverwrite(a, atm)
		if err != nil {
			log.Println("anime: error merging map into value", err)
			return
		}

	}

}

func (s *Service) userHandler(ctx context.Context) {

}

func (s *Service) userEntryHandler(ctx context.Context) {

}

func (s *Service) requestHandler(msg *nats.Msg) {
}
