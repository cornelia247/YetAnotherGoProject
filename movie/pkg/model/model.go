package model

import model "movieexample.com/metadata/pkg"

// MovieDetails includes movie metadata and it's aggregated rating

type MovieDetails struct {
	Rating   *float64       `json:"rating,omitempty"`
	Metadata model.Metadata `json:"metadata"`
}
