package mtproto

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"math/big"
)

var (
	__debug int32
)

const (
	DEBUG_LEVEL_NETWORK         = 0x01
	DEBUG_LEVEL_NETWORK_DETAILS = 0x02
	DEBUG_LEVEL_DECODE          = 0x04
	DEBUG_LEVEL_DECODE_DETAILS  = 0x08
)

type DecodeBuf struct {
	buf  []byte
	off  int
	size int
	err  error
}

func NewDecodeBuf(b []byte) *DecodeBuf {
	return &DecodeBuf{b, 0, len(b), nil}
}

func (m *DecodeBuf) Long() int64 {
	if m.err != nil {
		return 0
	}
	if m.off+8 > m.size {
		m.err = errors.New("DecodeLong")
		return 0
	}
	x := int64(binary.LittleEndian.Uint64(m.buf[m.off : m.off+8]))
	m.off += 8
	if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
		fmt.Println("Decode::Long::", x)
	}
	return x
}

func (m *DecodeBuf) Double() float64 {
	if m.err != nil {
		return 0
	}
	if m.off+8 > m.size {
		m.err = errors.New("DecodeDouble")
		return 0
	}
	x := math.Float64frombits(binary.LittleEndian.Uint64(m.buf[m.off : m.off+8]))
	m.off += 8
	if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
		fmt.Println("Decode::Double::", x)
	}
	return x
}

func (m *DecodeBuf) Int() int32 {
	if m.err != nil {
		return 0
	}
	if m.off+4 > m.size {
		m.err = errors.New("DecodeInt")
		return 0
	}
	x := binary.LittleEndian.Uint32(m.buf[m.off : m.off+4])
	m.off += 4
	if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
		fmt.Println("Decode::Int::", x)
	}
	return int32(x)
}

func (m *DecodeBuf) UInt() uint32 {
	if m.err != nil {
		return 0
	}
	if m.off+4 > m.size {
		m.err = errors.New("DecodeUInt")
		return 0
	}
	x := binary.LittleEndian.Uint32(m.buf[m.off : m.off+4])
	m.off += 4
	if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
		fmt.Println(fmt.Sprintf("Decode::UInt::%x", x))
	}
	return x
}

func (m *DecodeBuf) Bytes(size int) []byte {
	if m.err != nil {
		return nil
	}
	if m.off+size > m.size {
		m.err = errors.New("DecodeBytes")
		return nil
	}
	x := make([]byte, size)
	copy(x, m.buf[m.off:m.off+size])
	m.off += size
	if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
		if len(x) > 10 {
			fmt.Println("Decode::Bytes::", len(x), x[:10], " ...")
		} else {
			fmt.Println("Decode::Bytes::", len(x), x)
		}

	}
	return x
}

func (m *DecodeBuf) StringBytes() []byte {
	if m.err != nil {
		return nil
	}
	var size, padding int

	if m.off+1 > m.size {
		m.err = errors.New("DecodeStringBytes")
		return nil
	}
	size = int(m.buf[m.off])
	m.off++
	padding = (4 - ((size + 1) % 4)) & 3
	if size == 254 {
		if m.off+3 > m.size {
			m.err = errors.New("DecodeStringBytes")
			return nil
		}
		size = int(m.buf[m.off]) | int(m.buf[m.off+1])<<8 | int(m.buf[m.off+2])<<16
		m.off += 3
		padding = (4 - size%4) & 3
	}

	if m.off+size > m.size {
		m.err = errors.New("DecodeStringBytes: Wrong size")
		return nil
	}
	x := make([]byte, size)
	copy(x, m.buf[m.off:m.off+size])
	m.off += size

	if m.off+padding > m.size {
		m.err = errors.New("DecodeStringBytes: Wrong padding")
		return nil
	}
	m.off += padding
	if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
		if len(x) > 10 {
			fmt.Println("Decode::StringBytes::", len(x), x[:10], " ...")
		} else {
			fmt.Println("Decode::StringBytes::", len(x), x)
		}

	}
	return x
}

func (m *DecodeBuf) String() string {
	b := m.StringBytes()
	if m.err != nil {
		return ""
	}
	x := string(b)
	if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
		fmt.Println("Decode::String::", x)
	}
	return x
}

func (m *DecodeBuf) BigInt() *big.Int {
	b := m.StringBytes()
	if m.err != nil {
		return nil
	}
	y := make([]byte, len(b)+1)
	y[0] = 0
	copy(y[1:], b)
	x := new(big.Int).SetBytes(y)
	if __debug&DEBUG_LEVEL_DECODE_DETAILS != 0 {
		fmt.Println("Decode::BigInt::", x)
	}
	return x
}

func (d *DecodeBuf) dump() {
	fmt.Println(hex.Dump(d.buf[d.off:d.size]))
}
