//go:build !windows
// +build !windows

package chattr

import (
	"os"
	"syscall"
	"unsafe"
)

const ATTR_SUPPORTED = true

/*
Request flags.
*/
const (
	// from ioctl_list manpage
	FS_IOC_GETFLAGS uintptr = 0x80086601
	FS_IOC_SETFLAGS uintptr = 0x40086602
)

func ioctl(f *os.File, request uintptr, attrp *int32) error {

	argp := uintptr(unsafe.Pointer(attrp))

	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), request, argp)

	if errno != 0 {
		return os.NewSyscallError("ioctl", errno)
	}

	return nil
}

/*
GetAttrs retrieves the attributes of a file.
*/
func GetAttrs(f *os.File) (int32, error) {

	attr := int32(-1)

	err := ioctl(f, FS_IOC_GETFLAGS, &attr)

	return attr, err
}

/*
SetAttr sets the given attribute.
*/
func SetAttr(f *os.File, attr int32) error {

	attrs, err := GetAttrs(f)

	if err != nil {
		return err
	}

	attrs |= attr

	return ioctl(f, FS_IOC_SETFLAGS, &attrs)

}

/*
UnsetAttr unsets the given attribute.
*/
func UnsetAttr(f *os.File, attr int32) error {

	attrs, err := GetAttrs(f)

	if err != nil {
		return err
	}

	attrs ^= (attrs & attr)

	return ioctl(f, FS_IOC_SETFLAGS, &attrs)
}

/*
IsAttr checks whether the given attribute is set.
*/
func IsAttr(f *os.File, attr int32) (bool, error) {

	attrs, err := GetAttrs(f)

	if err != nil {
		return false, err
	}

	if (attrs & attr) != 0 {
		return true, nil
	} else {
		return false, nil
	}

}
