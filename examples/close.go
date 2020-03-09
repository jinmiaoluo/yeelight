package main

import (
	"log"

	"github.com/avarabyeu/yeelight"
)

func main() {
	y, err := yeelight.Discover()
	checkError(err)

	y.SetPower(false)
}

func checkError(err error) {
	if nil != err {
		log.Fatal(err)
	}
}
