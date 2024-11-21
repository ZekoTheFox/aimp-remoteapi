package internal

import "golang.org/x/sys/windows"

var (
	user32Dll   = windows.NewLazySystemDLL("user32.dll")
	kernel32Dll = windows.NewLazySystemDLL("kernel32.dll")

	// user32.dll FindWindowW
	FindWindow = user32Dll.NewProc("FindWindowW")
	// user32.dll SendMessageW
	SendMessage = user32Dll.NewProc("SendMessageW")

	// kernel32.dll OpenFileMappingW
	OpenFileMapping = kernel32Dll.NewProc("OpenFileMappingW")
	// kernel32.dll MapViewOfFile
	MapViewOfFile = kernel32Dll.NewProc("MapViewOfFile")
	// kernel32.dll UnmapViewOfFile
	UnmapViewOfFile = kernel32Dll.NewProc("UnmapViewOfFile")
	// kernel32.dll CloseHandle
	CloseHandle = kernel32Dll.NewProc("CloseHandle")
)

// see https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-user
const (
	WM_USER = 0x0400
)
