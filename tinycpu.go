package main

import "fmt"

var r0, r1, r2, r3 byte
var overflow, programcounter int

func init_cpu(){
	programcounter = 0
	overflow = 0 //Not used yet
	r0 = 0
	r1 = 0
	r2 = 0
	r3 = 0
}

func selectreg(op1 byte) *byte{
	switch op1 {
	case 0x00:
		return &r0
	case 0x01:
		return &r1
	case 0x02:
		return &r2
	case 0x03:
		return &r3
	}
	return &r0
}

func exec_instruc(instructions[] int, progsize int){
	for programcounter = 0; programcounter < progsize; programcounter++ {
		instruction := instructions[programcounter]
		// Get operands
		oper1 := byte((instruction >> 16) & 0xff)
		oper2 := byte((instruction >> 8) & 0xff)
		oper3 := byte(instruction & 0xff)
		// Get opcode
		opcode := byte((instruction >> 24) & 0xff)
			
		execute_instruction(opcode,oper1,oper2,oper3);
	}
}

func execute_instruction(opcode byte, oper1 byte, oper2 byte, oper3 byte){
	oper := selectreg(oper1) 
	switch opcode {
	// ADD
	case 0x01:
		*oper = oper2 + oper3
	// SUB
	case 0x02:
		*oper = oper2 - oper3
	// MOV
	case 0x03:
		*oper = oper2
	}
}

func main(){
	init_cpu()
	instructions:= []int{
        0x01001200, // MOV R0,0x12
        0x01014376, // ADD R1,0x43,0x76
        0x02028987, // SUB R2,0x89,0x87
	}
	exec_instruc(instructions, len(instructions))
	fmt.Printf("R0 : 0x%08x\nR1 : 0x%08x\nR2 : 0x%08x\nR3 : 0x%08x\nPC : 0x%08x\n",r0,r1,r2,r3,programcounter)
}