package AES

import (
	"gitee.com/fyfishie/AES/key"
	"gitee.com/fyfishie/AES/lineShift"
	"gitee.com/fyfishie/AES/mixCloumn"
	"gitee.com/fyfishie/AES/roundKey"
	"gitee.com/fyfishie/AES/sboxOp"
	"gitee.com/fyfishie/AES/utils"
)

func E(M [16]byte, KEY [16]byte) (C [16]byte) {
	rk := roundKey.Extend(KEY)
	CM := utils.Order2Block(M)
	CM = utils.BlockXor(rk[0], CM)
	for i := 0; i < 9; i++ {
		CM = sboxOp.Replace(CM)
		CM = lineShift.Shift(CM)
		CM = mixCloumn.BlockMix(CM)
		CM = utils.BlockXor(CM, rk[i+1])
	}
	CM = sboxOp.Replace(CM)
	CM = lineShift.Shift(CM)
	CM = utils.BlockXor(CM, rk[10])
	return utils.Block2Order(CM)
}

func D(C [16]byte, KEY [16]byte) (M [16]byte) {
	rk := roundKey.Extend(KEY)
	CM := utils.Order2Block(C)
	CM = utils.BlockXor(CM, rk[10])
	CM = lineShift.IShift(CM)
	CM = sboxOp.IReplace(CM)
	for i := 9; i > 0; i-- {
		CM = utils.BlockXor(CM, rk[i])
		CM = mixCloumn.IBlockMix(CM)
		CM = lineShift.IShift(CM)
		CM = sboxOp.IReplace(CM)
	}
	CM = utils.BlockXor(rk[0], CM)
	CM = utils.Block2Order(CM)
	return CM
}

func AESKey(sKey string) ([16]byte, error) {
	return key.MakeKey(sKey)
}
