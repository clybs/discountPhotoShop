package draw

import (
	"fmt"
	"math"
	"strings"
)

// Display outputs the result to terminal
func (d *Draw) Display(canvas [][]string) {
	for _, v := range canvas {
		data := strings.Join(v, " ")
		fmt.Println(data)
	}
}

func (d *Draw) getSimilarNeighbors(x, y int, originalCharacter string, canvas [][]string) [][]int {
	validNeighbors := make([][]int, 0)
	similarNeighbors := make([][]int, 0)

	// Get all neighbors
	neighbors := [][]int{
		{x - 1, y - 1}, {x, y - 1}, {x + 1, y - 1}, {x - 1, y},
		{x + 1, y}, {x - 1, y + 1}, {x, y + 1}, {x + 1, y + 1},
	}

	// Get max canvas width and height
	var canvasWidth, canvasHeight int

	canvasHeight = len(canvas)
	if canvasHeight > 0 {
		canvasWidth = len(canvas[0])
	}

	// Save valid neighbors
	for _, v := range neighbors {
		if v[0] < canvasWidth && v[1] < canvasHeight && v[0] >= 0 && v[1] >= 0 {
			validNeighbors = append(validNeighbors, v)
		}
	}

	// Select among valid neighbors which ones are similar in character
	for _, v := range validNeighbors {
		if canvas[v[1]][v[0]] == originalCharacter {
			similarNeighbors = append(similarNeighbors, v)
		}
	}

	return similarNeighbors
}

func (d *Draw) isPointInSlice(v []int, list [][]int) bool {
	for _, l := range list {
		if l[0] == v[0] && l[1] == v[1] {
			return true
		}
	}
	return false
}

func (d *Draw) isMiddlePoint(valX, valY, x1, y1, x2, y2, width, height int) (bool, int) {
	// Check if within the range
	xa := x1 <= valX && valX <= x2
	xb := x2 <= valX && valX <= x1

	ya := y1 <= valY && valY <= y2
	yb := y2 <= valY && valY <= y1

	inRange := (xa || xb) && (ya || yb)

	// Check if within the angle
	dxc := valX - x1
	dyc := valY - y1

	dxl := x2 - x1
	dyl := y2 - y1

	cross := dxc*dyl - dyc*dxl

	return cross == 0 && inRange, cross
}

func (d *Draw) processPoints(originalCharacter string, canvas [][]string, points chan [][]int, fillPoints chan<- [][]int) {
	// Create container for fill points
	fillPts := make([][]int, 0)

	for {
		pts, more := <-points
		if more {
			// Check for new points
			newPts := make([][]int, 0)

			for _, j := range pts {
				nearPts := d.getSimilarNeighbors(j[0], j[1], originalCharacter, canvas)
				newPts = append(newPts, nearPts...)
			}

			// Make sure there are no cross overs
			uniquePtsLocal := make([][]int, 0)
			for _, v := range newPts {
				if !d.isPointInSlice(v, uniquePtsLocal) {
					uniquePtsLocal = append(uniquePtsLocal, v)
				}
			}

			uniquePtsGlobal := make([][]int, 0)
			for _, v := range uniquePtsLocal {
				if !d.isPointInSlice(v, fillPts) {
					uniquePtsGlobal = append(uniquePtsGlobal, v)
				}
			}

			if len(uniquePtsGlobal) > 0 {
				fillPts = append(fillPts, uniquePtsGlobal...)
				points <- uniquePtsGlobal
			} else {
				close(points)
			}
		} else {
			fillPoints <- fillPts
			return
		}
	}
}

func (d *Draw) round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func (d *Draw) toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(d.round(num*output)) / output
}
