package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/gitshubham45/auctionSimulator/internal/auctionPkg"
	"github.com/gitshubham45/auctionSimulator/internal/utils"
)

const (
	NumAuctions          = 40  // number of auctions to run concurrently
	NumBiddersPerAuction = 100 // number of bidders participating in each auction
	NumAttributes        = 20  // number of attributes per auction
)

func main() {
	fmt.Println("Welcome to Auction Simulator")

	bidders := make([]auctionPkg.Bidder, NumBiddersPerAuction)

	for i := 0; i < 100; i++ {
		bidders[i] = auctionPkg.Bidder{
			ID: i + 1,
		}
	}

	auctions := make([]*auctionPkg.Auction, 40)

	for i := 0; i < NumAuctions; i++ {
		attr := make(auctionPkg.Attribute)
		var sum float64
		for j := 0; j < NumAttributes; j++ {
			val := rand.Float64() * 100
			attr[fmt.Sprintf("attr_%d", j+1)] = val
			sum += val
		}
		avg := sum / 20

		auctions[i] = &auctionPkg.Auction{
			ID:         i + 1,
			Attributes: attr,
			TimeoutSec: int(5 + avg/20), //  base timeout 5 + scaled by avg attribute
			BaseValue:  avg,             // base value is average attribute
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

	const outputDir = "../output"
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		fmt.Println("failed to create ouptput directory")
	}

	for _, auction := range auctions {
		if err := utils.WriteAuctionOutput(auction, outputDir); err != nil {
			fmt.Println("Error wrting to output directory")
		}
	}
}
