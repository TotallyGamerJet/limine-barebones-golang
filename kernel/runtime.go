package main

import "unsafe"

//go:linkname __g0 runtime.g0
var __g0 byte
var runtime_g0 = (*g)(unsafe.Pointer(&__g0))

type stack struct {
	lo uintptr
	hi uintptr
}

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

//go:linkname setg_gcc setg_gcc
func setg_gcc(*g)

//go:noinline
//go:nosplit
func initializeStack(sizeOfStackInBytes uint64, fn func()) {
	// MOVQ	$runtime·g0(SB), DI
	// LEAQ	(-64*1024+104)(SP), BX
	// MOVQ	BX, g_stackguard0(DI)
	// MOVQ	BX, g_stackguard1(DI)
	// MOVQ	BX, (g_stack+stack_lo)(DI)
	// MOVQ	SP, (g_stack+stack_hi)(DI)
	low := uintptr(unsafe.Add(unsafe.Pointer(&sizeOfStackInBytes), -65536)) // uintptr(sizeOfStackInBytes)
	runtime_g0.stackguard0 = low
	runtime_g0.stackguard1 = low
	runtime_g0.stack.lo = low
	runtime_g0.stack.hi = uintptr(unsafe.Pointer(&sizeOfStackInBytes))
	fn()
	hcf()
}
