package schema

type Series struct {
	Country      string
	Description  string
	Drm          bool
	EpisodeCount int
	Genre        string
	Image        struct {
		ShowImage string `json:"showImage"`
	} `json:"image"`
	Language    string
	NextEpisode struct {
		Channel     *string
		ChannelLogo string
		Date        *string
		Html        string
		Url         string
	}
	PrimaryColour string
	Seasons       []struct {
		Slug string
	}
	Slug      string `json:"slug"`
	Title     string `json:"title"`
	TvChannel string
}
type Request struct {
	Payload      []Series `json:"payload"`
	skip         int
	take         int
	totalRecords int
}
