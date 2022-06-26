package main

import "fmt"

var TempResponse string = "OK"

type Buyer struct {
	balance int64
}

func (b *Buyer) buy(count int64) string {
	//some logic for buying
	// return global constants accordingly
	k := b.balance
	fmt.Println("buying with balance ", k)
	return TempResponse
}
