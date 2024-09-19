package main

import (
	"unsafe"

	"github.com/totallygamerjet/limine-go"
)

// The Limine requests can be placed anywhere, but it is important that
// the compiler does not optimise them away, so, usually, they should
// be made volatile or equivalent.
var framebufferRequest = limine.FramebufferRequest{
	Id: [...]uint64{
		limine.CommonMagic0,
		limine.CommonMagic1,
		limine.FramebufferRequestId2,
		limine.FramebufferRequestId3,
	},
	Revision: 0,
}

// Halt and catch fire function.
func hcf() {
	for {
	}
}

// The following will be our kernel's entry point.
// If renaming _start to something else, make sure to change the
// assembly accordingly.
//
// This variable is necessary because Go has two ABIs: ABI0 and ABIInternal.
// ABIInternal is what we want but the linker links to ABI0 and assembly can't
// call ABIInternal directly unless its in the runtime. So the workaround is
// to create a function pointer and derefence it in the assembly.
var _start = start

// The reason the below pragma is needed:
//
// nosplit - Go has growable stacks. Since we haven't told the runtime how
// big the limine provided stack is disable the stack growth check.
//
//go:nosplit
func start() {
	// Ensure we got a framebuffer.
	if framebufferRequest.Response == nil || len(framebufferRequest.Response.Framebuffers()) < 1 {
		hcf()
	}
	// Fetch the first framebuffer.
	framebuffer := framebufferRequest.Response.Framebuffers()[0]
	// Note: we assume the framebuffer model is RGB with 32-bit pixels.
	for i := uint64(0); i < 100; i++ {
		pixOffset := uintptr(i*(framebuffer.Pitch/4) + i)
		*(*uint32)(unsafe.Add(framebuffer.Address, pixOffset*unsafe.Sizeof(uint32(0)))) = 0xffffff
	}
	// We're done, just hang...
	hcf()
}

// All Go programs expect there to be a main function even though this is never called
func main() {}
