package constants

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

// do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// do nothing
}
