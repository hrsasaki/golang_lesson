package main

import "strconv"

const (
	_  = iota
	KB = 1000 * iota
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB
	EB = 1000 * PB
	ZB = 1000 * EB
	YB = 1000 * ZB
)

func main() {
	println("KB: " + strconv.Itoa(KB))
	println("MB: " + strconv.Itoa(MB))
	println("GB: " + strconv.Itoa(GB))
	println("TB: " + strconv.Itoa(TB))
	println("PB: " + strconv.Itoa(PB))
	println("EB: " + strconv.Itoa(EB))
	// below: over 64 bit
	// println("1ZB = %dB", big.NewInt(ZB))
	// println("1YB = %dB", big.NewInt(YB))
}
