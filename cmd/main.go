package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gitshubham45/auctionSimulator/internal/auction"
	"github.com/gitshubham45/auctionSimulator/internal/bidder"
)

func main() {
	fmt.Println("Welcome to Auction Simulator")

	bidders := make([]*bidder.Bidder, 100)

	for i := 0; i < 100; i++ {
		bidders[i] = &bidder.Bidder{
			ID: i + 1,
		}
	}

	auctions := make([]*auction.Auction, 40)

	for i := 0; i < 40; i++ {
		attr := make(auction.Attribute)
		for j := 0; j < 20; j++ {
			attr[fmt.Sprintf("attr_%d", j+1)] = fmt.Sprintf("value_%d", j+1)
		}
		auctions[i] = &auction.Auction{
			ID:         i + 1,
			Attributes: attr,
			TimeoutSec: 5,
		}
	}

	var wg sync.WaitGroup
	start := time.Now()

	for _, a := range auctions {
		wg.Add(1)

		go func(auc *auction.Auction) {
			defer wg.Done()
			// Run Auctions
			fmt.Printf("Auction %d Completed: Winner %v, Duration %d md \n", auc.ID, auc.Winner, auc.DurationMs)
		}(a)
	}

	wg.Wait()
	fmt.Printf("Total time taken: %v\n", time.Since(start))
}
