package w32

import (
	"syscall"
	"unsafe"
)

var (
	modfirewallapi = syscall.NewLazyDLL("FirewallAPI.dll")

	procNetworkIsolationEnumAppContainers = modfirewallapi.NewProc("NetworkIsolationEnumAppContainers")
	procNetworkIsolationFreeAppContainers = modfirewallapi.NewProc("NetworkIsolationFreeAppContainers")
)

//NetworkIsolationEnumAppContainers ...
func NetworkIsolationEnumAppContainers(Flags DWORD, pdwNumPublicAppCs *DWORD, ppPublicAppCs *PINET_FIREWALL_APP_CONTAINER) DWORD {
	r, _, _ := procNetworkIsolationEnumAppContainers.Call(uintptr(Flags), uintptr(unsafe.Pointer(pdwNumPublicAppCs)), uintptr(unsafe.Pointer(ppPublicAppCs)))
	return DWORD(r)
}

//NetworkIsolationFreeAppContainers ...
func NetworkIsolationFreeAppContainers(pPublicAppCs PINET_FIREWALL_APP_CONTAINER) DWORD {
	r, _, _ := procNetworkIsolationFreeAppContainers.Call(uintptr(unsafe.Pointer(pPublicAppCs)))
	return DWORD(r)
}
