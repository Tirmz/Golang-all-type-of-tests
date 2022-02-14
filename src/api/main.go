package main

import (
	"Golang-all-type-of-tests/src/api/providers/locations_provider"
	"fmt"
)

func main() {
	country, err := locations_provider.GetCountry("AR")

	fmt.Println(err)
	fmt.Println(country)
}
