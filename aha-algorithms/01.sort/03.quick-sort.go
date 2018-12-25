package main

import (
	"fmt"
	"math/rand"
	"time"
)


var d = make([]int, 0)

func quickSort(left int, right int) {

	base := d[left]
	i := left
	j := right

	if left >= right {
		return
	}

	for left < right  {
		//if left === right {
		//	return
		//}


		for d[right] >= base && left < right {
			right--
		}

		for d[left] <= base && left < right {
			left++
		}

		if left < right {
			d[left], d[right] = d[right], d[left]
		}

	}


	quickSort(left, i-1)
	quickSort(i+1, right)
}

func main()  {
	newSeeder := rand.NewSource(time.Now().UnixNano())
	newRand := rand.New(newSeeder)

	//d := make([]int, 0)
	for i := 0; i < 20; i++ {
		d = append(d, newRand.Intn(101))
	}

	quickSort(0, 20)

	fmt.Println(d)
}
