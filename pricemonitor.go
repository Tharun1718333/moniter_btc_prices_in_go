package main

import "fmt"

type PriceMonitor struct {
	Price_getter
	Buyer
	initialPrice int64
	currentPrice int64
}

func (p *PriceMonitor) getHowManyToBuy() int64 {
	//use some logic here
	magic_number := p.currentPrice + p.initialPrice
	fmt.Println(magic_number)
	return 0
}
func (p *PriceMonitor) initilize() {
	p.initialPrice = p.Price_getter.getPrice()
}
func (p *PriceMonitor) update() {
	p.currentPrice = p.Price_getter.getPrice()
}
func (p *PriceMonitor) buy(count int64) string {
	//can do buying in phases
	p.Buyer.buy(count)
	return TempResponse
}
