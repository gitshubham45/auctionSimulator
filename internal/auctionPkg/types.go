package auctionPkg

import (
	"math/rand/v2"
	"time"
)

type Attribute map[string]float64

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

func (b *Bidder) PlaceBid(attributes Attribute, baseValue float64, timeoutSec, delayFactor int) (Bid, bool) {
	min := 10 * time.Millisecond
	max := time.Duration(timeoutSec) * time.Second
	delay := min + time.Duration(rand.Float64()*float64(max-min)) + time.Duration(delayFactor)*time.Millisecond

	// Actually wait (simulate network + reaction time)
	time.Sleep(delay)
	// decide based on baseValue of auction
	bidAmount := baseValue + rand.Float64()*100 // random bid amount
	return Bid{
		BidderId: b.ID,
		Amount:   bidAmount,
		PlacedAt: time.Now(),
	}, true
}
