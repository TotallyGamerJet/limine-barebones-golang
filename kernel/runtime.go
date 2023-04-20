package main

import "unsafe"

// This is some hacky stuff in order to link to the runtime.g0 symbol.
// It's not possible to link directly to the symbol as in var g0 g
// because it complains about missing go.info stuff.
// To circumvent this, use byte and then take the address of that
// to create runtime_g0. It's perfectly acceptable for it to be a
// pointer because the Go TLS register must always store a pointer
// to this struct.
//
//go:linkname __g0 runtime.g0
var __g0 byte
var runtime_g0 = (*g)(unsafe.Pointer(&__g0))

// copied from runtime2.go
type stack struct {
	lo uintptr
	hi uintptr
}

// copied from runtime2.go
type g struct {
	// Stack parameters.
	// stack describes the actual stack memory: [stack.lo, stack.hi).
	// stackguard0 is the stack pointer compared in the Go stack growth prologue.
	// It is stack.lo+StackGuard normally, but can be StackPreempt to trigger a preemption.
	// stackguard1 is the stack pointer compared in the C stack growth prologue.
	// It is stack.lo+StackGuard on g0 and gsignal stacks.
	// It is ~0 on other goroutine stacks, to trigger a call to morestackc (and crash).
	stack       stack   // offset known to runtime/cgo
	stackguard0 uintptr // offset known to liblink
	stackguard1 uintptr // offset known to liblink
}

func setg(*g)

// initializeStack sets the runtime.g0 variable's
// stack size to sizeOfStackInBytes and then calls into main.
// main should never return but if it does this function is
// guaranteed to never return. It is important that this function
// is nosplit because before this is called the stack check will
// cause the CPU to fault since there is garabage in the TLS register.
//
//go:nosplit
func initializeStack(sizeOfStackInBytes uint64) {
	// The following Go code replicates the assembly from the rt0_go(SB)
	// function in asm_amd64.s to setup the g0 variable
	low := uintptr(unsafe.Add(unsafe.Pointer(&sizeOfStackInBytes), -sizeOfStackInBytes))
	runtime_g0.stackguard0 = low
	runtime_g0.stackguard1 = low
	runtime_g0.stack.lo = low
	runtime_g0.stack.hi = uintptr(unsafe.Pointer(&sizeOfStackInBytes))
	setg(runtime_g0)
	main()
	hcf()
}
