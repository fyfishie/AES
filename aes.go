package AES

import (
	"gitee.com/fyfishie/AES/lineShift"
	"gitee.com/fyfishie/AES/mixCloumn"
	"gitee.com/fyfishie/AES/roundKey"
	"gitee.com/fyfishie/AES/sboxOp"
)

func E(M [16]byte, KEY [16]byte) (C [16]byte) {
	rk := roundKey.Extend(KEY)
	CM := order2Block(M)
	CM = blockXor(rk[0], CM)
	for i := 0; i < 9; i++ {
		CM = sboxOp.Replace(CM)
		CM = lineShift.Shift(CM)
		CM = mixCloumn.BlockMix(CM)
		CM = blockXor(CM, rk[i+1])
	}
	CM = sboxOp.Replace(CM)
	CM = lineShift.Shift(CM)
	CM = blockXor(CM, rk[10])
	return block2Order(CM)
}

func D(C [16]byte, KEY [16]byte) (M [16]byte) {
	rk := roundKey.Extend(KEY)
	CM := order2Block(C)
	CM = blockXor(CM, rk[10])
	CM = lineShift.IShift(CM)
	CM = sboxOp.IReplace(CM)
	for i := 9; i > 0; i-- {
		CM = blockXor(CM, rk[i])
		CM = mixCloumn.IBlockMix(CM)
		CM = lineShift.IShift(CM)
		CM = sboxOp.IReplace(CM)
	}
	CM = blockXor(rk[0], CM)
	CM = block2Order(CM)
	return CM
}
func blockXor(l [16]byte, r [16]byte) (c [16]byte) {
	c = [16]byte{}
	for i := 0; i < 16; i++ {
		c[i] = l[i] ^ r[i]
	}
	return c
}

func order2Block(o [16]byte) (block [16]byte) {
	block = [16]byte{}
	for i := 0; i < 16; i++ {
		col := i / 4
		row := i % 4
		block[row*4+col] = o[i]
	}
	return block
}

func block2Order(b [16]byte) (o [16]byte) {
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
