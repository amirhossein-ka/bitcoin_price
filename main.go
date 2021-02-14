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
	var currency string
	var price string
	fmt.Print("Enter the currency(UPPERCASE):\t")
	_, err := fmt.Scanln(&currency)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Print("Enter the output:\t ")
	_, err = fmt.Scanln(&price)
	if err != nil {
		log.Fatal(err)
	}


	urls := fmt.Sprintf("https://min-api.cryptocompare.com/data/price?fsym=%s&tsyms=%s", currency, price)
	resp, err := http.Get(urls)
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
	//_ = fmt.Sprint(i1)
	//_ = fmt.Sprint(j2)

	// remove "" from USD
	s = strings.Split(i1, "\"")
	typ = s[1]
	fmt.Println(typ)
	// remove } from numbers

	s = strings.Split(j2, "}")
	value = s[0]
	fmt.Println(value)

}
