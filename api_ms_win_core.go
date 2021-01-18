package w32

import (
	// "github.com/davecgh/go-spew/spew"
	"fmt"
	"syscall"
	"unsafe"
)

var (
	modmswincore = syscall.NewLazyDLL("Api-ms-win-core-version-l1-1-0.dll")

	procVerQueryValueA      = modmswincore.NewProc("VerQueryValueA")
	procGetFileVersionInfoW = modmswincore.NewProc("GetFileVersionInfoW")
)

//VerQueryValue windowsAPI
func VerQueryValue(data []byte) *VsFIXEDFILEINFO {
	var info *VsFIXEDFILEINFO
	var bytes uint32
	r, _, err := syscall.Syscall6(procVerQueryValueA.Addr(), 4, SpliceToPtr(data), StringToUTF16Ptr("\\"), uintptr(unsafe.Pointer(&info)), uintptr(unsafe.Pointer(&bytes)), 0, 0)
	if r != 1 {
		fmt.Println(err)
		return nil
	}
	return info
}

//GetFileVersionInfo windowsapi
func GetFileVersionInfo(file string) []byte {
	filep, _ := syscall.UTF16PtrFromString(file)
	buff := make([]byte, 2048)
	r, _, err := syscall.Syscall6(procGetFileVersionInfoW.Addr(), 4, uintptr(unsafe.Pointer(filep)), 0, 2048, uintptr(unsafe.Pointer(&buff[0])), 0, 0)
	if r != 1 {
		fmt.Println(err)
		return []byte{}
	}
	return buff
}
