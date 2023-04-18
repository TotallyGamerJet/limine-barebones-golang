package limine

type SMPInfo struct {
	Processor_id   uint32
	Gic_iface_no   uint32
	Mpidr          uint64
	Reserved       uint64
	Goto_address   GotoAddress
	Extra_argument uint64
}

type SMPResponse struct {
	Revision  uint64
	Flags     uint32
	Bsp_mpidr uint64
	Cpu_count uint64
	Cpus      **SMPInfo
}
