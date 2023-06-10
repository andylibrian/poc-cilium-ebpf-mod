package main

import (
	"fmt"

	"github.com/andylibrian/poc-cilium-ebpf-mod/pkg/ebpf/c/process_entry"
)

func main() {
	fmt.Println("Hello")

	x := process_entry.GetEbpfObject()
	fmt.Println(x.ExecveEntry.VerifierLog)
	fmt.Println(x.ExecveEntry)

	fmt.Println("Hello2")
}
