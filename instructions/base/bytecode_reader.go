package base

type BytecodeReader struct {
	code []byte
	pc   int
}

func (self *BytecodeReader) Reset(code []byte, pc int) {
	self.code = code
	self.pc = pc
}

func (self *BytecodeReader) ReadUint8() uint8 {
	ins := self.code[self.pc]
	self.pc++
	return ins
}

func (self *BytecodeReader) ReadInt8() int8 {
	return int8(self.ReadUint8())
}

func (self *BytecodeReader) ReadUint16() uint16 {
	b1 := uint16(self.ReadUint8())
	b2 := uint16(self.ReadUint8())
	return (b1 << 8) | b2
}

func (self *BytecodeReader) ReadInt16() int16 {
	return int16(self.ReadUint16())
}

func (self *BytecodeReader) ReadInt32() int32 {
	b1 := uint32(self.ReadUint16())
	b2 := uint32(self.ReadUint16())
	return int32((b1 << 16) | b2)
}
