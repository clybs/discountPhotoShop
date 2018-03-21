package cmd

import (
	"strconv"
	"strings"
)

func expectedParamLength(v string) int {
	switch strings.ToUpper(v) {
	case "C":
		return 2
	case "L":
		return 4
	case "R":
		return 4
	case "B":
		return 3
	case "Q":
		return -1
	}
	return 0
}

func getCommandParams(v string) []string {
	p := normalizeInput(v)
	return strings.Split(p, " ")
}

func isStringInSlice(v string, list []string) bool {
	for _, item := range list {
		if item == v {
			return true
		}
	}
	return false
}

func normalizeInput(v string) string {
	v = strings.Join(strings.Fields(v), " ")
	return strings.TrimSpace(v)
}

func toInt(v string) int {
	result, _ := strconv.Atoi(v)
	return result
}
