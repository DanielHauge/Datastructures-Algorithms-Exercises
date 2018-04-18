package main

import (
	"fmt"
	"math"
)

func main() {
	///// 10 tasks
	array := GetFibArray(100)
	fmt.Printf("%v\n",fractionCal(5))
	fmt.Printf("%v\n",calculateratio(array[6],array[5]))

	fmt.Printf("%v",findSqroot(5, 10))
	cal:=(1+findSqroot(5.0,2.0))/2
	fmt.Printf("%v\n",cal)

	fmt.Printf("%v\n", 1+2*math.Sin(math.Pi/10))
	fmt.Printf("%v", 1.61803398875)
	////// Mystery function
	fmt.Printf("%v", Mysery(3,25))



}

func calculateratio(n1 int, n2 int) float32{
	return float32(n1)/float32(n2)
}
func fractionCal(dex int)float64{
	res := -5.0
	mynumber := 1.0
	myNewnumber := 1.0
	for dex>0{
	dex--
	res = 1+mynumber/myNewnumber
	myNewnumber = res
	}
	return res
}

func GetFibArray(index int)[]int{
	previus := 1
	current := 1
	dex := 2
	array := []int{1, 1}

	for index>0{
		index--
		tempPrev := previus
		previus = current
		current = tempPrev+current
		array = append(array, current)
		dex++
	}
	return array

}
func findSqroot(OriginalNumber float64, guess float64)float64{

	CheckingGuess := OriginalNumber / guess
	nextgues := (CheckingGuess+guess) / 2.0
	if (guess == nextgues){
		return guess
	}else {
		return findSqroot(OriginalNumber, nextgues)
	}
}
func Mysery(a int, b int)int{
	if b == 0{ return 1}
	if b%2 == 0{return Mysery(a*a, b/2)}
	return Mysery(a*a, b/2)*a
}


