// Service that gets information every 10 seconds
package main

import (
	"fmt"
	"time"
)

func main() {
	for {

		fmt.Printf("Current time: %v\n", time.Now().Unix())
		time.Sleep(10 * time.Second)
	}
}
