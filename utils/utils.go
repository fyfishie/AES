package utils

func BlockXor(l [16]byte, r [16]byte) (c [16]byte) {
	c = [16]byte{}
	for i := 0; i < 16; i++ {
		c[i] = l[i] ^ r[i]
	}
	return c
}

func Order2Block(o [16]byte) (block [16]byte) {
	block = [16]byte{}
	for i := 0; i < 16; i++ {
		col := i / 4
		row := i % 4
		block[row*4+col] = o[i]
	}
	return block
}

func Block2Order(b [16]byte) (o [16]byte) {
	o = [16]byte{}
	oi := 0
	for col := 0; col < 4; col++ {
		for row := 0; row < 4; row++ {
			o[oi] = b[row*4+col]
			oi++
		}
	}
	return o
}
