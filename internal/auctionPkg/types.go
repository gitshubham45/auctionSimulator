package auctionPkg

import (
	"math/rand/v2"
	"time"
)

type Attribute map[string]string

type Bid struct {
	BidderId int       `json:"bidder_id"`
	Amount   float64   `json:"amount"`
	PlacedAt time.Time `json:"placed_at"`
}

type Auction struct {
	ID         int       `json:"id"`
	Attributes Attribute `json:"attributes"`
	BaseValue  float64   `json:"base_value"`
	TimeoutSec int       `json:"timeout_sec"` // limit time for auction
	StartTs    time.Time `json:"start_ts"`
	EndTs      time.Time `json:"end_ts"`
	Bids       []Bid     `json:"bids"`
	Winner     *Bid      `josn:"winner"`
	DurationMs int64     `json:"duration_ms"` // duration of auction
}

type Bidder struct {
	ID int
}

func (b *Bidder) PlaceBid(attributes Attribute) (Bid, bool) {
	// decide if want to bid or not
	// for now passinf 70% of bids
	if rand.Float64() < 0.7 {
		// decide based on baseValue a
		bidAmount := rand.Float64() * 100 // random bid amount
		return Bid{
			BidderId: b.ID,
			Amount:   bidAmount,
			PlacedAt: time.Now(),
		}, true
	}
	return Bid{}, false
}
