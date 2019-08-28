package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s := []int{11, 22, 33, 44, 55}
	fmt.Println("Slice:", s)
	fmt.Println("Inverted slice:", invert(s))

	big := bigMap(1000, 8)
	test(big)
	testPtr(&big)

	fmt.Println(sum(10, 20, 30))

	someSlice := []string{"one", "two", "three"}
	someSlice = appendString(someSlice, "four")
	fmt.Println(someSlice)
}

func invert(slice []int) []int {
	n := len(slice)
	invSlice := make([]int, 0)
	for i := 0; i < n; i++ {
		invSlice = append(invSlice, slice[n-1-i])
	}
	return invSlice
}

func randString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz")
	l := len(letters)
	b := make([]rune, n)

	for i := range b {
		b[i] = letters[rand.Intn(l)]
	}
	return string(b)
}

func revertString(s string) string {
	reverse := []rune(s)
	for i, j := 0, len(reverse)-1; i < j; i, j = i+1, j-1 {
		reverse[i], reverse[j] = reverse[j], reverse[i]
	}
	return string(reverse)
}

func bigMap(n int, m int) map[string]string {
	big := make(map[string]string)
	for i := 0; i < n; i++ {
		str := randString(m)
		revStr := revertString(str)
		big[str] = revStr
		// fmt.Println("A very BIG map:")
		// fmt.Println("Key:", str, "Value:", revStr)
	}
	return big
}

func test(t map[string]string) {
	start := time.Now()
	elapsed := time.Since(start)
	fmt.Println("Test elapsed:", elapsed)
}

func testPtr(t *map[string]string) {
	start := time.Now()
	elapsed := time.Since(start)
	fmt.Println("Test pointer elapsed:", elapsed)
}

func sum(params ...int) int {
	var sum int
	for _, p := range params {
		sum += p
	}
	return sum
}

func appendString(s []string, val string) []string {
	if cap(s) >= len(s)+1 {
		s = append(s, val)
		return s
	} else {
		new := make([]string, len(s)+1, cap(s)*2)
		copy(new, s)
		new = append(new, val)
		return new
	}
}
