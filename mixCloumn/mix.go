package mixCloumn

var mixTable = [4][4]byte{{0x02, 0x03, 0x01, 0x01}, {0x01, 0x02, 0x03, 0x01}, {0x01, 0x01, 0x02, 0x03}, {0x03, 0x01, 0x01, 0x02}}
var iMixTable = [4][4]byte{{0x0e, 0x0b, 0x0d, 0x09}, {0x09, 0x0e, 0x0b, 0x0d}, {0x0d, 0x09, 0x0e, 0x0b}, {0x0b, 0x0d, 0x09, 0x00e}}

func x1(B byte) byte {
	return B
}

func x2(B byte) byte {
	if B&0x80 == 0 {
		return B << 1
	}
	return (B << 1) ^ 0x1B

}

func x3(B byte) byte {
	return B ^ x2(B)
}

/*
0 4 8 12 1 5 9 13 2 6 10 14 3 7 11 15
*/
func BlockMix(M [16]byte) (C [16]byte) {
	C = [16]byte{}
	for index := 0; index < 16; index++ {
		row := index / 4
		col := index % 4
		rowMix := mixTable[row]
		colMix := getColMix(M, col)
		C[index] = multi(rowMix, colMix)
	}
	return C
}

func getColMix(M [16]byte, col int) (C [4]byte) {
	C = [4]byte{}
	offset := col % 4
	for i := 0; i < 4; i++ {
		C[i] = M[offset+i*4]
	}
	return C
}

func multi(row [4]byte, col [4]byte) byte {
	ors := [4]byte{}
	for i := 0; i < 4; i++ {
		switch row[i] {
		case 0x02:
			ors[i] = x2(col[i])
		case 0x01:
			ors[i] = x1(col[i])
		case 0x03:
			ors[i] = x3(col[i])
		case 0x09:
			ors[i] = x9(col[i])
		case 0x0b:
			ors[i] = xB(col[i])
		case 0x0d:
			ors[i] = xD(col[i])
		case 0x0e:
			ors[i] = xE(col[i])
		}
	}
	return ors[0] ^ ors[1] ^ ors[2] ^ ors[3]
}
func x9(B byte) byte {
	return x2(x2(x2(B))) ^ B
}

func xB(B byte) byte {
	return x2(x2(x2(B))) ^ x2(B) ^ B
}

func xD(B byte) byte {
	return x2(x2(x2(B))) ^ x2(x2(B)) ^ B
}

func xE(B byte) byte {
	return x2(x2(x2(B))) ^ x2(x2(B)) ^ x2(B)
}

func IBlockMix(C [16]byte) (M [16]byte) {
	M = [16]byte{}
	for index := 0; index < 16; index++ {
		row := index / 4
		col := index % 4
		rowMix := iMixTable[row]
		colMix := getColMix(C, col)
		M[index] = multi(rowMix, colMix)
	}
	return M
}
