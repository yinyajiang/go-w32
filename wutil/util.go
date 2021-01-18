package wutil

import (
	"github.com/yinyajiang/go-w32"
)

//IsHightWin10 ...
func IsHightWin10() bool {
	major, _, _ := w32.RtlGetNtVersionNumbers()
	return major >= 10
}

//GetSysBit 获取当前位数
func GetSysBit() int {
	var si w32.SystemInfo
	w32.GetNativeSystemInfo(&si)
	if si.ProcessorArchitecture == 6 || si.ProcessorArchitecture == 9 {
		return 64
	}
	return 32
}
