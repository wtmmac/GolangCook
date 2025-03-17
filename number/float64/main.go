package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	v := struct {
		Balance string
	}{
		Balance: "123.45",
	}

	balance, err := strconv.ParseFloat(v.Balance, 2)
	if err != nil {
		log.Fatalf("Error parsing float from string '%s': %v", v.Balance, err)
	}

	fmt.Printf("Parsed float: %f\n", balance)
}
