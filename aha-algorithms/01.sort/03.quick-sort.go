package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(d *[]int, left int, right int) {
	//fmt.Println(d, left, right)


	base := (*d)[left]

	for left < right  {
		//if left === right {
		//	return
		//}


		for right  {

			right--
		}




	}

	//return d
}

func main()  {
	newSeeder := rand.NewSource(time.Now().UnixNano())
	newRand := rand.New(newSeeder)

	d := make([]int, 0)
	for i := 0; i < 20; i++ {
		d = append(d, newRand.Intn(101))
	}

	fmt.Println(quickSort(&d))
}
