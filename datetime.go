package main

/*
As per the documentation

Operating systems provide both a “wall clock,” which is subject to changes for clock synchronization, and a “monotonic clock,” which is not.

The general rule is that the wall clock is for telling time and the monotonic clock is for measuring time

Time returned by time.Now contains both a wall clock reading and a monotonic clock reading;
later time-telling operations use the wall clock reading,
but later time-measuring operations, specifically comparisons and subtractions, use the monotonic clock reading.
*/

import "time"
import "fmt"

func main() {
	start := time.Now()
	fmt.Println(start)

	end := time.Now()
	fmt.Println(end)

	elapsed := end.Sub(start)
	fmt.Println(elapsed)
}
