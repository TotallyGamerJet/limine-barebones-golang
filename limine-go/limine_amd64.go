package limine

import "unsafe"

type SMPInfo struct {
	ProcessorId   uint32
	LapicId       uint32
	_             uint64 // Reserved
	GotoAddress   GotoAddress
	ExtraArgument uint64
}

type SMPResponse struct {
	Revision   uint64
	Flags      uint32
	BspLapicId uint32
	cpuCount   uint64
	cpus       **SMPInfo
}

func (s *SMPResponse) Cpus() []*SMPInfo {
	return unsafe.Slice(s.cpus, s.cpuCount)
}
