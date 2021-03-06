package wutil

import (
	"strings"

	"github.com/yinyajiang/go-w32"
)

//LoadAppList ...
func LoadAppList() (ret map[string]string) {
	ret = make(map[string]string, 0)
	_load := func(mask uint32) {
		var rtmask uint32
		if 0 != mask&w32.KEY_WOW64_64KEY {
			rtmask |= w32.RRF_SUBKEY_WOW6464KEY
		} else if 0 != mask&w32.KEY_WOW64_32KEY {
			rtmask |= w32.RRF_SUBKEY_WOW6432KEY
		}

		path := "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Uninstall"
		hUnistall := w32.RegOpenKeyEx(w32.HKEY_LOCAL_MACHINE, path, mask)
		if hUnistall != 0 {
			defer w32.RegCloseKey(hUnistall)
		}

		for i := 0; ; i++ {
			name := w32.RegEnumKeyEx(hUnistall, uint32(i))
			if len(name) == 0 {
				break
			}
			fullpath := path + "\\" + name
			packageName := name
			if strings.HasPrefix(packageName, "{") && strings.HasSuffix(packageName, "}") {
				packageName = ""
			}

			_, wzBuff := w32.RegGetRaw(w32.HKEY_LOCAL_MACHINE, fullpath, "DisplayName", w32.RRF_RT_ANY|rtmask)
			if len(wzBuff) != 0 {
				packageName = w32.UTF16ByteToString(wzBuff)
			}

			installLocation := ""
			_, wzBuff = w32.RegGetRaw(w32.HKEY_LOCAL_MACHINE, fullpath, "InstallLocation", w32.RRF_RT_ANY|rtmask)
			if len(wzBuff) != 0 {
				installLocation = w32.UTF16ByteToString(wzBuff)
			}
			if len(packageName) == 0 {
				continue
			}
			ret[packageName] = installLocation
		}
	}
	if 32 == w32.GetSysBit() {
		_load(w32.KEY_READ)
	} else {
		_load(w32.KEY_READ | w32.KEY_WOW64_64KEY)
		_load(w32.KEY_READ | w32.KEY_WOW64_32KEY)
	}
	return
}

//IsInstalled ..
func IsInstalled(name string) bool {
	apps := LoadAppList()
	name = strings.ToLower(name)
	for k := range apps {
		k = strings.ToLower(k)
		if -1 != strings.Index(k, name) {
			return true
		}
	}
	return false
}
