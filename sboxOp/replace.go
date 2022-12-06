package sboxOp

import "AES/lib"

func Replace(m [16]byte) (c [16]byte) {
	c = [16]byte{}
	for index, B := range m {
		CB := lib.S[int(B&0xf0>>4)][int(B&0x0f)]
		c[index] = CB
	}
	return c
}

func IReplace(c [16]byte) (m [16]byte) {
	m = [16]byte{}
	for index, C := range c {
		MB := lib.Is[int(C&0xf0)>>4][int(C&0x0f)]
		m[index] = MB
	}
	return m
}

func ReplaceOne(m byte) (c byte) {
	return lib.S[int(m&0xf0>>4)][int(m&0x0f)]
}

func IReplaceOne(c byte) (m byte) {
	return lib.Is[int(c&0xf0>>4)][int(c&0x0f)]
}
