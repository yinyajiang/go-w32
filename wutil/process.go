package wutil

import (
	"strings"
	"syscall"
	"unsafe"

	"github.com/yinyajiang/go-w32"
)

//IsProcessRuning ...
func IsProcessRuning(names ...string) (bret bool) {

	ProcessWalk(func(e32 w32.PROCESSENTRY32) bool {
		exefile := strings.ToLower(w32.UTF16SpliceToString(e32.SzExeFile[:]))
		for _, name := range names {
			name = strings.ToLower(name)
			if -1 != strings.Index(exefile, name) {
				bret = true
				return false
			}
		}
		return true
	})
	return
}

//GetProcessPID ...
func GetProcessPID(name string) (ret uint) {
	name = strings.ToLower(name)
	ProcessWalk(func(e32 w32.PROCESSENTRY32) bool {
		exefile := strings.ToLower(w32.UTF16SpliceToString(e32.SzExeFile[:]))
		if -1 != strings.Index(exefile, name) {
			ret = uint(e32.Th32ProcessID)
			return false
		}
		return true
	})
	return
}

//StartAdminProcess UAC启动
func StartAdminProcess(path string, arg []string) (pid uint) {
	path = strings.ReplaceAll(path, "/", "\\")
	para := ""
	if arg != nil {
		para = strings.Join(arg, " ")
	}

	var shExecInfo w32.SHELLEXECUTEINFOW
	shExecInfo.CbSize = w32.DWORD(unsafe.Sizeof(shExecInfo))
	shExecInfo.FMask = w32.SEE_MASK_NOCLOSEPROCESS
	shExecInfo.Hwnd = 0
	shExecInfo.LpVerb = syscall.StringToUTF16Ptr("runas")
	shExecInfo.LpFile = syscall.StringToUTF16Ptr(path)
	shExecInfo.LpParameters = syscall.StringToUTF16Ptr(para)
	shExecInfo.LpDirectory = nil
	shExecInfo.LpClass = nil
	shExecInfo.NShow = w32.SW_HIDE
	shExecInfo.HInstApp = 0
	if w32.ShellExecuteEx(&shExecInfo) {
		pid = w32.GetProcessId(shExecInfo.HProcess)
	}
	return
}

//StartTokenProcess ...
func StartTokenProcess(hPtoken w32.HANDLE, cmd string) (pid uint) {
	if 0 == hPtoken {
		return 0
	}

	var si w32.STARTUPINFOW
	si.Cb = uint32(unsafe.Sizeof(si))
	si.Flags = w32.STARTF_USESHOWWINDOW
	si.ShowWindow = w32.SW_HIDE
	var pi w32.PROCESS_INFORMATION
	if w32.CreateProcessWithToken(hPtoken, w32.LOGON_WITH_PROFILE, cmd, &si, &pi) {
		pid = w32.GetProcessId(pi.Process)
	}
	return
}

//ProcessWalk 遍历进程表
func ProcessWalk(fun func(w32.PROCESSENTRY32) bool) {
	hToolhelp := w32.CreateToolhelp32Snapshot(w32.TH32CS_SNAPPROCESS, 0)
	if 0 != hToolhelp {
		defer w32.CloseHandle(hToolhelp)
	}

	b, e32 := w32.Process32First(hToolhelp)
	if !b {
		return
	}
	if !fun(e32) {
		return
	}

	for b {
		b, e32 = w32.Process32Next(hToolhelp)
		if b && !fun(e32) {
			return
		}
	}
}
