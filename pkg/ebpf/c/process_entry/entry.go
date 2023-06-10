package process_entry

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc clang -cflags $BPF_CFLAGS -type event_data entry entry.bpf.c -- -I../../../../headers
func GetEbpfObject() *entryObjects {
	var bpfObj entryObjects
	err := loadEntryObjects(&bpfObj, nil)
	if err != nil {
		panic(err)
	}

	return &bpfObj
}
