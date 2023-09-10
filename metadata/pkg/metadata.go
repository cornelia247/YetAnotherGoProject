package model

// Metadata defines the movie metadata.

type Metadata struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"descripton"`
	Director    string `json:"director"`
}
