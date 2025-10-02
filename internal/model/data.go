package model

import "time"

type DataTimeRange struct {
	Seconds int64 `json:"seconds" query:"seconds"`
}

type DataPoint struct {
	Time time.Time `json:"time"`
	D    float64   `json:"d"`
}

type DataResponse struct {
	FHR       []DataPoint `json:"fhr,omitempty"`
	UC        []DataPoint `json:"uc,omitempty"`
	LastChunk bool        `json:"last_chunk"`
}