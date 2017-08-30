package run

import (
	"context"
	"os"
	"testing"
)

var testS *Service

func TestMain(m *testing.M) {
	testS = &Service{
		NATS:   getNATS(),
		Dgraph: getDG(),
	}

	os.Exit(m.Run())
}

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
