package main

import (
	"fmt"
)

func main() {
	var stokcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stokcode, enddate)
	fmt.Printf(url+"\n", stokcode, enddate)
	fmt.Println(target_url)
}
