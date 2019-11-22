package test1

import (
	"fmt"
	"github.com/abesuite/test1/crypto"
)

func TestDepend() {
	var i = crypto.TestRipemd160Size()
	fmt.Print(i)
}
