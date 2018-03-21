package artist

import (
	"reflect"
	"testing"
)

func TestArtist_CreateBlankCanvas(t *testing.T) {
	var a Artist
	tables := []struct {
		width  int
		height int
		result [][]string
	}{
		{
			5,
			4,
			[][]string{
				{"-", "-", "-", "-", "-"},
				{"|", " ", " ", " ", "|"},
				{"|", " ", " ", " ", "|"},
				{"|", " ", " ", " ", "|"},
				{"|", " ", " ", " ", "|"},
				{"-", "-", "-", "-", "-"}},
		},
		{
			4,
			4,
			[][]string{
				{"-", "-", "-", "-"},
				{"|", " ", " ", "|"},
				{"|", " ", " ", "|"},
				{"|", " ", " ", "|"},
				{"|", " ", " ", "|"},
				{"-", "-", "-", "-"}},
		},
	}

	for _, table := range tables {
		result := a.CreateBlankCanvas(table.width, table.height)

		if !reflect.DeepEqual(result, table.result) {
			t.Errorf("CreateBlankCanvas(%v, %v) was incorrect, got: %v, want: %v.", table.width, table.height, result, table.result)
		}
	}
}

func TestDraw_Fill(t *testing.T) {
	var a Artist
	tables := []struct {
		ink    string
		x      int
		y      int
		canvas [][]string
		result [][]string
	}{
		{
			"@",
			4,
			3,
			[][]string{
				{" ", " ", " ", " ", " "},
				{" ", "x", "x", "x", " "},
				{" ", "x", " ", "x", " "},
				{" ", "x", " ", "x", " "},
				{" ", "x", "x", "x", " "},
				{" ", " ", " ", " ", " "}},
			[][]string{
				{"-", "-", "-", "-", "-"},
				{"|", "x", "x", "x", "|"},
				{"|", "x", " ", "x", "|"},
				{"|", "x", " ", "x", "|"},
				{"|", "x", "x", "x", "|"},
				{"-", "-", "-", "-", "-"}},
		},
		{
			"@",
			2,
			2,
			[][]string{
				{" ", " ", " ", " ", " "},
				{" ", "x", "x", "x", " "},
				{" ", "x", " ", "x", " "},
				{" ", "x", " ", "x", " "},
				{" ", "x", "x", "x", " "},
				{" ", " ", " ", " ", " "}},
			[][]string{
				{"-", "-", "-", "-", "-"},
				{"|", "x", "x", "x", "|"},
				{"|", "x", "@", "x", "|"},
				{"|", "x", "@", "x", "|"},
				{"|", "x", "x", "x", "|"},
				{"-", "-", "-", "-", "-"}},
		},
		{
			"@",
			1,
			1,
			[][]string{
				{" ", " ", " ", " ", " "},
				{" ", " ", "x", "x", "x"},
				{" ", " ", "x", " ", " "},
				{" ", " ", "x", " ", " "},
				{"x", "x", "x", " ", " "},
				{" ", " ", " ", " ", " "}},
			[][]string{
				{"-", "-", "-", "-", "-"},
				{"|", "@", "x", "x", "|"},
				{"|", "@", "x", " ", "|"},
				{"|", "@", "x", " ", "|"},
				{"|", "x", "x", " ", "|"},
				{"-", "-", "-", "-", "-"}},
		},
		{
			"@",
			0,
			4,
			[][]string{
				{" ", " ", " ", " ", " "},
				{" ", " ", "x", "x", "x"},
				{" ", " ", "x", " ", " "},
				{" ", " ", "x", " ", " "},
				{"x", "x", "x", " ", " "},
				{" ", " ", " ", " ", " "}},
			[][]string{
				{"-", "-", "-", "-", "-"},
				{"|", " ", "@", "@", "|"},
				{"|", " ", "@", " ", "|"},
				{"|", " ", "@", " ", "|"},
				{"|", "@", "@", " ", "|"},
				{"-", "-", "-", "-", "-"}},
		},
	}

	for _, table := range tables {
		result := a.CreateFill(table.ink, table.x, table.y, table.canvas)

		if !reflect.DeepEqual(result, table.result) {
			t.Errorf("CreateFill(%v, %v, %v, %v) was incorrect, got: %v, want: %v.", table.ink, table.x, table.y, table.canvas, result, table.result)
		}
	}
}

func TestArtist_CreateLine(t *testing.T) {
	var a Artist
	tables := []struct {
		x1     int
		y1     int
		x2     int
		y2     int
		canvas [][]string
		result [][]string
	}{
		{
			1,
			1,
			2,
			1,
			[][]string{
				{" ", " ", " ", " "},
				{" ", " ", " ", " "},
				{" ", " ", " ", " "},
				{" ", " ", " ", " "}},
			[][]string{
				{"-", "-", "-", "-"},
				{"|", "x", "x", "|"},
				{"|", " ", " ", "|"},
				{"-", "-", "-", "-"}},
		},
		{
			1,
			1,
			3,
			1,
			[][]string{
				{" ", " ", " ", " ", " "},
				{" ", " ", " ", " ", " "},
				{" ", " ", " ", " ", " "},
				{" ", " ", " ", " ", " "}},
			[][]string{
				{"-", "-", "-", "-", "-"},
				{"|", "x", "x", "x", "|"},
				{"|", " ", " ", " ", "|"},
				{"-", "-", "-", "-", "-"}},
		},
		{
			1,
			0,
			4,
			3,
			[][]string{
				{" ", " ", " ", " ", " "},
				{" ", " ", " ", " ", " "},
				{" ", " ", " ", " ", " "},
				{" ", " ", " ", " ", " "}},
			[][]string{
				{"-", "-", "-", "-", "-"},
				{"|", " ", "x", " ", "|"},
				{"|", " ", " ", "x", "|"},
				{"-", "-", "-", "-", "-"}},
		},
		{
			4,
			0,
			0,
			4,
			[][]string{
				{" ", " ", " ", " ", " "},
				{" ", " ", " ", " ", " "},
				{" ", " ", " ", " ", " "},
				{" ", " ", " ", " ", " "},
				{" ", " ", " ", " ", " "}},
			[][]string{
				{"-", "-", "-", "-", "-"},
				{"|", " ", " ", "x", "|"},
				{"|", " ", "x", "x", "|"},
				{"|", "x", " ", " ", "|"},
				{"-", "-", "-", "-", "-"}},
		},
		{
			4,
			0,
			1,
			4,
			[][]string{
				{" ", " ", " ", " ", " "},
				{" ", " ", " ", " ", " "},
				{" ", " ", " ", " ", " "},
				{" ", " ", " ", " ", " "},
				{" ", " ", " ", " ", " "}},
			[][]string{
				{"-", "-", "-", "-", "-"},
				{"|", " ", " ", "x", "|"},
				{"|", " ", "x", "x", "|"},
				{"|", " ", "x", " ", "|"},
				{"-", "-", "-", "-", "-"}},
		},
		{
			3,
			0,
			1,
			4,
			[][]string{
				{" ", " ", " ", " ", " "},
				{" ", " ", " ", " ", " "},
				{" ", " ", " ", " ", " "},
				{" ", " ", " ", " ", " "},
				{" ", " ", " ", " ", " "}},
			[][]string{
				{"-", "-", "-", "-", "-"},
				{"|", " ", "x", "x", "|"},
				{"|", " ", "x", "x", "|"},
				{"|", "x", "x", " ", "|"},
				{"-", "-", "-", "-", "-"}},
		},
	}

	for _, table := range tables {
		result := a.CreateLine(table.x1, table.y1, table.x2, table.y2, table.canvas)

		if !reflect.DeepEqual(result, table.result) {
			t.Errorf("CreateLine(%v, %v, %v, %v, %v) was incorrect, got: %v, want: %v.", table.x1, table.y1, table.x2, table.y2, table.canvas, result, table.result)
		}
	}
}

func TestArtist_CreateRectangle(t *testing.T) {
	var a Artist
	tables := []struct {
		x1     int
		y1     int
		x2     int
		y2     int
		canvas [][]string
		result [][]string
	}{
		{
			1,
			1,
			2,
			2,
			[][]string{
				{" ", " ", " ", " "},
				{" ", " ", " ", " "},
				{" ", " ", " ", " "},
				{" ", " ", " ", " "}},
			[][]string{
				{"-", "-", "-", "-"},
				{"|", "x", "x", "|"},
				{"|", "x", "x", "|"},
				{"-", "-", "-", "-"}},
		},
		{
			0,
			0,
			3,
			3,
			[][]string{
				{" ", " ", " ", " "},
				{" ", " ", " ", " "},
				{" ", " ", " ", " "},
				{" ", " ", " ", " "}},
			[][]string{
				{"-", "-", "-", "-"},
				{"|", " ", " ", "|"},
				{"|", " ", " ", "|"},
				{"-", "-", "-", "-"}},
		},
		{
			2,
			0,
			4,
			2,
			[][]string{
				{" ", " ", " ", " ", " "},
				{" ", " ", " ", " ", " "},
				{" ", " ", " ", " ", " "},
				{" ", " ", " ", " ", " "}},
			[][]string{
				{"-", "-", "-", "-", "-"},
				{"|", " ", "x", " ", "|"},
				{"|", " ", "x", "x", "|"},
				{"-", "-", "-", "-", "-"}},
		},
		{
			4,
			2,
			2,
			0,
			[][]string{
				{" ", " ", " ", " ", " "},
				{" ", " ", " ", " ", " "},
				{" ", " ", " ", " ", " "},
				{" ", " ", " ", " ", " "}},
			[][]string{
				{"-", "-", "-", "-", "-"},
				{"|", " ", "x", " ", "|"},
				{"|", " ", "x", "x", "|"},
				{"-", "-", "-", "-", "-"}},
		},
	}

	for _, table := range tables {
		result := a.CreateRectangle(table.x1, table.y1, table.x2, table.y2, table.canvas)

		if !reflect.DeepEqual(result, table.result) {
			t.Errorf("CreateRectangle(%v, %v, %v, %v, %v) was incorrect, got: %v, want: %v.", table.x1, table.y1, table.x2, table.y2, table.canvas, result, table.result)
		}
	}
}
