package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pmylund/sortutil"
)

func main() {
	seed := int(time.Now().Unix())
	fmt.Println(seed)
	rand.Seed(int64(seed))
	bound := 25
	points := make([]point, 0) //makean empty slice
	fmt.Println(points)
	points = append(points, point{
		x: bound,
		y: 0,
	}, point{
		x: -bound,
		y: 0,
	})

	fmt.Println(points)

	sortutil.AscByField(points, "x")
	fmt.Println(points)
}
func insert(a []int, element int, index int) []int {
	return append(a[:index], append([]int{element}, a[index:]...)...)
}

type point struct {
	x int
	y int
}

//https://web.archive.org/web/20170812230846/http://www.gameprogrammer.com/fractal.html
