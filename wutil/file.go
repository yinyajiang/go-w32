package wutil

import (
	"fmt"

	"github.com/yinyajiang/go-w32"
)

//GetFileVersion 获取文件版本信息
func GetFileVersion(file string) string {
	data := w32.GetFileVersionInfo(file)
	if len(data) == 0 {
		return ""
	}
	fileInfo := w32.VerQueryValue(data)

	ver := fmt.Sprintf("%d.%d.%d.%d", fileInfo.FileVersionMS>>16,
		fileInfo.FileVersionMS&0xffff,
		fileInfo.FileVersionLS>>16,
		fileInfo.FileVersionLS&0xffff)
	return ver
}
