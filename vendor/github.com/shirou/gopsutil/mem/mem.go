package mem

import (
	"encoding/json"

	"github.com/shirou/gopsutil/internal/common"
)

var invoke common.Invoker = common.Invoke{}

// Memory usage statistics. Total, Available and Used contain numbers of bytes
// for human consumption.
//
// The other fields in this struct contain kernel specific values.
type VirtualMemoryStat struct {
	// Total amount of RAM on this system
	Total uint64 `json:"total" bson:"total,omitempty"`

	// RAM available for programs to allocate
	//
	// This value is computed from the kernel specific values.
	Available uint64 `json:"available" bson:"available,omitempty"`

	// RAM used by programs
	//
	// This value is computed from the kernel specific values.
	Used uint64 `json:"used" bson:"used,omitempty"`

	// Percentage of RAM used by programs
	//
	// This value is computed from the kernel specific values.
	UsedPercent float64 `json:"usedPercent" bson:"usedPercent,omitempty"`

	// This is the kernel's notion of free memory; RAM chips whose bits nobody
	// cares about the value of right now. For a human consumable number,
	// Available is what you really want.
	Free uint64 `json:"free" bson:"free,omitempty"`

	// OS X / BSD specific numbers:
	// http://www.macyourself.com/2010/02/17/what-is-free-wired-active-and-inactive-system-memory-ram/
	Active   uint64 `json:"active" bson:"active,omitempty"`
	Inactive uint64 `json:"inactive" bson:"inactive,omitempty"`
	Wired    uint64 `json:"wired" bson:"wired,omitempty"`

	// FreeBSD specific numbers:
	// https://reviews.freebsd.org/D8467
	Laundry uint64 `json:"laundry" bson:"laundry,omitempty"`

	// Linux specific numbers
	// https://www.centos.org/docs/5/html/5.1/Deployment_Guide/s2-proc-meminfo.html
	// https://www.kernel.org/doc/Documentation/filesystems/proc.txt
	// https://www.kernel.org/doc/Documentation/vm/overcommit-accounting
	Buffers        uint64 `json:"buffers" bson:"buffers,omitempty"`
	Cached         uint64 `json:"cached" bson:"cached,omitempty"`
	Writeback      uint64 `json:"writeback" bson:"writeback,omitempty"`
	Dirty          uint64 `json:"dirty" bson:"dirty,omitempty"`
	WritebackTmp   uint64 `json:"writebacktmp" bson:"writebacktmp,omitempty"`
	Shared         uint64 `json:"shared" bson:"shared,omitempty"`
	Slab           uint64 `json:"slab" bson:"slab,omitempty"`
	SReclaimable   uint64 `json:"sreclaimable" bson:"sreclaimable,omitempty"`
	PageTables     uint64 `json:"pagetables" bson:"pagetables,omitempty"`
	SwapCached     uint64 `json:"swapcached" bson:"swapcached,omitempty"`
	CommitLimit    uint64 `json:"commitlimit" bson:"commitlimit,omitempty"`
	CommittedAS    uint64 `json:"committedas" bson:"committedas,omitempty"`
	HighTotal      uint64 `json:"hightotal" bson:"hightotal,omitempty"`
	HighFree       uint64 `json:"highfree" bson:"highfree,omitempty"`
	LowTotal       uint64 `json:"lowtotal" bson:"lowtotal,omitempty"`
	LowFree        uint64 `json:"lowfree" bson:"lowfree,omitempty"`
	SwapTotal      uint64 `json:"swaptotal" bson:"swaptotal,omitempty"`
	SwapFree       uint64 `json:"swapfree" bson:"swapfree,omitempty"`
	Mapped         uint64 `json:"mapped" bson:"mapped,omitempty"`
	VMallocTotal   uint64 `json:"vmalloctotal" bson:"vmalloctotal,omitempty"`
	VMallocUsed    uint64 `json:"vmallocused" bson:"vmallocused,omitempty"`
	VMallocChunk   uint64 `json:"vmallocchunk" bson:"vmallocchunk,omitempty"`
	HugePagesTotal uint64 `json:"hugepagestotal" bson:"hugepagestotal,omitempty"`
	HugePagesFree  uint64 `json:"hugepagesfree" bson:"hugepagesfree,omitempty"`
	HugePageSize   uint64 `json:"hugepagesize" bson:"hugepagesize,omitempty"`
}

type SwapMemoryStat struct {
	Total       uint64  `json:"total" bson:"total,omitempty"`
	Used        uint64  `json:"used" bson:"used,omitempty"`
	Free        uint64  `json:"free" bson:"free,omitempty"`
	UsedPercent float64 `json:"usedPercent" bson:"usedPercent,omitempty"`
	Sin         uint64  `json:"sin" bson:"sin,omitempty"`
	Sout        uint64  `json:"sout" bson:"sout,omitempty"`
	PgIn        uint64  `json:"pgin" bson:"pgin,omitempty"`
	PgOut       uint64  `json:"pgout" bson:"pgout,omitempty"`
	PgFault     uint64  `json:"pgfault" bson:"pgfault,omitempty"`
}

func (m VirtualMemoryStat) String() string {
	s, _ := json.Marshal(m)
	return string(s)
}

func (m SwapMemoryStat) String() string {
	s, _ := json.Marshal(m)
	return string(s)
}
