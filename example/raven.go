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
	listen        bool
	help          bool
	restartMsg    = "Restart XML engine"
	factoryMsg    = "Factory Reset"
	listenMsg     = "Listen"
	showMsg       = "Show this message"
	helpMsg       = "raven " + "(v" + VERSION + ") " +
		"send commands to your raven device\n" +
		"Usage: [OPTIONS] DEVICE \n" +
		"\t-r " + restartMsg + " \n" +
		"\t-f " + factoryMsg + " \n" +
		"\t-l " + listenMsg + " \n" +
		"\t-h " + showMsg + " \n"
)

func init() {
	flag.BoolVar(&restart, "r", false, restartMsg)
	flag.BoolVar(&factory_reset, "f", false, factoryMsg)
	flag.BoolVar(&listen, "l", false, listenMsg)
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
	defer r.Close()

	// Factory Reset
	if factory_reset {
		err = r.FactoryReset()
		if err != nil {
			log.Fatal(err)
		}
	}

	// Restart
	if restart {
		err = r.Restart()
		if err != nil {
			log.Fatal(err)
		}
	}

	// Listen
	if listen {
		for {
			notify, err := r.Receive()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(notify)
		}
	}
}
