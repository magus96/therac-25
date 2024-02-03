package main

import (
	"fmt"
	"time"
)

var mode string

func setMode(m string) {
	time.Sleep(8 * time.Second)
	mode = m
}

func irradiate(duration int, radMode string) int {
	var rad int
	if radMode == "E" {
		rad = 10000
	} else {
		rad = 10
	}
	time.Sleep(time.Duration(duration) * time.Second)
	return rad * duration
}

func main() {
	go setMode("E")
	time.Sleep(8 * time.Second)
	go setMode("X")
	rad := irradiate(10, mode)
	fmt.Printf("You were irradiated with %d rads!!!\n", rad)
}
