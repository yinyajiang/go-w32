package w32

import (
	"os"
	"path/filepath"
	"strings"
)

//AbsPath 绝对路径
func AbsPath(path string) string {
	return AbsJoinPath(path)
}

//AbsJoinPath 拼接路径
func AbsJoinPath(paths ...string) string {
	if 0 == len(paths) {
		return ""
	}
	abs, err := filepath.Abs(paths[0])
	if err != nil {
		if strings.HasPrefix(paths[0], "./") || strings.HasPrefix(paths[0], ".\\") {
			return "./" + filepath.Join(paths...)
		}
	}
	paths[0] = abs
	return filepath.Join(paths...)
}

//IsExist 文件是否存在
func IsExist(path string) bool {
	if 0 == len(path) {
		return false
	}
	_, err := os.Stat(path)
	return err == nil
}

//IsHightWin10 ...
func IsHightWin10() bool {
	major, _, _ := RtlGetNtVersionNumbers()
	return major >= 10
}
