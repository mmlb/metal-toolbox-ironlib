package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/bombsimon/logrusr/v4"
	"github.com/metal-toolbox/ironlib"
	"github.com/sirupsen/logrus"
)

// This example invokes ironlib and prints out the BIOS features on supported platforms
// a sample output can be seen in the biosconfig.json file

func main() {
	l := logrus.New()
	l.Formatter = &logrus.JSONFormatter{}
	l.Level = logrus.TraceLevel
	logger := logrusr.New(l)

	device, err := ironlib.New(logger)
	if err != nil {
		logger.Error(err, "creating ironlib manager")
		os.Exit(1)
	}

	features, err := device.GetBIOSConfiguration(context.TODO())
	if err != nil {
		logger.Error(err, "getting bios config")
		os.Exit(1)
	}

	j, err := json.MarshalIndent(features, " ", "  ")
	if err != nil {
		logger.Error(err, "formatting json")
		os.Exit(1)
	}

	fmt.Println(string(j))
}
