package main

import (
	"fmt"
	"math/rand"
)

func sum(s []float64, c chan float64) {
	sum := 0.0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	NUM := 1000
	s := []float64{}
	a := []float64{}
	s1 := rand.NewSource(42)
	r1 := rand.New(s1)

	c := make(chan float64)
	for i := 1; i < NUM; i++ {
		s = append(s, r1.Float64())
		if (i % (NUM/10)) == 0 {
			fmt.Println(i)
			go sum(s[i-NUM/10:i], c)
		}
	}
	for i := 1; i < (NUM/100) ; i++ {
		a = append(a, <-c/10)
	}
	fmt.Println(a)
}
