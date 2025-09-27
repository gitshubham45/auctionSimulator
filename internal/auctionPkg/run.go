package auctionPkg

import (
	"context"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

func RunAuction(ctx context.Context, auction *Auction, bidders []Bidder) {
	auction.StartTs = time.Now()

	delayFactorStr := os.Getenv("DELAY_FACTOR")
	delayFactorValue := 100

	if delayFactorStr != "" {
		if val, err := strconv.Atoi(delayFactorStr); err == nil && val > 0 {
			delayFactorValue = val
		} else {
			log.Printf("Invalid or negative DELAY_FACTOR '%s', using default %d", delayFactorStr, delayFactorValue)
		}
	}

	// channel to collect bids
	bidsCh := make(chan Bid, len(bidders))

	var wg sync.WaitGroup
	for _, b := range bidders {
		wg.Add(1)
		go func(b Bidder) {
			// place bid
			defer wg.Done()
			bid, ok := b.PlaceBid(auction.Attributes, auction.BaseValue, auction.TimeoutSec, delayFactorValue)
			if ok {
				bidsCh <- bid
			}
		}(b)
	}

	// closing bids channel after all bids collected
	go func() {
		wg.Wait()
		close(bidsCh)
	}()

	// auction timeout context
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Duration(auction.TimeoutSec)*time.Second)
	defer cancel()

	// checking timout of bids
	for {
		select {
		case bid, ok := <-bidsCh:
			if !ok {
				// if no more bids (within time limit)
				auction.EndTs = time.Now()
				goto end
			}
			auction.Bids = append(auction.Bids, bid)
		case <-timeoutCtx.Done():
			auction.EndTs = time.Now()
			goto end
		}
	}

end:

	// find winner with highes bid
	var winner *Bid
	var maxAmount float64 = -1
	for _, bid := range auction.Bids {
		if bid.Amount > maxAmount {
			maxAmount = bid.Amount
			winner = &bid

		}
	}

	auction.Winner = winner
	auction.DurationMs = auction.EndTs.Sub(auction.StartTs).Milliseconds()
}
