package wutil

import (
	"github.com/yinyajiang/go-w32"
)

//IsHighWin7Sp1 ...
func IsHighWin7Sp1() bool {
	dwMajorVersion, dwMinorVersion, dwBuild := w32.GetSystemVersion()

	if dwMajorVersion > 6 {
		return true
	} else if dwMajorVersion == 6 && dwMinorVersion > 1 {
		return true
	} else if dwMajorVersion == 6 && dwMinorVersion == 1 && dwBuild > 7600 {
		return true
	}
	return false
}
