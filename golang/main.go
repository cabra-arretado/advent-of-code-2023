package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	number := flag.String("n", "", "number")
	if err := flag.Parse(); err != nil { // Use short declaration inside the if statement
		fmt.Println("Error parsing flags:", err)
		os.Exit(1)
	}

	fmt.Printf("Hello, %s\n", *number)
}
