package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/avarabyeu/yeelight"
)

func main() {
	flag.Parse()
	actions := flag.Args()
	y, err := yeelight.Discover()
	if nil != err {
		log.Fatal(err)
	}
	if len(actions) == 1 {
		action := actions[0]
		switch action {
		case "open":
			y.SetPower(true)
		case "close":
			y.SetPower(false)
		case "up", "increase":
			y.IncreaseBright()
		case "down", "decrease":
			y.DecreaseBright()
		default:
			fmt.Printf("Try: %s <open|close|up|down|increase|decrease>\n", os.Args[0])
		}
	} else {
		y.SetPower(true)
	}
}
