package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

const N  = 100
const M  = 10

var d = make([]int, 0)

func quickSort(left int, right int) {

	if left >= right {
		return
	}

	base := d[left]
	i := left
	j := right


	for i < j  {

		for d[j] >= base && i < j {
			j--
		}

		for d[i] <= base && i < j {
			i++
		}

		if i < j {
			d[i], d[j] = d[j], d[i]
			highlight(i, j)
		}

	}

	d[i], d[left] = d[left], d[i]
	highlightSwitchBase(left, i)

	quickSort(left, i-1)
	quickSort(i+1, right)
}

func highlight(i, j int)  {
	red := color.New(color.FgWhite, color.Bold, color.BgGreen)
	//fmt.Println()
	for index, v := range d {
		if index == i || index == j {
			red.Printf("%d", v)
			fmt.Printf(" ")
		} else {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Println()
}
func highlightSwitchBase(i, j int)  {
	red := color.New(color.FgWhite, color.Bold, color.BgYellow)
	//fmt.Println()
	for index, v := range d {
		if index == i || index == j {
			red.Printf("%d", v)
			fmt.Printf(" ")
		} else {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Println()
}

func main()  {
	newSeeder := rand.NewSource(time.Now().UnixNano())
	newRand := rand.New(newSeeder)

	//d := make([]int, 0)
	for i := 0; i < M; i++ {
		d = append(d, newRand.Intn(N + 1))
	}

	fmt.Println(d)

	quickSort(0, M - 1)

	fmt.Println(d)
}
