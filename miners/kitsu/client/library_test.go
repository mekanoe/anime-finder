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

	le, err := kc.GetUserLibraryEntriesAll("161540")
	if err != nil {
		t.Error(err)
		return
	}

	for _, a := range le.Included {
		fmt.Println(a.Attributes.CanonicalTitle)
	}
}

// func TestLibGetAllAnime(t *testing.T) {
// 	t.Skip()

// 	le, err := kc.GetUserLibraryEntriesAll("161540")
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}

// 	a, err := kc.GetAnimesFromLibraryEntries(le)

// 	fmt.Println(a)
// }
