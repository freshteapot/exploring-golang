package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	_, week := t.ISOWeek()
	fmt.Printf("Todays week is %d\n", week)
}
