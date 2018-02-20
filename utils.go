package main

import (
	"math"
	"strconv"
)

func parseFloat(s string) float64 {
	res, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return math.Floor(res + .5)
}

func parseInt(s string) int64 {
	res, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		panic(err)
	}
	return res
}
