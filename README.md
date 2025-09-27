# 🎯 Auction Simulator

A concurrent auction simulation system that runs 40 auctions in parallel with 100 bidders each, measuring performance under controlled resource constraints.

## 📌 Objective

The goal of this project is to design and implement an **Auction Simulator** that:

- Runs **40 auctions concurrently**
- Simulates **100 bidders per auction**, each making decisions based on item attributes
- Enforces **per-auction timeouts**
- Collects valid bids and declares winners
- Measures **end-to-end execution time** (from first auction start to last completion)
- Standardizes resource usage (vCPU, RAM) for consistent and reproducible behavior

This project satisfies the requirements of a scalable, deterministic simulation suitable for performance benchmarking and concurrency testing.

---

## 🔧 Features

✅ **Concurrent Auctions**  
- 40 auctions run simultaneously using Go goroutines.

✅ **Smart Bidders**  
- 100 unique bidders participate across all auctions.
- Each bidder evaluates 20-item attributes (simulated) and decides:
  - Whether to place a bid (~70% probability)
  - With a random reaction delay (realistic network + decision time)

✅ **Timeout-Based Auctions**  
- Each auction runs for a defined duration (`TimeoutSec`)
- Late bids are rejected after timeout
- Winner = highest bid among timely submissions

✅ **Performance Measurement**  
- Total runtime measured from **start of first auction** to **end of last auction**
- Output includes:
  - Per-auction winner (in JSON)
  - Duration
  - Bid statistics

✅ **Resource Standardization**  
- Controls vCPU usage via `GOMAXPROCS`
- Optional environment override: `SIM_VCPU=4`
- Memory limits can be added via `debug.SetMemoryLimit()` (future extension)

✅ **Clean Output & Logging**  
- Each auction result saved as `output/auction_XXX.json`
- Console logs show progress and final stats

---

## 📦 Requirements

- Go 1.20 or higher
- macOS/Linux/Windows
- No external dependencies

---

## ⚙️ How to Run

### 1. Clone the Repository

```bash
git clone https://github.com/gitshubham45/auctionSimulator.git
cd auctionSimulator
```

### 2.  Create .env File
```bash
cp .env.example .env
```

You can edit .env to customize settings:
```bash
SIM_VCPU=5
DELAY_FACTOR=400
SEMPAPHORE_LIMIT_FACTOR=6
```
### 3. Install Dependencies
```bash
go mod tidy
```
### 4. Run 
```bash
cd cmd
go run main.go
```