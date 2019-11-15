//from https://tour.golang.org/welcome/1
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Packages: ")

	fmt.Println("time > The time is", time.Now())

	fmt.Printf("math > Now you have %g problems.\n", math.Sqrt(7))

	fmt.Println("math/rand > My favorite number is", rand.Intn(10))
}
