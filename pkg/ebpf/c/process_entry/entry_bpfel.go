// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || amd64p32 || arm || arm64 || mips64le || mips64p32le || mipsle || ppc64le || riscv64
// +build 386 amd64 amd64p32 arm arm64 mips64le mips64p32le mipsle ppc64le riscv64

package process_entry

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type entryEventData struct {
	Pid            uint32
	Tgid           uint32
	Uid            uint32
	Gid            uint32
	SyscallNr      int32
	Comm           [16]uint8
	Cwd            [32]uint8
	BinaryFilepath [256]uint8
	UserComm       [256][256]uint8
}

// loadEntry returns the embedded CollectionSpec for entry.
func loadEntry() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_EntryBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load entry: %w", err)
	}

	return spec, err
}

// loadEntryObjects loads entry and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*entryObjects
//	*entryPrograms
//	*entryMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadEntryObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadEntry()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// entrySpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type entrySpecs struct {
	entryProgramSpecs
	entryMapSpecs
}

// entrySpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type entryProgramSpecs struct {
	ExecveEntry *ebpf.ProgramSpec `ebpf:"execve_entry"`
}

// entryMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type entryMapSpecs struct {
	Event *ebpf.MapSpec `ebpf:"event"`
}

// entryObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadEntryObjects or ebpf.CollectionSpec.LoadAndAssign.
type entryObjects struct {
	entryPrograms
	entryMaps
}

func (o *entryObjects) Close() error {
	return _EntryClose(
		&o.entryPrograms,
		&o.entryMaps,
	)
}

// entryMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadEntryObjects or ebpf.CollectionSpec.LoadAndAssign.
type entryMaps struct {
	Event *ebpf.Map `ebpf:"event"`
}

func (m *entryMaps) Close() error {
	return _EntryClose(
		m.Event,
	)
}

// entryPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadEntryObjects or ebpf.CollectionSpec.LoadAndAssign.
type entryPrograms struct {
	ExecveEntry *ebpf.Program `ebpf:"execve_entry"`
}

func (p *entryPrograms) Close() error {
	return _EntryClose(
		p.ExecveEntry,
	)
}

func _EntryClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed entry_bpfel.o
var _EntryBytes []byte
