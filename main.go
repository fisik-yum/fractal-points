package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := int(time.Now().Unix())
	fmt.Println(seed)
	rand.Seed(int64(seed))
	uBound := 25
	lBound := -uBound
	x := make([]int, 0) //makean empty in
	//y := make([]int, 2, 2)
	x = append(x, lBound)
	x = append(x, uBound)
	fmt.Println(x)
	x = insert(x, 12, 1) //insert at index 1 of x
	fmt.Println(x)
}
func insert(a []int, element int, index int) []int {
	return append(a[:index], append([]int{element}, a[index:]...)...)
}

//https://web.archive.org/web/20170812230846/http://www.gameprogrammer.com/fractal.html
