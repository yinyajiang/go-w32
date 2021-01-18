package wutil

import (
	"strconv"
	"strings"
	"unsafe"

	"github.com/yinyajiang/go-w32"
)

//UwpAppDesc ...
type UwpAppDesc struct {
	strWorkDir       string
	strContainerName string
	strDisplayName   string
	strPackageName   string
	strDesc          string
	strExePath       string
}

//LoadUWPDesc ...
func LoadUWPDesc(name string) (bret bool, desc UwpAppDesc) {
	defer func(n string) {
		if !bret {
			bret, desc = loadUWPDescFromReg(n)
		}
	}(name)

	if !IsHightWin10() {
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

		desc.strWorkDir = w32.UTF16PtrToString(publicAppCs.WorkingDirectory)
		desc.strContainerName = w32.UTF16PtrToString(publicAppCs.AppContainerName)
		desc.strDisplayName = w32.UTF16PtrToString(publicAppCs.DisplayName)
		desc.strPackageName = w32.UTF16PtrToString(publicAppCs.PackageFullName)
		desc.strDesc = w32.UTF16PtrToString(publicAppCs.Description)
		if strings.HasSuffix(desc.strWorkDir, "/") || strings.HasSuffix(desc.strWorkDir, `\`) {
			desc.strWorkDir = desc.strWorkDir[0 : len(desc.strWorkDir)-1]
		}
		desc.strExePath = desc.strWorkDir + `\` + desc.strDisplayName + ".exe"
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
	sys := "x" + strconv.Itoa(GetSysBit())
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
	wzBuff := w32.RegGetRaw(hRootKey, fullRegPath, "PackageID")
	if len(wzBuff) > 0 {
		desc.strPackageName = w32.UTF16ByteToString(wzBuff)
	} else {
		desc.strPackageName = subpath
	}

	wzBuff = w32.RegGetRaw(hRootKey, fullRegPath, "DisplayName")
	if len(wzBuff) > 0 {
		desc.strDisplayName = w32.UTF16ByteToString(wzBuff)
		desc.strDesc = desc.strDisplayName
	}

	wzBuff = w32.RegGetRaw(hRootKey, fullRegPath, "PackageRootFolder")
	if len(wzBuff) > 0 {
		desc.strWorkDir = w32.UTF16ByteToString(wzBuff)
	} else {
		return
	}

	if len(desc.strDisplayName) > 0 {
		desc.strExePath = desc.strWorkDir + "\\" + desc.strDisplayName + ".exe"
		if !w32.IsExist(desc.strExePath) {
			desc.strExePath = ""
		}
	}
	bRet = true
	return
}
