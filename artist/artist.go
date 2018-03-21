package artist

import "github.com/clybs/discountPhotoShop/draw"

// Artist struct
type Artist struct{}

var pen draw.Draw
var ink = "x"

func (a *Artist) createBorder(canvas [][]string) [][]string {
	return pen.Border("-", "|", "|", "-", canvas)
}

// CreateBlankCanvas creates a blank canvas with border
func (a *Artist) CreateBlankCanvas(width, height int) [][]string {
	c := pen.Canvas(" ", width, height+2)
	return a.createBorder(c)
}

// CreateFill creates a fill on the canvas
func (a *Artist) CreateFill(color string, x, y int, canvas [][]string) [][]string {
	c := pen.Fill(color, x, y, canvas)
	return a.createBorder(c)
}

// CreateLine creates a line on the canvas
func (a *Artist) CreateLine(x1, y1, x2, y2 int, canvas [][]string) [][]string {
	c := pen.Line(ink, x1, y1, x2, y2, canvas)
	return a.createBorder(c)
}

// CreateRectangle creates a rectangle on the canvas
func (a *Artist) CreateRectangle(x1, y1, x2, y2 int, canvas [][]string) [][]string {
	c := pen.Rectangle(ink, x1, y1, x2, y2, canvas)
	return a.createBorder(c)
}
