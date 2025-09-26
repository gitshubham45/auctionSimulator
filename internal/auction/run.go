package auction

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func RunAuction(ctx context.Context, auction *Auction, bidders []Bidder) {
	auction.StartTs = time.Now()

	// channel to collect bids
	bidsCh := make(chan Bid, len(bidders))
	fmt.Println(bidsCh)

	var wg sync.WaitGroup
	for _, b := range bidders {
		wg.Add(1)
		go func(b Bidder) {
			// place bid


		}(b)
	}

}
