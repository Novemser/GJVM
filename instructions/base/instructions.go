package base

import "jvmgo/rtda"

// Instruction 指令接口
type Instruction interface {
	FetchOperands()
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct{}
type BranchInstruction struct {
	Offset int
}
type Index8Instruction struct {
	Index uint
}
type Index16Instruction struct {
	Index uint
}

func (self *NoOperandsInstruction) FetchOperands() {
	// do nothing
}

func (self *BranchInstruction) FetchOperands() {

}
