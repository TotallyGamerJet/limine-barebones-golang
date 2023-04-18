package main

import (
	"unsafe"

	"github.com/totallygamerjet/limine-barebones-golang/limine-go"
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
func hcf()

// The following will be our kernel's entry point.
// If renaming _start() to something else, make sure to change the
// linker script accordingly. Also change the linkname pragma.
//
// Here is why each pragma below is needed:
//
// linkname - Go symbols are a combination of their package and func name.
// So this func is actually github.com/totallygamerjet/limine/main._start
// This creates a link to this function just as _start
//
// noinline - Without this the Go compiler optimizes this function away
//
// nosplit - Go has growable stacks. Since we haven't set up the stack,
// yet it must be disabled.
//
//go:linkname _start _start
//go:noinline
//go:nosplit
func _start() {
	// Ensure we got a framebuffer.
	if framebufferRequest.Response == nil || len(framebufferRequest.Response.Framebuffers()) < 1 {
		hcf()
	}
	// Fetch the first framebuffer.
	framebuffer := framebufferRequest.Response.Framebuffers()[0]
	// Note: we assume the framebuffer model is RGB with 32-bit pixels.
	for i := uint64(0); i < 100; i++ {
		fbPtr := framebuffer.Address
		*(*uint32)(unsafe.Add(fbPtr, unsafe.Sizeof(uint32(0))*uintptr(i*(framebuffer.Pitch/4)+i))) = 0xffffff
	}
	// We're done, just hang...
	hcf()
}

// All Go programs expect there is be a main function
func main() {
	/* ensure that the _start function isn't optimized away */
	_start()
}