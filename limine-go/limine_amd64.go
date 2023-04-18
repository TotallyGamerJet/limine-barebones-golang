package limine

type SMPInfo struct {
	Processor_id   uint32
	Lapic_id       uint32
	Reserved       uint64
	Goto_address   GotoAddress
	Extra_argument uint64
}

type SMPResponse struct {
	Revision     uint64
	Flags        uint32
	Bsp_lapic_id uint32
	Cpu_count    uint64
	Cpus         **SMPInfo
}
