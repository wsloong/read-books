package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Sleepping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
