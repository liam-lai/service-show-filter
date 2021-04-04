package schema

type FilteredDrm struct {
	Image string `json:"image"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
}
type Response struct {
	Response []*FilteredDrm `json:"response,omitempty"`
}
