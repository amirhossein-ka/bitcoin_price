package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	var typ string
	var value string
	//urls = "https://min-api.cryptocompare.com/data/price?fsym=BTC&tsyms=USD"
	resp, err := http.Get("https://min-api.cryptocompare.com/data/price?fsym=BTC&tsyms=USD")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	out := string(body)
	//fmt.Println(reflect.TypeOf(out))
	// split USD and values from each other
	s := strings.Split(out, ":")
	i1, j2 := s[0], s[1]
	fmt.Sprint(i1)
	fmt.Sprint(j2)
	// remove "" from USD
	s = strings.Split(out, "\"")
	typ = s[1]
	fmt.Println(typ)
	// remove } from numbers
	s = strings.Split(j2, "}")
	value = s[0]
	fmt.Println(value)

}
