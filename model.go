package restapisample

//Job contains the business model data.
type Job struct {
	Company     string `json:"company"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}
