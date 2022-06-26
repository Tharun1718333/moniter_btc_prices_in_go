package main

import (
	"time"
)

var SecondsTowait int = 5

func main() {
	price_getter := Price_getter{"bitcoin", "idr"}
	buyer := Buyer{18}
	price_monitor := PriceMonitor{price_getter, buyer, 0, 0}
	price_monitor.initilize()
	for {
		price_monitor.update()
		k := price_monitor.getHowManyToBuy()
		if k != 0 {
			price_monitor.buy(k)
		}
		time.Sleep(time.Duration(SecondsTowait) * time.Second)
	}
}
