// Code generated by github.com/go-enjin/golang-org-x-text/internal/gen/bitfield. DO NOT EDIT.

package bitfield

type myInt uint32

func (m myInt) fob() uint16 {
	return uint16((m >> 16) & 0xffff)
}

func (m myInt) baz() int8 {
	return int8((m >> 11) & 0x1f)
}

func (m myInt) bar() myUint8 {
	return myUint8((m >> 8) & 0x7)
}

func (m myInt) Bool() bool {
	const bit = 1 << 7
	return m&bit == bit
}

func (m myInt) Baz() int8 {
	return int8((m >> 4) & 0x7)
}
