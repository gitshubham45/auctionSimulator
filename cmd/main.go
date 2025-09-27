package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/gitshubham45/auctionSimulator/internal/auctionPkg"
	"github.com/gitshubham45/auctionSimulator/internal/utils"
	"github.com/joho/godotenv"
)

const (
	NumAuctions          = 40  // number of auctions to run concurrently
	NumBiddersPerAuction = 100 // number of bidders participating in each auction
	NumAttributes        = 20  // number of attributes per auction
)

func main() {

	fmt.Println("-------Welcome to Auction Simulator-------")
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	mem := humanizeBytes(m.Sys)
	fmt.Printf("Initial memory usage: %v bytes\n", mem)

	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	vcpuEnv := os.Getenv("SIM_VCPU")
	vcpu := runtime.NumCPU()
	if v, err := strconv.Atoi(vcpuEnv); err == nil && v > 0 {
		vcpu = v
	}

	runtime.GOMAXPROCS(vcpu)
	fmt.Printf("Using GOMAXPROCS=%d (vCPU units). NumCPU reported: %d\n", runtime.GOMAXPROCS(0), runtime.NumCPU())

	// we can set this to runtime.GOMAXPROCS(0) * X where X is factor of goroutines per CPU.
	SemaphoreLimitFactor := os.Getenv("SEMPAPHORE_LIMIT_FACTOR")
	SemaphoreLimitFactorValue := 4
	if v, err := strconv.Atoi(SemaphoreLimitFactor); err == nil && v > 0 {
		SemaphoreLimitFactorValue = v
	}
	semCap := runtime.GOMAXPROCS(0) * SemaphoreLimitFactorValue
	if semCap < 1 {
		semCap = 1
	}

	sem := utils.NewSemaphore(semCap)
	fmt.Printf("Semaphore concurrency limit = %d\n", semCap)

	bidders := make([]auctionPkg.Bidder, NumBiddersPerAuction)

	for i := 0; i < NumBiddersPerAuction; i++ {
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
		sem.Acquire()

		go func(auc *auctionPkg.Auction) {
			defer wg.Done()
			defer sem.Release()
			// Run Auctions
			auctionPkg.RunAuction(context.Background(), auc, bidders)
			winnerJSON, _ := json.Marshal(auc.Winner)
			fmt.Printf("Auction %d Completed: Winner %s\n", auc.ID, winnerJSON)
			fmt.Println("Total Bids", len(auc.Bids))
			fmt.Println("-----------------------------------------------------------------")
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

	runtime.ReadMemStats(&m)
	mem = humanizeBytes(m.Sys)
	fmt.Printf("Final memory usage: %v bytes\n", mem)
}

func humanizeBytes(s uint64) string {
	if s < 1024 {
		return fmt.Sprintf("%d B", s)
	}
	exp := int(math.Log(float64(s)) / math.Log(1024))
	suffix := "KMGTPE"[exp-1]
	val := float64(s) / math.Pow(1024, float64(exp))
	return fmt.Sprintf("%.2f %cB", val, suffix)
}
