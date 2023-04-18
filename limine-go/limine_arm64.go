package limine

import "unsafe"

type SMPInfo struct {
	ProcessorId   uint32
	GicIfaceNo    uint32
	Mpidr         uint64
	_             uint64 // Reserved
	GotoAddress   GotoAddress
	ExtraArgument uint64
}

type SMPResponse struct {
	Revision uint64
	Flags    uint32
	BspMpidr uint64
	cpuCount uint64
	cpus     **SMPInfo
}

func (s *SMPResponse) Cpus() []*SMPInfo {
	return unsafe.Slice(s.cpus, s.cpuCount)
}
