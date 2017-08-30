package run

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"time"

	dgraph "github.com/dgraph-io/dgraph/client"
	"github.com/imdario/mergo"
	"github.com/kayteh/anime-finder/util"
	"github.com/kayteh/anime-finder/util/dqrack"
	nats "github.com/nats-io/go-nats"
	"google.golang.org/grpc"
)

type ctxStr string

var (
	CtxSvc = ctxStr("service")
	CtxMsg = ctxStr("message")
)

type ClientHandler func(ctx context.Context)

type Service struct {
	NATS   *nats.Conn
	Dgraph *dgraph.Dgraph
	Dq     *dqrack.Dqrack
}

func NewService() *Service {
	s := &Service{
		NATS:   getNATS(),
		Dgraph: getDG(),
	}

	s.Dq = dqrack.New(s.Dgraph)

	return s
}

func getNATS() *nats.Conn {
	opts := nats.Options{
		Url:            "nats://" + util.GetenvOrDie("NATS_ADDR"),
		AllowReconnect: true,
		MaxReconnect:   10,
		ReconnectWait:  5 * time.Second,
		Timeout:        1 * time.Second,
	}

	nc, err := opts.Connect()
	if err != nil {
		log.Fatalln("couldn't connect to NATS,", err)
	}

	return nc
}

func getDG() *dgraph.Dgraph {
	g, err := grpc.Dial(util.GetenvOrDie("DGRAPH_ADDR"), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return dgraph.NewDgraphClient([]*grpc.ClientConn{g}, dgraph.DefaultOptions, ".tmp")
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
	s.NATS.QueueSubscribe("ingress:anime", "graphsvc", s.wrap(s.animeHandler))
	s.NATS.QueueSubscribe("ingress:user", "graphsvc", s.wrap(s.userHandler))
	s.NATS.QueueSubscribe("ingress:userentries", "graphsvc", s.wrap(s.userEntryHandler))
	s.NATS.Subscribe("graphsvc:req:*", s.requestHandler)
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
	return
	// msg := ctx.Value(CtxMsg).(SvcMsg)

	// asTmp := msg.Data["anime"].([]interface{})

	// for _, at := range asTmp {
	// 	var a types.Anime
	// 	err := jsonCast(&a, at)
	// 	if err != nil {
	// 		log.Println("anime: failed json cast:", err)
	// 		return
	// 	}

	// 	buf := bytes.Buffer{}
	// 	a.WriteMutation(&buf)

	// 	r := &dgraph.Req{}
	// 	r.SetQuery(buf.String())
	// 	// log.Panicln(buf.String())
	// 	_, err = s.Dgraph.Run(getDGContext(), r)
	// 	if err != nil {
	// 		log.Println("anime: mutation write failed:", err)
	// 		return
	// 	}
	// }

}

func (s *Service) userHandler(ctx context.Context) {

}

func (s *Service) userEntryHandler(ctx context.Context) {

}

func (s *Service) requestHandler(msg *nats.Msg) {
}

// Helper func for casting weird types into good types.
func jsonCast(dst, src interface{}) error {
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(src)
	if err != nil {
		return err
	}

	return json.NewDecoder(buf).Decode(dst)
}
