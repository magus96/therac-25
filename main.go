package main

import (
	"fmt"
	"time"
)

var mode string

func setMode(m string) {
	time.Sleep(2 * time.Second)
	mode = m
}

func irradiate(duration int, radMode string) int {
	var rad int
	if radMode == "E" {
		rad = 10000
	} else {
		rad = 10
	}
	return rad * duration
}

func main() {
	go setMode("E")
	time.Sleep(2 * time.Second)
	go setMode("X")
	rad := irradiate(10, mode)
	time.Sleep(10 * time.Second)
	fmt.Printf("You were irradiated with %d rads!!!\n", rad)
}
