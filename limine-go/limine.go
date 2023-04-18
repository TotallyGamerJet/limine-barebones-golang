// Package limine boot protocol is a modern, minimal, fast, and extensible boot protocol, with a focus on backwards
// and forwards compatibility, created from the experience gained by working on the stivale boot protocols.
package limine

import (
	"unsafe"
)

const (
	CommonMagic0, CommonMagic1 = 0xc7b1dd30df4c8b88, 0x0a82e883a194f07b
)

type UUID struct {
	A uint32
	B uint16
	C uint16
	D [8]uint8
}

const (
	MediaTypeGeneric = 0
	MediaTypeOptical = 1
	MediaTypeTftp    = 2
	FramebufferRgb   = 1
)

type File struct {
	Revision       uint64
	Address        unsafe.Pointer
	Size           uint64
	path           *byte
	cmdline        *byte
	MediaType      uint32
	_              uint32 // Unused
	TftpIp         uint32
	TftpPort       uint32
	PartitionIndex uint32
	MbrDiskId      uint32
	GptDiskUuid    UUID
	GptPartUuid    UUID
	PartUuid       UUID
}

func (f *File) Data() []byte {
	return unsafe.Slice((*byte)(f.Address), f.Size)
}

func (f *File) Path() string {
	return unsafe.String(f.path, findNull(f.path))
}

func (f *File) Cmdline() string {
	return unsafe.String(f.cmdline, findNull(f.cmdline))
}

// Boot Info

const (
	BootloaderInfoRequestId2 = 0xf55038d8e2a1202f
	BootloaderInfoRequestId3 = 0x279426fcf5f59740
)

type BootloaderInfoResponse struct {
	Revision uint64
	name     *byte
	version  *byte
}

func (b *BootloaderInfoResponse) Name() string {
	return unsafe.String(b.name, findNull(b.name))
}

func (b *BootloaderInfoResponse) Version() string {
	return unsafe.String(b.version, findNull(b.version))
}

type BootloaderInfoRequest struct {
	Id       [4]uint64
	Revision uint64
	Response *BootloaderInfoResponse
}

// Stack size

const (
	StackSizeRequestId2 = 0x224ef0460a8e8926
	StackSizeRequestId3 = 0xe1cb0fc25f46ea3d
)

type StackSizeResponse struct {
	Revision uint64
}

type StackSizeRequest struct {
	Id        [4]uint64
	Revision  uint64
	Response  *StackSizeResponse
	StackSize uint64
}

// HHDM

const (
	HHDMRequestId2 = 0x48dcf1cb8ad2b852
	HHDMRequestId3 = 0x63984e959a98244b
)

type HHDMResponse struct {
	Revision uint64
	Offset   uint64
}

type HHDMRequest struct {
	Id       [4]uint64
	Revision uint64
	Response *HHDMResponse
}

// Framebuffer

const FrameBufferRGB byte = 1

const (
	FramebufferRequestId2 = 0x9d5827dcd881dd75
	FramebufferRequestId3 = 0xa3148604f6fab11b
)

type VideoMode struct {
	Pitch          uint64
	Width          uint64
	Height         uint64
	Bpp            uint16
	MemoryModel    uint8
	RedMaskSize    uint8
	RedMaskShift   uint8
	GreenMaskSize  uint8
	GreenMaskShift uint8
	BlueMaskSize   uint8
	BlueMaskShift  uint8
}

type Framebuffer struct {
	Address        unsafe.Pointer
	Width          uint64
	Height         uint64
	Pitch          uint64
	Bpp            uint16
	MemoryModel    uint8
	RedMaskSize    uint8
	RedMaskShift   uint8
	GreenMaskSize  uint8
	GreenMaskShift uint8
	BlueMaskSize   uint8
	BlueMaskShift  uint8
	_              [7]uint8
	EdidSize       uint64
	Edid           unsafe.Pointer
	modeCount      uint64
	modes          **VideoMode
}

func (f *Framebuffer) Data() []byte {
	return unsafe.Slice((*byte)(f.Address), f.Pitch*f.Height)
}

func (f *Framebuffer) EdidData() []byte {
	if f.EdidSize == 0 {
		return nil
	}
	return unsafe.Slice((*byte)(f.Edid), f.EdidSize)
}

func (f *Framebuffer) Modes() []*VideoMode {
	return unsafe.Slice(f.modes, f.modeCount)
}

type FramebufferResponse struct {
	Revision         uint64
	framebufferCount uint64
	framebuffers     **Framebuffer
}

func (f *FramebufferResponse) Framebuffers() []*Framebuffer {
	return unsafe.Slice(f.framebuffers, f.framebufferCount)
}

type FramebufferRequest struct {
	Id       [4]uint64
	Revision uint64
	Response *FramebufferResponse
}

// Terminal

const (
	TerminalRequestId2 = 0xc8ac59310c2b0844 // Deprecated: do not use if possible.
	TerminalRequestId3 = 0xa68d0c7265d38878 // Deprecated: do not use if possible.
)

const (
	TerminalCbDec          = 10 // Deprecated: do not use if possible.
	TerminalCbBell         = 20 // Deprecated: do not use if possible.
	TerminalCbPrivateId    = 30 // Deprecated: do not use if possible.
	TerminalCbStatusReport = 40 // Deprecated: do not use if possible.
	TerminalCbPosReport    = 50 // Deprecated: do not use if possible.
	TerminalCbKbdLeds      = 60 // Deprecated: do not use if possible.
	TerminalCbMode         = 70 // Deprecated: do not use if possible.
	TerminalCbLinux        = 80 // Deprecated: do not use if possible.
)

const (
	// Deprecated: do not use if possible.
	TerminalCtxSize = ^uint64(1 - 1) // -1
	// Deprecated: do not use if possible.
	TerminalCtxSave = ^uint64(2 - 1) // -2
	// Deprecated: do not use if possible.
	TerminalCtxRestore = ^uint64(3 - 1) // -3
	// Deprecated: do not use if possible.
	TerminalFullRefresh = ^uint64(4 - 1) // -4
)

// Response revision 1
const (
	// Deprecated: do not use if possible.
	TerminalOobOutputGet = ^uint64(10 - 1) // -10
	// Deprecated: do not use if possible.
	TerminalOobOutputSet = ^uint64(11 - 1) // -11
)

const (
	TerminalOobOutputOcrnl  = 1 << iota // Deprecated: do not use if possible.
	TerminalOobOutputOfdel              // Deprecated: do not use if possible.
	TerminalOobOutputOfill              // Deprecated: do not use if possible.
	TerminalOobOutputOlcuc              // Deprecated: do not use if possible.
	TerminalOobOutputOnlcr              // Deprecated: do not use if possible.
	TerminalOobOutputOnlret             // Deprecated: do not use if possible.
	TerminalOobOutputOnocr              // Deprecated: do not use if possible.
	TerminalOobOutputOpost              // Deprecated: do not use if possible.
)

// Deprecated: do not use if possible.
type TerminalWrite uintptr

// Deprecated: do not use if possible.
func (w TerminalWrite) Call(*Terminal, *byte, uint64) {
	// TODO:
}

// Deprecated: do not use if possible.
type TerminalCallback uintptr

// Deprecated: do not use if possible.
func (c TerminalCallback) Call(*Terminal, uint64, uint64, uint64, uint64) {
	// TODO:
}

// Deprecated: do not use if possible.
type Terminal struct {
	Columns     uint64
	Rows        uint64
	Framebuffer *Framebuffer
}

// Deprecated: do not use if possible.
type TerminalResponse struct {
	Revision      uint64
	terminalCount uint64
	terminals     **Terminal
	Write         TerminalWrite
}

// Deprecated: do not use if possible.
func (t *TerminalResponse) Terminals() []*Terminal {
	return unsafe.Slice(t.terminals, t.terminalCount)
}

// Deprecated: do not use if possible.
type TerminalRequest struct {
	Id       [4]uint64
	Revision uint64
	Response *TerminalResponse
	Callback TerminalCallback
}

// 5-level paging

const (
	FiveLevelPagingRequestId2 = 0x94469551da9b3192
	FiveLevelPagingRequestId3 = 0xebe5e86db7382888
)

type FiveLevelPagingResponse struct {
	Revision uint64
}

type FiveLevelPagingRequest struct {
	Id       [4]uint64
	Revision uint64
	Response *FiveLevelPagingResponse
}

// SMP

const (
	SmpX2apic = 1
)

const (
	SMPRequestId2 = 0x95a67b819a1b857e
	SMPRequestId3 = 0xa0b61b723b6a73e0
)

type GotoAddress uintptr

func (g GotoAddress) Call(*SMPInfo) {
	// TODO:
}

type SMPRequest struct {
	Id       [4]uint64
	Revision uint64
	Response *SMPResponse
	Flags    uint64
}

// Memory Map

const (
	MemmapUsable                = 0
	MemmapReserved              = 1
	MemmapAcpiReclaimable       = 2
	MemmapAcpiNvs               = 3
	MemmapBadMemory             = 4
	MemmapBootloaderReclaimable = 5
	MemmapKernelAndModules      = 6
	MemmapFramebuffer           = 7
)

const (
	MemMapRequestId2 = 0x67cf3d9d378a806f
	MemMapRequestId3 = 0xe304acdfc50c3c62
)

type MemMapEntry struct {
	Base   uint64
	Length uint64
	Type   uint64
}

type MemMapResponse struct {
	Revision   uint64
	entryCount uint64
	entries    **MemMapEntry
}

func (m *MemMapResponse) Entries() []*MemMapEntry {
	return unsafe.Slice(m.entries, m.entryCount)
}

type MemMapRequest struct {
	Id       [4]uint64
	Revision uint64
	Response *MemMapResponse
}

// Entry Point

const (
	EntryPointRequestId2 = 0x13d86c035a1cd3e1
	EntryPointRequestId3 = 0x2b0caa89d8f3026a
)

type EntryPoint uintptr

func (e EntryPoint) Call() {
	// TODO:
}

type EntryPointResponse struct {
	Revision uint64
}

type EntryPointRequest struct {
	Id       [4]uint64
	Revision uint64
	Response *EntryPointResponse
	Entry    EntryPoint
}

// Kernel File

const (
	KernelFileRequestId2 = 0xad97e90e83f1ed67
	KernelFileRequestId3 = 0x31eb5d1c5ff23b69
)

type KernelFileResponse struct {
	Revision   uint64
	KernelFile *File
}

type KernelFileRequest struct {
	Id       [4]uint64
	Revision uint64
	Response *KernelFileResponse
}

// Module

const (
	InternalModuleRequired = 1
)

const (
	ModuleRequestId2 = 0x3e7e279702be32af
	ModuleRequestId3 = 0xca1c4f3bd1280cee
)

type InternalModule struct {
	Path    *byte
	Cmdline *byte
	Flags   uint64
}

type ModuleResponse struct {
	Revision    uint64
	moduleCount uint64
	modules     **File
}

func (m *ModuleResponse) Modules() []*File {
	return unsafe.Slice(m.modules, m.moduleCount)
}

type ModuleRequest struct {
	Id                  [4]uint64
	Revision            uint64
	Response            *ModuleResponse
	internalModuleCount uint64
	internalModules     **InternalModule
}

func (m *ModuleRequest) InternalModules() []*InternalModule {
	return unsafe.Slice(m.internalModules, m.internalModuleCount)
}

// RSDP

const (
	RSDPRequestId2 = 0xc5e77b6b397e7b43
	RSDPRequestId3 = 0x27637845accdcf3c
)

type RSDPResponse struct {
	Revision uint64
	Address  unsafe.Pointer
}

type RSDPRequest struct {
	Id       [4]uint64
	Revision uint64
	Response *RSDPResponse
}

// SMBIOS

const (
	SMBIOSRequestId2 = 0x9e9046f11e095391
	SMBIOSRequestId3 = 0xaa4a520fefbde5ee
)

type SMBIOSResponse struct {
	Revision uint64
	Entry32  unsafe.Pointer
	Entry64  unsafe.Pointer
}

type SMBIOSsRequest struct {
	Id       [4]uint64
	Revision uint64
	Response *SMBIOSResponse
}

// EFI System Table

const (
	EfiSystemTableRequestId2 = 0x5ceba5163eaaf6d6
	EfiSystemTableRequestId3 = 0x0a6981610cf65fcc
)

type EfiSystemTableResponse struct {
	Revision uint64
	Address  unsafe.Pointer
}

type EfiSystemTableRequest struct {
	Id       [4]uint64
	Revision uint64
	Response *EfiSystemTableResponse
}

// Boot time

const (
	BootTimeRequestId2 = 0x502746e184c088aa
	BootTimeRequestId3 = 0xfbc5ec83e6327893
)

type BootTimeResponse struct {
	Revision uint64
	BootTime int64
}

type BootTimeRequest struct {
	Id       [4]uint64
	Revision uint64
	Response *BootTimeResponse
}

// Kernel address

const (
	KernelAddressRequestId2 = 0x71ba76863cc55f63
	KernelAddressRequestId3 = 0xb2644a48c516a487
)

type KernelAddressResponse struct {
	Revision     uint64
	PhysicalBase uint64
	VirtualBase  uint64
}

type KernelAddressRequest struct {
	Id       [4]uint64
	Revision uint64
	Response *KernelAddressResponse
}

// Device Tree Blob

const (
	DTBRequestId2 = 0xb40ddb48fb54bac7
	DTBRequestId3 = 0x545081493f81ffb7
)

type DTBResponse struct {
	Revision uint64
	DtbPtr   unsafe.Pointer
}

type DTBRequest struct {
	Id       [4]uint64
	Revision uint64
	Response *DTBResponse
}

// findNull returns the length of the null-terminated C string
func findNull(b *byte) uint {
	ptr := *(*unsafe.Pointer)(unsafe.Pointer(&b))
	if ptr == nil {
		return 0
	}
	var length uint
	for {
		if *(*byte)(unsafe.Add(ptr, uintptr(length))) == '\x00' {
			break
		}
		length++
	}
	return length
}
