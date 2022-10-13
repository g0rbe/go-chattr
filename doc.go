/*
Package go_chattr is a package for change attribute of a file on Linux, similar to the chattr command.

Example to set the immutable attribute to a file:

	file, err := os.OpenFile("file.txt", os.O_RDONLY, 0666)

	if err != nil {
	    panic(err)
	}

	defer file.Close()

	err = chattr.SetAttr(file, chattr.FS_IMMUTABLE_FL)

	if err != nil {
	    panic(err)
	}
*/
package chattr
