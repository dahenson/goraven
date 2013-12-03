package main

import (
	"bitbucket.org/LocalHero/libraven"
	"flag"
	"fmt"
	"log"
	"os"
)

const VERSION = "0.1"

var (
	restart    bool
	help       bool
	restartMsg = "Restart XML engine"
	factoryMsg = "Factory Reset"
	showMsg    = "Show this message"
	helpMsg    = "raven " + "(v" + VERSION + ") " +
		"send commands to your raven device\n" +
		"Usage: [-r|-f] device ... \n" +
		"\t-r " + restartMsg + " \n" +
		"\t-f " + factoryMsg + " \n" +
		"\t-h " + showMsg + " \n"
)

func init() {
	flag.BoolVar(&restart, "r", false, restartMsg)
	flag.BoolVar(&help, "h", false, showMsg)
}

var usage = func() {
	fmt.Print(helpMsg)
	os.Exit(0)
}

func parseFlags() {
	flag.Usage = usage
	flag.Parse()
	if len(os.Args) == 1 || help {
		flag.Usage()
	}
}

func main() {
	parseFlags()

	r, err := libraven.Connect("/dev/ttyUSB0")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Disconnect()

	if restart {
		err = r.Restart()
		if err != nil {
			log.Fatal(err)
		}
	}
}
