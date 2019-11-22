package crypto

import "golang.org/x/crypto/ripemd160"

func TestRipemd160Size() int {
	return ripemd160.Size
}
