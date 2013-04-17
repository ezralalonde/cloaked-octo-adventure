package main

import (
	"cloaked-octo-adventure/swap"
	"fmt"
)

func main() {
	var lines int
	_, err := fmt.Scanf("%d", &lines)
	if err != nil {
		panic("could not read lines")
	}
	for ii := 0; ii < lines; ii++ {
		var width int
		var train []int

		_, err := fmt.Scanf("%d", &width)
		if err != nil {
			fmt.Println(lines, width, err)
			panic("could not read width")
		}
		for jj := 0; jj < width; jj++ {
			var value int
			_, err := fmt.Scanf("%d", &value)
			if err != nil {
				panic("could not read value")
			}
			train = append(train, value)
		}
		fmt.Printf("Optimal train swapping takes %v swaps.\n", swap.Swap(train))
	}
}
