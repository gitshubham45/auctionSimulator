package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/gitshubham45/auctionSimulator/internal/auctionPkg"
)

func main() {
	fmt.Println("Welcome to Auction Simulator")

	bidders := make([]auctionPkg.Bidder, 100)

	for i := 0; i < 100; i++ {
		bidders[i] = auctionPkg.Bidder{
			ID: i + 1,
		}
	}

	auctions := make([]*auctionPkg.Auction, 40)

	for i := 0; i < 40; i++ {
		attr := make(auctionPkg.Attribute)
		for j := 0; j < 20; j++ {
			attr[fmt.Sprintf("attr_%d", j+1)] = fmt.Sprintf("value_%d", j+1)
		}
		auctions[i] = &auctionPkg.Auction{
			ID:         i + 1,
			Attributes: attr,
			TimeoutSec: 5,
		}
	}

	var wg sync.WaitGroup
	start := time.Now()

	for _, auction := range auctions {
		wg.Add(1)

		go func(auc *auctionPkg.Auction) {
			defer wg.Done()
			// Run Auctions
			auctionPkg.RunAuction(context.Background(), auc, bidders)
			fmt.Printf("Auction %d Completed: Winner %v, Duration %d md \n", auc.ID, auc.Winner, auc.DurationMs)
		}(auction)
	}

	wg.Wait()
	fmt.Printf("Total time taken: %v\n", time.Since(start))
}
