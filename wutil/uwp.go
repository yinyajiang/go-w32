package wutil

import (
	"strconv"
	"strings"
	"unsafe"

	"github.com/yinyajiang/go-w32"
)

//UwpAppDesc ...
type UwpAppDesc struct {
	WorkDir       string
	ContainerName string
	DisplayName   string
	PackageName   string
	Desc          string
	ExePath       string
}

//LoadUWPDesc ...
func LoadUWPDesc(name string) (bret bool, desc UwpAppDesc) {
	defer func(n string) {
		if !bret {
			bret, desc = loadUWPDescFromReg(n)
		}
	}(name)

	if !w32.IsHightWin10() {
		return
	}

	var dwNumPublicAppCs w32.DWORD
	var pPublicAppCs w32.PINET_FIREWALL_APP_CONTAINER
	w32.NetworkIsolationEnumAppContainers(0, &dwNumPublicAppCs, &pPublicAppCs)
	if pPublicAppCs != nil {
		defer w32.NetworkIsolationFreeAppContainers(pPublicAppCs)
	}

	appCsIndex := func(pPublicAppCs w32.PINET_FIREWALL_APP_CONTAINER, index int) w32.PINET_FIREWALL_APP_CONTAINER {
		return w32.PINET_FIREWALL_APP_CONTAINER(w32.PrtIndex(pPublicAppCs, index))
	}

	for i := 0; i < int(dwNumPublicAppCs); i++ {
		publicAppCs := appCsIndex(pPublicAppCs, i)
		if nil == unsafe.Pointer((publicAppCs.AppContainerName)) {
			continue
		}

		tmpName := w32.UTF16PtrToString(publicAppCs.AppContainerName)
		if -1 == strings.Index(tmpName, name) {
			continue
		}

		desc.WorkDir = w32.UTF16PtrToString(publicAppCs.WorkingDirectory)
		desc.ContainerName = w32.UTF16PtrToString(publicAppCs.AppContainerName)
		desc.DisplayName = w32.UTF16PtrToString(publicAppCs.DisplayName)
		desc.PackageName = w32.UTF16PtrToString(publicAppCs.PackageFullName)
		desc.Desc = w32.UTF16PtrToString(publicAppCs.Description)
		if strings.HasSuffix(desc.WorkDir, "/") || strings.HasSuffix(desc.WorkDir, `\`) {
			desc.WorkDir = desc.WorkDir[0 : len(desc.WorkDir)-1]
		}
		desc.ExePath = desc.WorkDir + `\` + desc.DisplayName + ".exe"
		bret = true
		break
	}
	return
}

func loadUWPDescFromReg(name string) (bRet bool, desc UwpAppDesc) {
	if len(name) == 0 {
		return
	}
	hRootKey := w32.HKEY_CURRENT_USER
	strRegPath := "SOFTWARE\\Classes\\Local Settings\\Software\\Microsoft\\Windows\\CurrentVersion\\AppModel\\Repository\\Packages"
	hMainKey := w32.RegOpenKeyEx(hRootKey, strRegPath, w32.KEY_READ)
	if 0 == hMainKey {
		return
	}
	defer w32.RegCloseKey(hMainKey)
	sys := "x" + strconv.Itoa(w32.GetSysBit())
	subpath := ""
	for dwIndex := 0; ; dwIndex++ {
		val := w32.RegEnumKeyEx(hMainKey, uint32(dwIndex))
		if len(val) == 0 {
			break
		}
		if -1 == strings.Index(val, name) {
			continue
		}
		if -1 == strings.Index(val, sys) {
			continue
		}
		subpath = val
		break
	}
	if len(subpath) == 0 {
		return
	}

	fullRegPath := strRegPath + "\\" + subpath
	_, wzBuff := w32.RegGetRawAll(hRootKey, fullRegPath, "PackageID")
	if len(wzBuff) > 0 {
		desc.PackageName = w32.UTF16ByteToString(wzBuff)
	} else {
		desc.PackageName = subpath
	}

	_, wzBuff = w32.RegGetRawAll(hRootKey, fullRegPath, "DisplayName")
	if len(wzBuff) > 0 {
		desc.DisplayName = w32.UTF16ByteToString(wzBuff)
		desc.Desc = desc.DisplayName
	}

	_, wzBuff = w32.RegGetRawAll(hRootKey, fullRegPath, "PackageRootFolder")
	if len(wzBuff) > 0 {
		desc.WorkDir = w32.UTF16ByteToString(wzBuff)
	} else {
		return
	}

	if len(desc.DisplayName) > 0 {
		desc.ExePath = desc.WorkDir + "\\" + desc.DisplayName + ".exe"
		if !w32.IsExist(desc.ExePath) {
			desc.ExePath = ""
		}
	}
	bRet = true
	return
}
