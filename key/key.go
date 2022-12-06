package key

import "errors"

func MakeKey(s string) (key [16]byte, err error) {
	key = [16]byte{}
	if len(s) == 0 {
		return [16]byte{}, errors.New("s is zero length")
	}
	count := 0
	for {
		for i := 0; i < len(s); i++ {
			key[count] = s[i]
			count++
			if count == 16 {
				return key, nil
			}
		}
	}

}
