package main

import (
	"beertown/mathematics"
	"fmt"
)

func main() {

	fmt.Println("Enter the number to Find the Best beers")
	var n int
	fmt.Scanf("%d", &n)

	goodBeers := mathematics.FindBeers(n)

	if goodBeers == nil {
		fmt.Println("No Good beer found :(")
	} else {
		fmt.Printf("Good beer numbers are : %+v\n", goodBeers)
	}

}
