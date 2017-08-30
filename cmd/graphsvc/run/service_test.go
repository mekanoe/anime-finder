package run

import (
	"context"
	"testing"
)

var testS = NewService()

func getTestContext(dat map[string]interface{}) context.Context {
	return context.WithValue(context.Background(), CtxMsg, SvcMsg{
		Data: dat,
	})
}

func TestAnimeHandler(t *testing.T) {
	testS.animeHandler(getTestContext(
		map[string]interface{}{
			"anime": []interface{}{
				testAnime,
			},
		},
	))
}
