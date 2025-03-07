package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"log/slog"
)

// AppName : App name
const AppName = "epgo"

// Version : Version
const Version = "1.1.3"

// Config : Config file (struct)
var Config config
var Config2 string
var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

func main() {
	log.SetOutput(os.Stdout)
	var configure = flag.String("configure", "", "= Create or modify the configuration file. [filename.yaml]")
	var config = flag.String("config", "", "= Get data from Schedules Direct with configuration file. [filename.yaml]")

	var h = flag.Bool("h", false, ": Show help")

	flag.Parse()
	Config2 = *config
	logger.Info("epgo revamped", "Version", Version, "Forked", "By Chuchodavids")

	if *h {
		fmt.Println()
		flag.Usage()
		os.Exit(0)
	}

	if len(*configure) != 0 {
		err := Configure(*configure)
		if err != nil {
			logger.Error("could not open configuration", "error", err)
		}
		os.Exit(0)
	}

	if len(*config) != 0 {
		var sd SD
		err := sd.Update(*config)
		if err != nil {
			ShowErr(err)
		}
	}
}

// ShowErr : Show error on screen
func ShowErr(err error) {
	var msg = fmt.Sprintf("%s", err)
	logger.Error(msg)

}
