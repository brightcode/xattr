package xattr

func getxattr(path string, name string, value *byte, size int, pos int, options int) (n int, errno int) {
	r0, _, e1 := Syscall6(SYS_GETXATTR, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(StringBytePtr(name))), uintptr(unsafe.Pointer(value)), uintptr(size), uintptr(pos), uintptr(options))
	n = int(r0)
	errno = int(e1)
	return
}

func listxattr(path string, namebuf *byte, size int, options int) (n int, errno int) {
	r0, _, e1 := Syscall6(SYS_LISTXATTR, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(namebuf)), uintptr(size), uintptr(options), 0, 0)
	n = int(r0)
	errno = int(e1)
	return
}

func setxattr(path string, name string, value *byte, size int, pos int, options int) (errno int) {
	_, _, e1 := Syscall6(SYS_SETXATTR, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(StringBytePtr(name))), uintptr(unsafe.Pointer(value)), uintptr(size), uintptr(pos), uintptr(options))
	errno = int(e1)
	return
}

func removexattr(path string, name string, options int) (errno int) {
	_, _, e1 := Syscall(SYS_REMOVEXATTR, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(StringBytePtr(name))), uintptr(options))
	errno = int(e1)
	return
}