package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gitshubham45/auctionSimulator/internal/auctionPkg"
)

func WriteAuctionOutput(a *auctionPkg.Auction, outputDir string) error {
	filename := fmt.Sprintf("auction_%03d.json", a.ID)
	filepath := filepath.Join(outputDir, filename)

	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", " ")
	return enc.Encode(a)
}
