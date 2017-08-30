package kitsuclient

type LibraryEntry struct {
	Attributes struct {
		CreatedAt       string      `json:"createdAt"`
		FinishedAt      interface{} `json:"finishedAt"`
		Notes           interface{} `json:"notes"`
		Private         bool        `json:"private"`
		Progress        int64       `json:"progress"`
		ProgressedAt    string      `json:"progressedAt"`
		Rating          string      `json:"rating"`
		RatingTwenty    int64       `json:"ratingTwenty"`
		ReactionSkipped string      `json:"reactionSkipped"`
		ReconsumeCount  int64       `json:"reconsumeCount"`
		Reconsuming     bool        `json:"reconsuming"`
		StartedAt       string      `json:"startedAt"`
		Status          string      `json:"status"`
		UpdatedAt       string      `json:"updatedAt"`
		VolumesOwned    int64       `json:"volumesOwned"`
	} `json:"attributes"`
	ID    string `json:"id"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Relationships struct {
		Anime struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"anime"`
		Drama struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"drama"`
		Manga struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"manga"`
		Media struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"media"`
		MediaReaction struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"mediaReaction"`
		NextUnit struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"nextUnit"`
		Review struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"review"`
		Unit struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"unit"`
		User struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"user"`
	} `json:"relationships"`
	Type string `json:"type"`
}

type LibraryEntriesWithAnime struct {
	Data     []LibraryEntry `json:"data"`
	Included []Anime        `json:"included"`
	Links    struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
}

type AnimeResponse struct {
	Data Anime `json:"data"`
}

type Anime struct {
	Attributes struct {
		AbbreviatedTitles []string `json:"abbreviatedTitles"`
		AgeRating         string   `json:"ageRating"`
		AgeRatingGuide    string   `json:"ageRatingGuide"`
		AverageRating     string   `json:"averageRating"`
		CanonicalTitle    string   `json:"canonicalTitle"`
		CoverImage        struct {
			Large    string `json:"large"`
			Original string `json:"original"`
			Small    string `json:"small"`
			Tiny     string `json:"tiny"`
		} `json:"coverImage"`
		CoverImageTopOffset int64  `json:"coverImageTopOffset"`
		CreatedAt           string `json:"createdAt"`
		EndDate             string `json:"endDate"`
		EpisodeCount        int64  `json:"episodeCount"`
		EpisodeLength       int64  `json:"episodeLength"`
		FavoritesCount      int64  `json:"favoritesCount"`
		Nsfw                bool   `json:"nsfw"`
		PopularityRank      int64  `json:"popularityRank"`
		PosterImage         struct {
			Large    string `json:"large"`
			Medium   string `json:"medium"`
			Original string `json:"original"`
			Small    string `json:"small"`
			Tiny     string `json:"tiny"`
		} `json:"posterImage"`
		RatingFrequencies map[string]string `json:"ratingFrequencies"`
		RatingRank        int64             `json:"ratingRank"`
		ShowType          string            `json:"showType"`
		Slug              string            `json:"slug"`
		StartDate         string            `json:"startDate"`
		Status            string            `json:"status"`
		Subtype           string            `json:"subtype"`
		Synopsis          string            `json:"synopsis"`
		Tba               interface{}       `json:"tba"`
		Titles            struct {
			En   string `json:"en"`
			EnJp string `json:"en_jp"`
			JaJp string `json:"ja_jp"`
		} `json:"titles"`
		UpdatedAt      string `json:"updatedAt"`
		UserCount      int64  `json:"userCount"`
		YoutubeVideoID string `json:"youtubeVideoId"`
	} `json:"attributes"`
	ID string `json:"id"`
}

func (a Anime) GetName() string {
	return "a" + a.ID
}

func (a Anime) GetData() interface{} {
	return a.Attributes
}

type User struct {
	Attributes struct {
		About               string        `json:"about"`
		AboutFormatted      interface{}   `json:"aboutFormatted"`
		Avatar              interface{}   `json:"avatar"`
		Bio                 string        `json:"bio"`
		Birthday            interface{}   `json:"birthday"`
		CommentsCount       int64         `json:"commentsCount"`
		CoverImage          interface{}   `json:"coverImage"`
		CreatedAt           string        `json:"createdAt"`
		FacebookID          string        `json:"facebookId"`
		FavoritesCount      int64         `json:"favoritesCount"`
		FeedCompleted       bool          `json:"feedCompleted"`
		FollowersCount      int64         `json:"followersCount"`
		FollowingCount      int64         `json:"followingCount"`
		Gender              string        `json:"gender"`
		LifeSpentOnAnime    int64         `json:"lifeSpentOnAnime"`
		LikesGivenCount     int64         `json:"likesGivenCount"`
		LikesReceivedCount  int64         `json:"likesReceivedCount"`
		Location            interface{}   `json:"location"`
		MediaReactionsCount int64         `json:"mediaReactionsCount"`
		Name                string        `json:"name"`
		PastNames           []interface{} `json:"pastNames"`
		PostsCount          int64         `json:"postsCount"`
		ProExpiresAt        interface{}   `json:"proExpiresAt"`
		ProfileCompleted    bool          `json:"profileCompleted"`
		RatingSystem        string        `json:"ratingSystem"`
		RatingsCount        int64         `json:"ratingsCount"`
		ReviewsCount        int64         `json:"reviewsCount"`
		Theme               string        `json:"theme"`
		Title               interface{}   `json:"title"`
		UpdatedAt           string        `json:"updatedAt"`
		WaifuOrHusbando     interface{}   `json:"waifuOrHusbando"`
		Website             interface{}   `json:"website"`
	} `json:"attributes"`
	ID    string `json:"id"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Relationships struct {
		Blocks struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"blocks"`
		Favorites struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"favorites"`
		Followers struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"followers"`
		Following struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"following"`
		LibraryEntries struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"libraryEntries"`
		LinkedAccounts struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"linkedAccounts"`
		NotificationSettings struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"notificationSettings"`
		OneSignalPlayers struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"oneSignalPlayers"`
		PinnedPost struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"pinnedPost"`
		ProfileLinks struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"profileLinks"`
		Reviews struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"reviews"`
		Stats struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"stats"`
		UserRoles struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"userRoles"`
		Waifu struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"waifu"`
	} `json:"relationships"`
	Type string `json:"type"`
}

type UserFilterRequest struct {
	Data  []User `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
	Meta struct {
		Count int64 `json:"count"`
	} `json:"meta"`
}

type UserRequest struct {
	Data  User `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
}
