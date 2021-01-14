package w32

import (
	"syscall"
	"unsafe"
)

var (
	procRtlGetNtVersionNumbers = modntdll.NewProc("RtlGetNtVersionNumbers")
)

//RtlGetNtVersionNumbers windowsAPI
func RtlGetNtVersionNumbers() (Major, Minor, BuildNumber uint32) {
	syscall.Syscall(procRtlGetNtVersionNumbers.Addr(), 3, uintptr(unsafe.Pointer(&Major)), uintptr(unsafe.Pointer(&Minor)), uintptr(unsafe.Pointer(&BuildNumber)))
	return
}
