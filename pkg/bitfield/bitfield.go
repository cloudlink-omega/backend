package bitfield

import (
	"fmt"
)

// Bitfield interface with common bit manipulation methods

type Bitfield interface {
	Set(pos uint)
	Clear(pos uint)
	Toggle(pos uint)
	Read(pos uint) bool
	String() string
}

// Implementations for uint8, uint16, uint32, and uint64

type Bitfield8 uint8

func (b *Bitfield8) Set(pos uint) {
	*b |= 1 << pos
}

func (b *Bitfield8) Clear(pos uint) {
	*b &^= 1 << pos
}

func (b *Bitfield8) Toggle(pos uint) {
	*b ^= 1 << pos
}

func (b Bitfield8) Read(pos uint) bool {
	return b&(1<<pos) != 0
}

func (b Bitfield8) String() string {
	return fmt.Sprintf("%08b", b)
}

type Bitfield16 uint16

func (b *Bitfield16) Set(pos uint) {
	*b |= 1 << pos
}

func (b *Bitfield16) Clear(pos uint) {
	*b &^= 1 << pos
}

func (b *Bitfield16) Toggle(pos uint) {
	*b ^= 1 << pos
}

func (b Bitfield16) Read(pos uint) bool {
	return b&(1<<pos) != 0
}

func (b Bitfield16) String() string {
	return fmt.Sprintf("%016b", b)
}

type Bitfield32 uint32

func (b *Bitfield32) Set(pos uint) {
	*b |= 1 << pos
}

func (b *Bitfield32) Clear(pos uint) {
	*b &^= 1 << pos
}

func (b *Bitfield32) Toggle(pos uint) {
	*b ^= 1 << pos
}

func (b Bitfield32) Read(pos uint) bool {
	return b&(1<<pos) != 0
}

func (b Bitfield32) String() string {
	return fmt.Sprintf("%032b", b)
}

type Bitfield64 uint64

func (b *Bitfield64) Set(pos uint) {
	*b |= 1 << pos
}

func (b *Bitfield64) Clear(pos uint) {
	*b &^= 1 << pos
}

func (b *Bitfield64) Toggle(pos uint) {
	*b ^= 1 << pos
}

func (b Bitfield64) Read(pos uint) bool {
	return b&(1<<pos) != 0
}

func (b Bitfield64) String() string {
	return fmt.Sprintf("%064b", b)
}

/* func main() {
	var bf8 Bitfield8 = 0
	bf8.Set(7)
	fmt.Printf("Bitfield8: %s (%d)", bf8, bf8)
}
*/
