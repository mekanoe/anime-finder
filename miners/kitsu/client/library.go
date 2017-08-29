package kitsuclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	ErrTimedOut = errors.New("kitsu client: multi-fetch timed out")
)

// https://kitsu.io/api/edge/users/2/library-entries?page[limit]=500&page[offset]=...
func (k *Kitsu) GetUserLibraryEntriesAll(id string) (le LibraryEntriesWithAnime, err error) {
	end := false

	go func() {
		time.AfterFunc(5*time.Second, func() {
			if !end {
				log.Println("kitsu client: GetUserLibraryEntriesAll timed out, user =>", id)
			}

			end = true
			err = ErrTimedOut
		})
	}()

	limit := 500
	offset := 0
	for end == false {
		// log.Println("getting data for", id, "||", offset, "->", limit+offset)
		tle, err := k.GetUserLibraryEntriesOffset(id, limit, offset)
		if err != nil {
			return le, err
		}

		if offset == 0 {
			le = tle
		}

		le.Data = append(le.Data, tle.Data...)

		offset += limit

		if le.Meta.Count < offset {
			end = true
			break
		}
	}

	return
}

func (k *Kitsu) GetUserLibraryEntriesOffset(id string, limit int, offset int) (le LibraryEntriesWithAnime, err error) {
	rsp, err := http.Get(fmt.Sprintf("%s/users/%s/library-entries?include=anime&page[limit]=%d&page[offset]=%d", k.baseURL, id, limit, offset))
	if err != nil {
		return
	}

	err = json.NewDecoder(rsp.Body).Decode(&le)
	if err != nil {
		return
	}

	return le, nil
}
