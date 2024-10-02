package main

import (
	"fmt"
)

func main() {
	helloWorld()
	loops()
	sum := sum(5, 4)
	sumVal, diffVal := sumAndDifference(7, 5)
	fmt.Println(sum, sumVal, diffVal)

	// Arrays
	var a4 [4]int
	a5 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a4, a5)

	a4_cpy := a4
	a4_cpy[0] = 1 // Only a4_cpy is changed, a4_cpy and a4 are two separate instances

	// Slices
	s3 := []int{1, 2, 3}
	s4 := make([]int, 4)
	var s5 [][]float64
	s3_cpy := s3
	s3_cpy[0] = 5 // s3_cpy and s3 point to the same instance, which means both are updated
	s6 := append(s4, 7, 8, 9)
	s7 := append(s6, []int{1, 2, 3}...)
	fmt.Println(s3, s4, s5, s6, s7)

	m := map[string]int{"one": 1, "two": 2}
	fmt.Println(m, m["one"])

	if val, ok := m["two"]; ok {
		fmt.Println(val)
	}

	// Type switch
	var data interface{}
	data = ""
	switch c := data.(type) {
	case string:
		fmt.Println(c, "is a string")
	case int64:
		fmt.Printf("%d is an int64\n", c)
	default:
		// all other cases
	}

	for key, value := range map[string]int{"one": 1, "two": 2, "three": 3} {
		fmt.Printf("key=%s, value=%d\n", key, value)
	}

	for key := range map[string]int{"one": 1, "two": 2, "three": 3} {
		fmt.Printf("key=%s", key)
	}

	// Function literals
	boolVal := func() bool {
		return 1 < 2
	}
	fmt.Println(boolVal())

	fmt.Println("Add + double two numbers:",
		func(a, b int) int {
			return (a + b) * 2
		}(10, 2))

	fmt.Println(sentenceFactory("the")("what", "hell"))

	deferConcept()

	concurrencyConcept()
}

func helloWorld() {
	fmt.Println("Hello world")
}

func loops() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := range 3 {
		fmt.Println(j)
	}

	for k := 0; k < 3; k++ {
		fmt.Println(k)
	}

	l := 1
	for {
		fmt.Println(l)
		if l == 4 {
			break
		}
		l++
	}

	for m := range 6 {
		if m%2 == 0 {
			continue
		}
		fmt.Println(m)
	}
}

/*
func functionName(params) returnType {}
*/
func sum(a int, b int) int {
	return a + b
}

/*
multi return functions
func functionName(params) (returnType, returnType) {} || func functionName(params) (returnVal, returnVal)
*/
func sumAndDifference(a, b int) (sum, difference int) {
	return a + b, a - b
}

func sentenceFactory(input string) func(before, after string) string {
	return func(before, after string) string {
		return fmt.Sprintf("%s %s %s", before, input, after)
	}
}

func deferConcept() bool {
	defer fmt.Println("deferred statements execute in LIFO sequence")
	defer fmt.Println("This line is executed first because")
	return true
}

func inc(i int, c chan int) {
	c <- i + 1 // <- is the "send" operator when a channel appears on the left.
}

func concurrencyConcept() {
	c := make(chan int)

	go inc(0, c)
	go inc(10, c)
	go inc(-805, c)

	fmt.Println(<-c, <-c, <-c)
}
