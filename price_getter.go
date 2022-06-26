package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ResponseBody struct {
	Bitcoin Bitcoin
}
type Bitcoin struct {
	Idr int64
}
type Price_getter struct {
	coinId       string
	currencyType string
}

func (p *Price_getter) getPrice() int64 {
	url := "https://api.coingecko.com/api/v3/simple/price?ids="
	url += p.coinId
	url += "&vs_currencies="
	url += p.currencyType
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var urlResonse ResponseBody
	koh := json.Unmarshal(body, &urlResonse)
	if koh != nil {
		panic(koh)
	}
	return urlResonse.Bitcoin.Idr
}
