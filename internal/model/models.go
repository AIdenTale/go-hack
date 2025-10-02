package model

type BPM struct {
	Time string  `json:"time"`
	BPM  float64 `json:"bpm"`
}

type Trac struct {
	Time string  `json:"time"`
	Trac float64 `json:"trac"`
}

type PregnantData struct {
	ID           int64    `json:"id"`
	TimeRecorded string   `json:"time_recorded"`
	BPM          *float64 `json:"bpm,omitempty"`
	Trac         *float64 `json:"trac,omitempty"`
}
