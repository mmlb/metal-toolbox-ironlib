package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"github.com/metal-toolbox/ironlib"
)

// This example invokes ironlib and prints out the BIOS features on supported platforms
// a sample output can be seen in the biosconfig.json file

func main() {
	logger := slog.Default()

	device, err := ironlib.New(logger)
	if err != nil {
		logger.Error("creating ironlib manager", "error", err)
		os.Exit(1)
	}

	features, err := device.GetBIOSConfiguration(context.TODO())
	if err != nil {
		logger.Error("getting bios config", "error", err)
		os.Exit(1)
	}

	j, err := json.MarshalIndent(features, " ", "  ")
	if err != nil {
		logger.Error("formatting json", "error", err)
		os.Exit(1)
	}

	fmt.Println(string(j))
}
