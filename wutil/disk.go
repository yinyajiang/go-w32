package wutil

import (
	"path/filepath"

	"github.com/yinyajiang/go-w32"
	tools "github.com/yinyajiang/go-ytools/utils"
)

//DiskStatus 磁盘信息
type DiskStatus struct {
	All  uint64
	Used uint64
	Free uint64
}

//DiskUsage 获取路径的磁盘信息
func DiskUsage(path string) (disk DiskStatus, bret bool) {
	vol := filepath.VolumeName(tools.AbsPath(path))
	b, freeBytesAvailable, totalNumberOfBytes, totalNumberOfFreeBytes := w32.GetDiskFreeSpaceEx(vol)

	return DiskStatus{
		All:  totalNumberOfBytes,
		Used: totalNumberOfBytes - totalNumberOfFreeBytes,
		Free: freeBytesAvailable,
	}, b
}

//GetMaxPartition 获取最大分区
func GetMaxFreePartition() (ret string) {
	ps := w32.GetLogicalDriveStrings()
	lastFree := uint64(0)
	for _, p := range ps {
		t := w32.GetDriveType(p)
		if t != w32.DRIVE_FIXED {
			continue
		}
		usage, b := DiskUsage(p)
		if !b {
			continue
		}
		if usage.Free > lastFree {
			ret = p
			lastFree = usage.Free
		}
	}
	return
}
