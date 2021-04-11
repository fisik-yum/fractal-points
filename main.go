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
	y := make([]int, 0)
	x = append(x, lBound, uBound)
	y = append(y, 0, 0)
	fmt.Println(x)
	fmt.Println(y)
	//x = insert(x, 25, 1)
	fmt.Println(x)
	for i := 0; i <= 2; i++ {
		midpoint := int((x[i] + x[i+1]) / 2)
		x = insert(x, midpoint, i+1)
		fmt.Println(x)
	}
}
func insert(a []int, element int, index int) []int {
	return append(a[:index], append([]int{element}, a[index:]...)...)
}

//https://web.archive.org/web/20170812230846/http://www.gameprogrammer.com/fractal.html
