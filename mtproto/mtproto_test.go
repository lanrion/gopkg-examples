package mtproto

import (
	"fmt"
	"testing"
)

func TestNewEncodeBuf(t *testing.T) {
	x := NewEncodeBuf(20)
	x.UInt(0x60469778)
	b := []byte{1, 2, 3, 4, 5}
	x.Bytes(b)
	fmt.Println(x.buf)

	deb := NewDecodeBuf(x.buf)
	deb.Bytes(20)
	fmt.Println("123: ", deb.String())
}
