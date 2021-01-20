package wutil

import (
	"fmt"
	"os/exec"
	"strings"
	"unsafe"

	"github.com/yinyajiang/go-w32"
)

//IsAdminPrivilege 是否具有管理员权限
func IsAdminPrivilege() bool {
	ph := w32.GetCurrentProcess()
	if ph == 0 {
		return false
	}
	// Get current process token
	hToken := w32.OpenProcessToken(ph, w32.TOKEN_QUERY)
	if hToken == 0 {
		return false
	}
	defer w32.CloseHandle(hToken)
	var tokenEle w32.TOKEN_ELEVATION

	// Retrieve token elevation information
	if 0 == w32.GetTokenInformation(hToken, w32.TokenElevation, uintptr(unsafe.Pointer(&tokenEle)), uint32(unsafe.Sizeof(tokenEle))) {
		return false
	}
	return 1 == tokenEle.TokenIsElevated
}

//GetProcessToken ...
func GetProcessToken(pid uintptr) (hNewToken w32.HANDLE) {
	hProcess := w32.OpenProcess(w32.PROCESS_ALL_ACCESS, false, pid)
	if 0 == hProcess {
		return 0
	}
	defer w32.CloseHandle(hProcess)
	hToken := w32.OpenProcessToken(hProcess, w32.TOKEN_DUPLICATE)
	if 0 == hToken {
		return 0
	}
	defer w32.CloseHandle(hToken)

	hNewToken = w32.DuplicateTokenEx(hToken, w32.TOKEN_ALL_ACCESS, w32.SecurityImpersonation, w32.TokenPrimary)
	return
}

//GetGeneralPID 获取普通权限进程的PID
func GetGeneralPID() (pid uint) {
	pid = GetProcessPID("explorer.exe")
	if pid != 0 {
		return
	}

	hWnd := w32.FindWindow("Shell_TrayWnd", "")
	if 0 == hWnd {
		hWnd = w32.FindWindowEx(0, 0, "ReBarWindow32", "")
	}
	if 0 == hWnd {
		hWnd = w32.FindWindowEx(0, 0, "MSTaskSwWClass", "")
	}
	if 0 == hWnd {
		return
	}
	_, pid = w32.GetWindowThreadProcessId(hWnd)
	return
}

//DropPrivilegeStartProcess ...
func DropPrivilegeStartProcess(name string, arg ...string) (err error) {
	if IsAdminPrivilege() {
		pid := GetGeneralPID()
		if pid == 0 {
			err = fmt.Errorf("GetGeneralPID fail")
			return
		}
		token := GetProcessToken(uintptr(pid))
		if token == 0 {
			err = fmt.Errorf("GetProcessToken fail")
			return
		}
		defer w32.CloseHandle(token)
		para := `cmd start  /c "` + name + `"`
		for _, a := range arg {
			if strings.HasPrefix(a, `"`) && strings.HasSuffix(a, `"`) {
				para += ` ` + a
			} else {
				para += ` "` + a + `"`
			}
		}
		pid = StartTokenProcess(token, para)
		if pid != 0 {
			return
		}
	}

	cmd := exec.Command(name, arg...)
	if nil != cmd {
		err = cmd.Start()
	} else {
		err = fmt.Errorf("Create command fail")
	}
	return
}
