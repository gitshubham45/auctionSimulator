package auction

import "time"

type Attribute map[string]string

type Bid struct {
	BidderId   int       `json:"bidder_id"`
	Amount     float64   `json:"amount"`
	ReceivedAt time.Time `json:"received_at"`
}

type Auction struct {
	ID         int       `json:"id"`
	Attributes Attribute `json:"attributes"`
	TimeoutSec int       `json:"timeout_sec"` // limit time for auction
	StartTs    time.Time `json:"start_ts"`
	EndTs      time.Time `json:"end_ts"`
	Bids       []Bid     `json:"bids"`
	Winner     *Bid      `josn:"winner"`
	DurationMs int64     `json:"duration_ms"` // duration of auction
}
