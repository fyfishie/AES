package lineShift

var lineShiftTable = [16]int{0, 1, 2, 3, 5, 6, 7, 4, 10, 11, 8, 9, 15, 12, 13, 14}
var iLineShiftTable = [16]int{0, 1, 2, 3, 7, 4, 5, 6, 10, 11, 8, 9, 13, 14, 15, 12}

func Shift(m [16]byte) (c [16]byte) {
	c = [16]byte{}
	for i := 0; i < 16; i++ {
		c[i] = m[lineShiftTable[i]]
	}
	return c
}

func IShift(c [16]byte) (m [16]byte) {
	m = [16]byte{}
	for i := 0; i < 16; i++ {
		m[i] = c[iLineShiftTable[i]]
	}
	return m
}
