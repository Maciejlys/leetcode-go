package leetcode

func findTheDifference(s string, t string) byte {
	var sBytes byte
	var tBytes byte

	for _, r := range s {
		sBytes += byte(r)
	}

	for _, r := range t {
		tBytes += byte(r)
	}

	return tBytes - sBytes
}
