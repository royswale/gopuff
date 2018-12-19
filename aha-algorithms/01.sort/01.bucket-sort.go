package main

import (
	"fmt"
	"math/rand"
)

// max number, i.e. the bucket size
const T  = 1000
//const T  = 1000000000

// bucket sort
func bucketSort(a *[]int) *[]int {

	b := make([]int, T)

	for _, v := range *a {
		b[v]++
	}

	// ascending
	//for i := 0; i < T; i++ {
	//	for j := 0; j < b[i]; j++ {
	//		fmt.Println(i)
	//	}
	//}

	// descending
	for i := T - 1; i >= 0; i-- {
		for j := 0; j < b[i]; j++ {
			fmt.Println(i)
		}
	}
	//fmt.Println(b)

	return a
}

func main()  {

	a := make([]int, 0)

	// generate 10 pseudo-random number [0, 100)
	for i := 0; i < 10; i++ {
		a = append(a, rand.Intn(T))
		//fmt.Println(len(a), cap(a))
	}


	fmt.Println(bucketSort(&a))

}