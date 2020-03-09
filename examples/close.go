package main

import (
	"github.com/avarabyeu/yeelight"
	"log"
)

func main() {
	y, err := yeelight.Discover()
	checkError(err)

	y.SetPower(true)
}

func checkError(err error) {
	if nil != err {
		log.Fatal(err)
	}
}
