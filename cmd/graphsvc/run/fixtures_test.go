package run

import (
	"testing"

	"github.com/kayteh/anime-finder/types"
)

var testAnime = map[string]interface{}{
	"attributes": map[string]interface{}{
		"abbreviatedTitles": []string{
			"Ajin",
		},
		"ageRating":      "R",
		"ageRatingGuide": "17+ (violence & profanity)",
		"averageRating":  "76.5",
		"canonicalTitle": "Ajin",
		"coverImage": map[string]string{
			"large":    "https://media.kitsu.io/anime/cover_images/11368/large.jpg?1452818186",
			"original": "https://media.kitsu.io/anime/cover_images/11368/original.jpg?1452818186",
			"small":    "https://media.kitsu.io/anime/cover_images/11368/small.jpg?1452818186",
			"tiny":     "https://media.kitsu.io/anime/cover_images/11368/tiny.jpg?1452818186",
		},
		"coverImageTopOffset": 200,
		"createdAt":           "2015-09-15T18:38:36.109Z",
		"endDate":             "2016-04-09",
		"episodeCount":        13,
		"episodeLength":       24,
		"favoritesCount":      86,
		"nsfw":                false,
		"popularityRank":      568,
		"posterImage": map[string]string{
			"large":    "https://media.kitsu.io/anime/poster_images/11368/large.jpg?1451099754",
			"medium":   "https://media.kitsu.io/anime/poster_images/11368/medium.jpg?1451099754",
			"original": "https://media.kitsu.io/anime/poster_images/11368/original.jpg?1451099754",
			"small":    "https://media.kitsu.io/anime/poster_images/11368/small.jpg?1451099754",
			"tiny":     "https://media.kitsu.io/anime/poster_images/11368/tiny.jpg?1451099754",
		},
		"ratingFrequencies": map[string]string{
			"10": "128",
			"11": "7",
			"12": "302",
			"13": "11",
			"14": "603",
			"15": "25",
			"16": "761",
			"17": "16",
			"18": "400",
			"19": "4",
			"2":  "18",
			"20": "444",
			"3":  "0",
			"4":  "28",
			"5":  "0",
			"6":  "22",
			"7":  "1",
			"8":  "74",
			"9":  "1",
		},
		"ratingRank": 965,
		"showType":   "TV",
		"slug":       "ajin",
		"startDate":  "2016-01-16",
		"status":     "finished",
		"subtype":    "TV",
		"synopsis":   "Mysterious immortal humans known as \"Ajin\" first appeared 17 years ago in Africa. Upon their discovery, they were labeled as a threat to mankind, as they might use their powers for evil and were incapable of being destroyed. Since then, whenever an Ajin is found within society, they are to be arrested and taken into custody immediately.\r\nStudying hard to become a doctor, Kei Nagai is a high schooler who knows very little about Ajin, only having seen them appear in the news every now and then. Students are taught that these creatures are not considered to be human, but Kei doesn't pay much attention in class. As a result, his perilously little grasp on this subject proves to be completely irrelevant when he survives an accident that was supposed to claim his life, signaling his rebirth as an Ajin and the start of his days of torment. However, as he finds himself alone on the run from the entire world, Kei soon realizes that more of his species may be a lot closer than he thinks.\r\n[Written by MAL Rewrite]",
		"tba":        nil,
		"titles": map[string]string{
			"en":    "Ajin: Demi-Human",
			"en_jp": "Ajin",
			"ja_jp": "\u4e9c\u4eba",
		},
		"updatedAt":      "2017-08-22T22:43:29.824Z",
		"userCount":      7345,
		"youtubeVideoId": "V62kcgCXNJU",
	},
	"id": "11368",
}

func TestJsonCast(t *testing.T) {
	var a types.Anime
	err := jsonCast(&a, testAnime)
	if err != nil {
		t.Fatal(err)
	}

	// log.Println(a.Attributes.CanonicalTitle)
}

func TestAnimeNode(t *testing.T) {
	var a types.Anime
	err := jsonCast(&a, testAnime)
	if err != nil {
		t.Fatal(err)
	}

	_, err = testS.Dq.GetNode(a)
	if err != nil {
		t.Error(err)
		return
	}
}
