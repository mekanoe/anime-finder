// Package main is the service for the data miner.
package main

import (
	"log"
	"sync"
	"time"

	"github.com/kayteh/anime-finder/miners"
	_ "github.com/kayteh/anime-finder/miners/kitsu"
	"github.com/kayteh/anime-finder/util"
	"github.com/nats-io/go-nats"
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

	_, err = miners.NewMiner(nc)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("running worker")

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
