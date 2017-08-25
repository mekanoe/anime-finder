package kitsuclient

import (
	"fmt"
	"testing"
)

var (
	kc, _ = NewClient()
)

func TestLibGetAll(t *testing.T) {
	t.Skip()

	_, err := kc.GetUserLibraryEntriesAll("161540")
	if err != nil {
		t.Error(err)
		return
	}

	// fmt.Println(le)
}

func TestLibGetAllAnime(t *testing.T) {
	t.Skip()

	le, err := kc.GetUserLibraryEntriesAll("161540")
	if err != nil {
		t.Error(err)
		return
	}

	a, err := kc.GetAnimesFromLibraryEntries(le)

	fmt.Println(a)
}
