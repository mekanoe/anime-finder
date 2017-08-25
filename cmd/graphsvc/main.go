package main

import (
	"log"
	"sync"
	"time"

	"github.com/kayteh/anime-finder/util"
	nats "github.com/nats-io/go-nats"
)

func main() {
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

	svc := &Service{
		n: nc,
	}

	svc.Mount()

	log.Println("running worker")

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
