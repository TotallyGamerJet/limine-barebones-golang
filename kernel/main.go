package main

import (
	_ "embed"
	"github.com/totallygamerjet/limine-go"
	"unsafe"
	_ "unsafe"
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
// If renaming _start() to something else, make sure to change the
// linker script accordingly. Also change the linkname pragma.
//
// Here is why each pragma below is needed:
//
// cgo_export_static - Exports the symbol for the external linker to see
//
// linkname - Go symbols are a combination of their package and func name.
// So this func is actually main._start This creates a link to this function
// just as _start so that the exported C name knows which functions to call.
//
// nosplit - Go has growable stacks. Since we haven't told the runtime how
// big the limine provided stack is disable the stack growth check.
//
//go:cgo_export_static _start _start
func _start() {
	//initializeStack(65536, main)
	//hcf()
	// Ensure we got a framebuffer.
	if framebufferRequest.Response == nil || len(framebufferRequest.Response.Framebuffers()) < 1 {
		hcf()
	}
	// Fetch the first framebuffer.
	framebuffer := framebufferRequest.Response.Framebuffers()[0]
	lfb := unsafe.Slice((*uint32)(framebuffer.Address), framebuffer.Pitch*framebuffer.Height)
	// Note: we assume the framebuffer model is RGB with 32-bit pixels.
	for i := uint64(0); i < 100; i++ {
		putPixel(lfb, int(framebuffer.Width), int(framebuffer.Pitch), int(i), int(i), 0xffffff)
	}
	initPSF(lfb, int(framebuffer.Pitch))
	drawString(lfb, int(framebuffer.Width), int(framebuffer.Pitch), "Hello World!", 0, 0, 0xFFFFFFFF, 0xFF00FFFF)
	// We're done, just hang...
	hcf()
}

func putPixel(screen []uint32, _, pitch, x, y int, color uint32) {
	where := y*(pitch/4) + x
	if where > len(screen) {
		return
	}
	screen[where] = color
}

//go:noinline
func main() {
	// Ensure we got a framebuffer.
	if framebufferRequest.Response == nil || len(framebufferRequest.Response.Framebuffers()) < 1 {
		hcf()
	}
	// Fetch the first framebuffer.
	framebuffer := framebufferRequest.Response.Framebuffers()[0]
	lfb := unsafe.Slice((*uint32)(framebuffer.Address), framebuffer.Pitch*framebuffer.Height)
	// Note: we assume the framebuffer model is RGB with 32-bit pixels.
	for i := uint64(0); i < 100; i++ {
		putPixel(lfb, int(framebuffer.Width), int(framebuffer.Pitch), int(i), int(i), 0xffffff)
	}
	initPSF(lfb, int(framebuffer.Pitch))
	drawString(lfb, int(framebuffer.Width), int(framebuffer.Pitch), "Hello World!", 0, 0, 0xFFFFFFFF, 0xFF00FFFF)
	// We're done, just hang...
	hcf()
}
