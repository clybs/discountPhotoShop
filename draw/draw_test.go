package draw

import (
	"reflect"
	"testing"
)

func TestIsMiddlePoint(t *testing.T) {
	var d Draw
	tables := []struct {
		valX   int
		valY   int
		x1     int
		y1     int
		x2     int
		y2     int
		width  int
		height int
		result bool
	}{
		{
			2,
			1,
			1,
			1,
			2,
			1,
			5,
			4,
			true,
		},
		{
			2,
			2,
			0,
			1,
			5,
			1,
			5,
			4,
			false,
		},
	}

	for _, table := range tables {
		result, _ := d.isMiddlePoint(table.valX, table.valY, table.x1, table.y1, table.x2, table.y2, table.width, table.height)

		if !reflect.DeepEqual(result, table.result) {
			t.Errorf("isMiddlePoint(%v, %v, %v, %v, %v, %v, %v, %v) was incorrect, got: %v, want: %v.", table.valX, table.valY, table.x1, table.y1, table.x2, table.y2, table.width, table.height, result, table.result)
		}
	}
}

func TestDraw_Border(t *testing.T) {
	var d Draw
	tables := []struct {
		top    string
		left   string
		right  string
		bottom string
		canvas [][]string
		result [][]string
	}{
		{
			"a",
			"b",
			"c",
			"d",
			[][]string{
				{"*", "*"},
				{"*", "*"}},
			[][]string{
				{"a", "a"},
				{"d", "d"}},
		}, {
			"a",
			"b",
			"c",
			"d",
			[][]string{
				{"*", "*", "*"},
				{"*", "*", "*"},
				{"*", "*", "*"}},
			[][]string{
				{"a", "a", "a"},
				{"b", "*", "c"},
				{"d", "d", "d"}},
		},
		{
			"a",
			"b",
			"c",
			"d",
			[][]string{
				{"*", "*", "*", "*"},
				{"*", "*", "*", "*"},
				{"*", "*", "*", "*"},
				{"*", "*", "*", "*"}},
			[][]string{
				{"a", "a", "a", "a"},
				{"b", "*", "*", "c"},
				{"b", "*", "*", "c"},
				{"d", "d", "d", "d"}},
		},
		{
			"a",
			"b",
			"c",
			"d",
			[][]string{
				{"*"},
				{"*"}},
			[][]string{
				{"a"},
				{"d"}},
		},
		{
			"a",
			"b",
			"c",
			"d",
			[][]string{
				{"*"}},
			[][]string{
				{"a"}},
		},
	}

	for _, table := range tables {
		result := d.Border(table.top, table.left, table.right, table.bottom, table.canvas)

		if !reflect.DeepEqual(result, table.result) {
			t.Errorf("Border(%v, %v, %v, %v, %v) was incorrect, got: %v, want: %v.", table.top, table.left, table.right, table.bottom, table.canvas, result, table.result)
		}
	}
}

func TestDraw_Canvas(t *testing.T) {
	var d Draw
	tables := []struct {
		spacer string
		width  int
		height int
		result [][]string
	}{
		{
			"*",
			1,
			4,
			[][]string{
				{"*"},
				{"*"},
				{"*"},
				{"*"}},
		},
		{
			"*",
			2,
			3,
			[][]string{
				{"*", "*"},
				{"*", "*"},
				{"*", "*"}},
		},
	}

	for _, table := range tables {
		result := d.Canvas(table.spacer, table.width, table.height)

		if !reflect.DeepEqual(result, table.result) {
			t.Errorf("Canvas(%v, %v) was incorrect, got: %v, want: %v.", table.width, table.height, result, table.result)
		}
	}
}

func TestDraw_Fill(t *testing.T) {
	var d Draw
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
				{" ", "*", "*", "*", " "},
				{" ", "*", " ", "*", " "},
				{" ", "*", " ", "*", " "},
				{" ", "*", "*", "*", " "},
				{" ", " ", " ", " ", " "}},
			[][]string{
				{"@", "@", "@", "@", "@"},
				{"@", "*", "*", "*", "@"},
				{"@", "*", " ", "*", "@"},
				{"@", "*", " ", "*", "@"},
				{"@", "*", "*", "*", "@"},
				{"@", "@", "@", "@", "@"}},
		},
		{
			"@",
			2,
			2,
			[][]string{
				{" ", " ", " ", " ", " "},
				{" ", "*", "*", "*", " "},
				{" ", "*", " ", "*", " "},
				{" ", "*", " ", "*", " "},
				{" ", "*", "*", "*", " "},
				{" ", " ", " ", " ", " "}},
			[][]string{
				{" ", " ", " ", " ", " "},
				{" ", "*", "*", "*", " "},
				{" ", "*", "@", "*", " "},
				{" ", "*", "@", "*", " "},
				{" ", "*", "*", "*", " "},
				{" ", " ", " ", " ", " "}},
		},
		{
			"@",
			1,
			1,
			[][]string{
				{" ", " ", " ", " ", " "},
				{" ", " ", "*", "*", "*"},
				{" ", " ", "*", " ", " "},
				{" ", " ", "*", " ", " "},
				{"*", "*", "*", " ", " "},
				{" ", " ", " ", " ", " "}},
			[][]string{
				{"@", "@", "@", "@", "@"},
				{"@", "@", "*", "*", "*"},
				{"@", "@", "*", " ", " "},
				{"@", "@", "*", " ", " "},
				{"*", "*", "*", " ", " "},
				{" ", " ", " ", " ", " "}},
		},
		{
			"@",
			0,
			4,
			[][]string{
				{" ", " ", " ", " ", " "},
				{" ", " ", "*", "*", "*"},
				{" ", " ", "*", " ", " "},
				{" ", " ", "*", " ", " "},
				{"*", "*", "*", " ", " "},
				{" ", " ", " ", " ", " "}},
			[][]string{
				{" ", " ", " ", " ", " "},
				{" ", " ", "@", "@", "@"},
				{" ", " ", "@", " ", " "},
				{" ", " ", "@", " ", " "},
				{"@", "@", "@", " ", " "},
				{" ", " ", " ", " ", " "}},
		},
	}

	for _, table := range tables {
		result := d.Fill(table.ink, table.x, table.y, table.canvas)

		if !reflect.DeepEqual(result, table.result) {
			t.Errorf("Fill(%v, %v, %v, %v) was incorrect, got: %v, want: %v.", table.ink, table.x, table.y, table.canvas, result, table.result)
		}
	}
}

func TestDraw_Line(t *testing.T) {
	var d Draw
	tables := []struct {
		ink    string
		x1     int
		y1     int
		x2     int
		y2     int
		canvas [][]string
		result [][]string
	}{
		{
			"@",
			1,
			1,
			2,
			1,
			[][]string{
				{"*", "*", "*", "*"},
				{"*", "*", "*", "*"},
				{"*", "*", "*", "*"},
				{"*", "*", "*", "*"}},
			[][]string{
				{"*", "*", "*", "*"},
				{"*", "@", "@", "*"},
				{"*", "*", "*", "*"},
				{"*", "*", "*", "*"}},
		},
		{
			"@",
			1,
			1,
			3,
			1,
			[][]string{
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"}},
			[][]string{
				{"*", "*", "*", "*", "*"},
				{"*", "@", "@", "@", "*"},
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"}},
		},
		{
			"@",
			1,
			2,
			2,
			2,
			[][]string{
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"}},
			[][]string{
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"},
				{"*", "@", "@", "*", "*"},
				{"*", "*", "*", "*", "*"}},
		},
		{
			"@",
			4,
			0,
			0,
			0,
			[][]string{
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"}},
			[][]string{
				{"@", "@", "@", "@", "@"},
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"}},
		},
		{
			"@",
			1,
			1,
			1,
			2,
			[][]string{
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"}},
			[][]string{
				{"*", "*", "*", "*", "*"},
				{"*", "@", "*", "*", "*"},
				{"*", "@", "*", "*", "*"},
				{"*", "*", "*", "*", "*"}},
		},
		{
			"@",
			1,
			0,
			4,
			3,
			[][]string{
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"}},
			[][]string{
				{"*", "@", "@", "*", "*"},
				{"*", "*", "@", "*", "*"},
				{"*", "*", "*", "@", "@"},
				{"*", "*", "*", "*", "@"}},
		},
		{
			"@",
			4,
			0,
			0,
			4,
			[][]string{
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"}},
			[][]string{
				{"*", "*", "*", "*", "@"},
				{"*", "*", "*", "@", "*"},
				{"*", "*", "@", "@", "*"},
				{"*", "@", "*", "*", "*"},
				{"@", "*", "*", "*", "*"}},
		},
		{
			"@",
			4,
			0,
			1,
			4,
			[][]string{
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"}},
			[][]string{
				{"*", "*", "*", "*", "@"},
				{"*", "*", "*", "@", "@"},
				{"*", "*", "@", "@", "*"},
				{"*", "*", "@", "*", "*"},
				{"*", "@", "*", "*", "*"}},
		},
		{
			"@",
			3,
			0,
			1,
			4,
			[][]string{
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"}},
			[][]string{
				{"*", "*", "*", "@", "*"},
				{"*", "*", "@", "@", "*"},
				{"*", "*", "@", "@", "*"},
				{"*", "@", "@", "*", "*"},
				{"*", "@", "*", "*", "*"}},
		},
	}

	for _, table := range tables {
		result := d.Line(table.ink, table.x1, table.y1, table.x2, table.y2, table.canvas)

		if !reflect.DeepEqual(result, table.result) {
			t.Errorf("Line(%v, %v, %v, %v, %v, %v) was incorrect, got: %v, want: %v.", table.ink, table.x1, table.y1, table.x2, table.y2, table.canvas, result, table.result)
		}
	}
}

func TestDraw_Rectangle(t *testing.T) {
	var d Draw
	tables := []struct {
		ink    string
		x1     int
		y1     int
		x2     int
		y2     int
		canvas [][]string
		result [][]string
	}{
		{
			"@",
			1,
			1,
			2,
			2,
			[][]string{
				{"*", "*", "*", "*"},
				{"*", "*", "*", "*"},
				{"*", "*", "*", "*"},
				{"*", "*", "*", "*"}},
			[][]string{
				{"*", "*", "*", "*"},
				{"*", "@", "@", "*"},
				{"*", "@", "@", "*"},
				{"*", "*", "*", "*"}},
		},
		{
			"@",
			0,
			0,
			3,
			3,
			[][]string{
				{"*", "*", "*", "*"},
				{"*", "*", "*", "*"},
				{"*", "*", "*", "*"},
				{"*", "*", "*", "*"}},
			[][]string{
				{"@", "@", "@", "@"},
				{"@", "*", "*", "@"},
				{"@", "*", "*", "@"},
				{"@", "@", "@", "@"}},
		},
		{
			"@",
			2,
			0,
			4,
			2,
			[][]string{
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"}},
			[][]string{
				{"*", "*", "@", "@", "@"},
				{"*", "*", "@", "*", "@"},
				{"*", "*", "@", "@", "@"},
				{"*", "*", "*", "*", "*"}},
		},
		{
			"@",
			4,
			2,
			2,
			0,
			[][]string{
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"},
				{"*", "*", "*", "*", "*"}},
			[][]string{
				{"*", "*", "@", "@", "@"},
				{"*", "*", "@", "*", "@"},
				{"*", "*", "@", "@", "@"},
				{"*", "*", "*", "*", "*"}},
		},
	}

	for _, table := range tables {
		result := d.Rectangle(table.ink, table.x1, table.y1, table.x2, table.y2, table.canvas)

		if !reflect.DeepEqual(result, table.result) {
			t.Errorf("Rectangle(%v, %v, %v, %v, %v, %v) was incorrect, got: %v, want: %v.", table.ink, table.x1, table.y1, table.x2, table.y2, table.canvas, result, table.result)
		}
	}
}

func BenchmarkDraw_Fill(b *testing.B) {
	var d Draw
	var canvas = [][]string{
		{" ", " ", " ", " ", " "},
		{" ", "*", "*", "*", " "},
		{" ", "*", " ", "*", " "},
		{" ", "*", " ", "*", " "},
		{" ", "*", "*", "*", " "},
		{" ", " ", " ", " ", " "}}
	for n := 0; n < b.N; n++ {
		d.Fill("@", 4, 3, canvas)
	}
}

// 50000	     31459 ns/op
