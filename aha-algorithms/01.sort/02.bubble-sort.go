package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(d *[]int) *[]int {
	fmt.Println(d)

	l := len(*d)
	fmt.Println(l)

	for i := 0; i < l-1; i++ {
		fmt.Println("round ", i)
		//fmt.Println((*d)[i])
		for j := 0; j < l-1-i; j++ {
			if (*d)[j] < (*d)[j+1] {
				(*d)[j], (*d)[j+1] = (*d)[j+1], (*d)[j]
			}
		}
		fmt.Println(d)
	}

	//for index, value := range *d {
	//	fmt.Println(index, value)
	//}

	return d

}

func main() {

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	d := make([]int, 0)

	for i := 0; i < 20; i++ {
		//for i := 0; i < 10; i++ {
		d = append(d, r.Intn(101))
	}

	fmt.Println(bubbleSort(&d))
}
