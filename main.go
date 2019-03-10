package main

import (
	"github.com/brian-armstrong/gpio"
	"log"
	"time"
)

func main() {

	for {
		var i uint = 0
		for ; i < 26; i++ {
			pin := gpio.NewOutput(i, false)
			if err := pin.High(); err != nil {
				log.Printf("Can't set high value to %d pin: %v", pin.Number, err)
			}
			time.Sleep(time.Second / 2)
			if err := pin.Low(); err != nil {
				log.Printf("Can't set low value to %d pin: %v", pin.Number, err)
			}
		}
		time.Sleep(time.Second / 2)
	}
}
