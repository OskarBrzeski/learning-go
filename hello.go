package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	bytes_count, err := fmt.Println(quote.Go())
	fmt.Println(bytes_count)
	fmt.Println(err)
}
