// overflow

package main

import "math"

func main() {
	// 오버플로우 (범위초과)
	var n1 uint8 = math.MaxUint16 + 1
	// 오버플로우 (범위미만)
	var n2 uint8 = -1
}
