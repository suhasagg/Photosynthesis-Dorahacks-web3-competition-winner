// Code generated by command: go run asm.go -out=fp_amd64.s -stubs=fp_stub.go. DO NOT EDIT.

// +build amd64,!purego

package farm

func Fingerprint64(s []byte) uint64

func Fingerprint32(s []byte) uint32
