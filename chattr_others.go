//go:build windows
// +build windows

package chattr

import "os"

const ATTR_SUPPORTED = false

/*
GetAttrs retrieves the attributes of a file.
*/
func GetAttrs(f *os.File) (int32, error) {
	return 0, ErrOSNotSupported
}

/*
SetAttr sets the given attribute.
*/
func SetAttr(f *os.File, attr int32) error {
	return ErrOSNotSupported
}

/*
UnsetAttr unsets the given attribute.
*/
func UnsetAttr(f *os.File, attr int32) error {
	return ErrOSNotSupported
}

/*
IsAttr checks whether the given attribute is set.
*/
func IsAttr(f *os.File, attr int32) (bool, error) {
	return false, ErrOSNotSupported
}
