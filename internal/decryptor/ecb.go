package decryptor

import "crypto/cipher"

//reference - https://gist.github.com/DeanThompson/17056cc40b4899e3e7f4

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbDecrypter ecb

func NewDecrypterWithModeECB(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}

func (x *ecbDecrypter) BlockSize() int { return x.blockSize }

// This function should be called only after adding checks such that the panic scenarios do not occur.
// These should be handled as rzp Error. The function making call to this function should initially have
// condition checks to ensure the PANIC DOES NOT HAPPEN
// This will be abstracted to a separate package such that there will be wrapper function
// to do the checks
func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
