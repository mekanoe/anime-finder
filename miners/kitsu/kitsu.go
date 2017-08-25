package kitsu

import (
	"context"
	"log"

	"github.com/kayteh/anime-finder/miners"
)

func init() {
	miners.RegisterMinerClient("kitsu", Handle)
}

func Handle(ctx context.Context) {
	msg := ctx.Value(miners.CtxMsg).(miners.MinerMsg)

	log.Println(msg)
}
