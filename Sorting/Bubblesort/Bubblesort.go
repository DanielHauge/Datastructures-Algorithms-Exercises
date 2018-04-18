package main

import "fmt"



func main() {
	toBeSorted := []int{1,3,2,4,8,6,7,2,3,0}
	sortedarray := bubblesort(toBeSorted)
	fmt.Println(sortedarray)

}


func bubblesort(input []int)[]int {
	res := input
	n := len(input)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n-1; i++{
			if res[i-1] > input[i]{

			res[i], res[i-1] = res[i-1], res[i]
			swapped = true
			}
		}
	}

	return res
}


