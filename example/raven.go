package main

import (
	"flag"
	"fmt"
	"github.com/dahenson/goraven"
	"log"
	"os"
)

const VERSION = "0.1"

var (
	restart       bool
	factory_reset bool
	help          bool
	restartMsg    = "Restart XML engine"
	factoryMsg    = "Factory Reset"
	showMsg       = "Show this message"
	helpMsg       = "raven " + "(v" + VERSION + ") " +
		"send commands to your raven device\n" +
		"Usage: [OPTIONS] DEVICE \n" +
		"\t-r " + restartMsg + " \n" +
		"\t-f " + factoryMsg + " \n" +
		"\t-h " + showMsg + " \n"
)

func init() {
	flag.BoolVar(&restart, "r", false, restartMsg)
	flag.BoolVar(&factory_reset, "f", false, factoryMsg)
	flag.BoolVar(&help, "h", false, showMsg)
}

var usage = func() {
	fmt.Print(helpMsg)
	os.Exit(0)
}

func parseFlags() {
	flag.Usage = usage
	flag.Parse()
	n := flag.NArg()
	if len(os.Args) == 1 || help || n == 0 {
		flag.Usage()
	}
}

func main() {
	parseFlags()

	dev := flag.Arg(0)
	r, err := goraven.Connect(dev)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Disconnect()

	if factory_reset {
		err = r.FactoryReset()
		if err != nil {
			log.Fatal(err)
		}
	}

	if restart {
		err = r.Restart()
		if err != nil {
			log.Fatal(err)
		}
	}
}
