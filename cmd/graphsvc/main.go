package main

import (
	"log"
	"sync"

	"github.com/kayteh/anime-finder/cmd/graphsvc/run"
)

func main() {

	run.NewService().Mount()

	log.Println("running worker")

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
