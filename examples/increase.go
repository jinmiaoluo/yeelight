package main

import (
	"log"

	"github.com/avarabyeu/yeelight"
)

func main() {
	y, err := yeelight.Discover()
	checkError(err)

	y.IncreaseBright()
}

func checkError(err error) {
	if nil != err {
		log.Fatal(err)
	}
}
