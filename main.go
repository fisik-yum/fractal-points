package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/pmylund/sortutil"
)

func main() {
	seed := int(time.Now().Unix())
	fmt.Println(seed)
	rand.Seed(int64(seed))
	bound := 25.0
	points := make([]point, 0) //makean empty slice
	fmt.Println(points)
	points = append(points, point{
		x: bound,
		y: 0,
	}, point{
		x: -bound,
		y: 0,
	})

	//fmt.Println(points)

	sortutil.AscByField(points, "x")
	//fmt.Println(points)

	points = createPoints(points, 4)
	fmt.Println("Created slice of points", strconv.Itoa(len(points)), "elements long")
	fmt.Println(points)
}
func createPoints(points []point, iterations int) []point {
	for x := 1; x <= iterations; x++ {
		j := len(points)
		pointsTwo := make([]point, 0)
		//fmt.Println(x)
		//points1 := make([]point, 0)
		for i := 0; i < j-1; i++ {
			//fmt.Println(midpoint(points, i))
			pointsTwo = append(pointsTwo, midpoint(points, i))
			//fmt.Println(pointsTwo)

		}

		//sortutil.AscByField(points, "x")
		points = append(points, pointsTwo...)
		sortutil.AscByField(points, "x")
		//fmt.Println(points)

	}
	return points
}

func midpoint(a []point, index int) point {
	//fmt.Println(point{x: ((a[index].x + a[index+1].x) / 2), y: ((a[index].y + a[index+1].y) / 2)})
	return point{x: ((a[index].x + a[index+1].x) / 2), y: ((a[index].y + a[index+1].y) / 2)}
}

type point struct {
	x float64
	y float64
}

/*func insert(a []int, element int, index int) []int {
	return append(a[:index], append([]int{element}, a[index:]...)...)
}*/

//https://web.archive.org/web/20170812230846/http://www.gameprogrammer.com/fractal.html
