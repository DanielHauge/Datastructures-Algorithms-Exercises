package main

import (
	"testing"
	"math/rand"
)

func BenchmarkRandom(b *testing.B) {

	array := []int{rand.Int(), rand.Int(),rand.Int(),rand.Int(),rand.Int(),rand.Int(),rand.Int(),rand.Int(),rand.Int(),rand.Int(),rand.Int()}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		bubblesort(array)
	}
}

func BenchmarkAllreadySorted(b *testing.B) {

	array := []int{1,2,3,4,5,6,7,8,9, 10, 11}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bubblesort(array)
	}
}

func BenchmarkFirst(b *testing.B) {

	array := []int{12,2,3,4,5,6,7,8,9, 10, 11}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bubblesort(array)
	}
}

func BenchmarkLast(b *testing.B) {

	array := []int{1,2,3,4,5,6,7,8,9, 10, 0}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bubblesort(array)
	}
}
