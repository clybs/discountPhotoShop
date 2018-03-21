package draw

import (
	"math"
)

// Draw struct
type Draw struct{}

// Border creates a border
func (d *Draw) Border(top, left, right, bottom string, canvas [][]string) [][]string {
	newCanvas := make([][]string, len(canvas))
	for row := 0; row < len(canvas); row++ {
		newCanvas[row] = make([]string, len(canvas[row]))
		for col := 0; col < len(canvas[row]); col++ {
			if row == 0 {
				// Populate top border
				newCanvas[row][col] = top
			} else if row == len(canvas)-1 {
				// Populate bottom border
				newCanvas[row][col] = bottom
			} else if col == 0 {
				// Populate left border
				newCanvas[row][col] = left
			} else if col == len(canvas[row])-1 {
				// Populate right border
				newCanvas[row][col] = right
			} else {
				// Populate inside details
				newCanvas[row][col] = canvas[row][col]
			}
		}
	}
	return newCanvas
}

// Canvas creates a blank canvas
func (d *Draw) Canvas(spacer string, width, height int) [][]string {
	canvas := make([][]string, height)
	for row := 0; row < height; row++ {
		canvas[row] = make([]string, width)
		for col := 0; col < width; col++ {
			canvas[row][col] = spacer
		}
	}
	return canvas
}

// Fill bucket fills an area with specific character
func (d *Draw) Fill(ink string, x, y int, canvas [][]string) [][]string {
	// Get the character that will be replaced
	var originalCharacter string
	points := make(chan [][]int, 1)
	done := make(chan [][]int, 1)

	for row := 0; row < len(canvas); row++ {
		for col := 0; col < len(canvas[row]); col++ {
			if row == y && col == x {
				originalCharacter = canvas[row][col]
				break
			}
		}
	}

	// Start running threads
	go d.processPoints(originalCharacter, canvas, points, done)

	// Pass the original points
	points <- [][]int{{x, y}}

	// Pass the calculated fill points here
	fillPoints := <-done

	// Recreate the new canvas and bucket fill it
	for row := 0; row < len(canvas); row++ {
		for col := 0; col < len(canvas[row]); col++ {
			for _, j := range fillPoints {
				if j[0] == col && j[1] == row {
					canvas[row][col] = ink
				}
			}
		}
	}

	return canvas
}

// Line creates a line
func (d *Draw) Line(ink string, x1, y1, x2, y2 int, canvas [][]string) [][]string {
	newCanvas := make([][]string, len(canvas))
	isStraight := x1 == x2 || y1 == y2

	for row := 0; row < len(canvas); row++ {
		newCanvas[row] = make([]string, len(canvas[row]))

		// Check if straight line
		if isStraight {
			for col := 0; col < len(canvas[row]); col++ {
				middlePoint, _ := d.isMiddlePoint(col, row, x1, y1, x2, y2, len(canvas), len(canvas[row]))
				if middlePoint {
					// Draw points
					newCanvas[row][col] = ink
				} else {
					// Populate inside details
					newCanvas[row][col] = canvas[row][col]
				}
			}
		} else {
			// This is an inclined line. Do a look ahead and select lowest point.
			var lowestPointX float64
			for col := 0; col < len(canvas[row]); col++ {
				// Get possible lowest point
				_, middlePoint := d.isMiddlePoint(col, row, x1, y1, x2, y2, len(canvas), len(canvas[row]))
				tempLowestPointX := math.Abs(float64(middlePoint))
				if col == 0 {
					lowestPointX = tempLowestPointX
					continue
				}
				if tempLowestPointX < lowestPointX {
					lowestPointX = tempLowestPointX
				}
			}

			// Match result with lowest point
			for col := 0; col < len(canvas[row]); col++ {
				// Get possible lowest point
				_, middlePoint := d.isMiddlePoint(col, row, x1, y1, x2, y2, len(canvas), len(canvas[row]))

				isMiddlePointY := math.Abs(float64(col-middlePoint)) == 1
				isMiddlePointX := math.Abs(float64(middlePoint)) == lowestPointX
				isMiddlePoint := isMiddlePointX || isMiddlePointY

				// Match found
				if isMiddlePoint {
					// Draw points
					newCanvas[row][col] = ink
				} else {
					// Populate inside details
					newCanvas[row][col] = canvas[row][col]
				}
			}
		}
	}
	return newCanvas
}

// Rectangle creates a rectangle
func (d *Draw) Rectangle(ink string, x1, y1, x2, y2 int, canvas [][]string) [][]string {
	// Create the side 1 line
	newCanvas := d.Line(ink, x1, y1, x2, y1, canvas)

	// Create the side 2 line
	newCanvas = d.Line(ink, x2, y1, x2, y2, newCanvas)

	// Create the side 3 line
	newCanvas = d.Line(ink, x2, y2, x1, y2, newCanvas)

	// Create the side 4 line and return
	return d.Line(ink, x1, y2, x1, y1, newCanvas)
}
