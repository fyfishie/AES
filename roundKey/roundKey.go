package roundKey

import "gitee.com/fyfishie/AES/sboxOp"

var TTable = [10][4]byte{
	{0x01, 0x00, 0x00, 0x00}, {0x02, 0x00, 0x00, 0x00},
	{0x04, 0x00, 0x00, 0x00}, {0x08, 0x00, 0x00, 0x00},
	{0x10, 0x00, 0x00, 0x00}, {0x20, 0x00, 0x00, 0x00},
	{0x40, 0x00, 0x00, 0x00}, {0x80, 0x00, 0x00, 0x00},
	{0x1b, 0x00, 0x00, 0x00}, {0x36, 0x00, 0x00, 0x00}}

func Extend(key [16]byte) [11][16]byte {
	ws := getW(key)
	wss := [44][4]byte{}
	wss[0] = ws[0]
	wss[1] = ws[1]
	wss[2] = ws[2]
	wss[3] = ws[3]
	for i := 4; i < 44; i++ {
		if i%4 != 0 {
			wss[i][0] = wss[i-4][0] ^ wss[i-1][0]
			wss[i][1] = wss[i-4][1] ^ wss[i-1][1]
			wss[i][2] = wss[i-4][2] ^ wss[i-1][2]
			wss[i][3] = wss[i-4][3] ^ wss[i-1][3]
		} else {
			TS := T(wss[i-1], i/4-1)
			wss[i][0] = wss[i-4][0] ^ TS[0]
			wss[i][1] = wss[i-4][1] ^ TS[1]
			wss[i][2] = wss[i-4][2] ^ TS[2]
			wss[i][3] = wss[i-4][3] ^ TS[3]
		}
	}
	res := [11][16]byte{}
	for i := 0; i < 11; i++ {
		r := [16]byte{}
		for j := 0; j < 16; j++ {
			r[j] = wss[i*4+j%4][j/4]
		}
		res[i] = r
	}
	return res
}

func getW(key [16]byte) [4][4]byte {
	ws := [4][4]byte{}
	for index, k := range key {
		row := index / 4
		col := index % 4
		ws[col][row] = k
	}
	return ws
}

func T(M [4]byte, round int) [4]byte {
	tmp := M[0]
	M[0] = M[1]
	M[1] = M[2]
	M[2] = M[3]
	M[3] = tmp
	for index, m := range M {
		M[index] = sboxOp.ReplaceOne(m)
	}
	M[0] = M[0] ^ TTable[round][0]
	M[1] = M[1] ^ TTable[round][1]
	M[2] = M[2] ^ TTable[round][2]
	M[3] = M[3] ^ TTable[round][3]
	return M
}
