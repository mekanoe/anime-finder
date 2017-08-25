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

type LibraryEntries struct {
	Data []LibraryEntry `json:"data"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`

	Links struct {
		First string `json:"first"`
		Next  string `json:"next"`
		Last  string `json:"last"`
	} `json:"links"`
}

// https://kitsu.io/api/edge/users/2/library-entries?page[limit]=500&page[offset]=...
func (k *Kitsu) GetUserLibraryEntriesAll(id string) (le LibraryEntries, err error) {
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

func (k *Kitsu) GetUserLibraryEntriesOffset(id string, limit int, offset int) (le LibraryEntries, err error) {
	rsp, err := http.Get(fmt.Sprintf("%s/users/%s/library-entries?page[limit]=%d&page[offset]=%d", k.baseURL, id, limit, offset))
	if err != nil {
		return
	}

	err = json.NewDecoder(rsp.Body).Decode(&le)
	if err != nil {
		return
	}

	return le, nil
}

func animeFetcher(id int, jobs <-chan LibraryEntry, results chan<- Anime) {
	for j := range jobs {
		log.Println("worker", id, "on", j.ID)

		rsp, err := http.Get(j.Relationships.Anime.Links.Related)
		if err != nil {
			log.Println("worker", id, "errored on rsp for", j.ID, "=>", err)
		}

		var ar AnimeResponse
		err = json.NewDecoder(rsp.Body).Decode(&ar)
		if err != nil {
			log.Println("worker", id, "errored on decode for", j.ID, "=>", err)
		}

		if ar.Data.ID != "" {
			results <- ar.Data
		}
		log.Println("worker", id, "on", j.ID, "DONE")
	}
}

func (k *Kitsu) GetAnimesFromLibraryEntries(le LibraryEntries) (as []Anime, err error) {
	jobs := make(chan LibraryEntry, len(le.Data))
	results := make(chan Anime)

	wc := 10

	if le.Meta.Count < 10 {
		wc = le.Meta.Count
	}

	// spawn workers
	for w := 1; w <= wc; w++ {
		go animeFetcher(w, jobs, results)
	}

	// enqueue jobs
	for _, e := range le.Data {
		jobs <- e
	}
	close(jobs)

	// output results
	for _ = range le.Data {
		as = append(as, <-results)
	}

	return
}
