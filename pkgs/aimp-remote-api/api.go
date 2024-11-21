package aimpremoteapi

import (
	"errors"
	"fmt"
	"unsafe"

	"gitbub.com/zekothefox/aimp-remoteapi/pkgs/internal"
	"golang.org/x/sys/windows"
)

func getRemoteWindow() uintptr {
	className, _ := windows.UTF16PtrFromString(AIMPRemoteAccessClass)
	window, _, err := internal.FindWindow.Call(uintptr(unsafe.Pointer(className)), 0)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		fmt.Println("AIMP-RemoteApi Error:", err.Error())
		return 0
	}

	return window
}

// Get a property's current value from AIMP
// See constants under `AIMP_NotifyProperty...` for available properties.
//
// Returns `0` if there an error, but the property may also return `0` as well
func GetProperty(property uint) (uint, error) {
	window := getRemoteWindow()

	result, _, err := internal.SendMessage.Call(window, AIMP_WMProperty, uintptr(property|AIMP_NotifyPropertyGet), 0)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		fmt.Println("AIMP-RemoteApi error when sending Get message:", err.Error())
		return 0, errors.New("Failed to send message to Remote API window. (" + err.Error() + ")")
	}

	return uint(result), nil
}

// Set a property's value in AIMP
// See constants under `AIMP_NotifyProperty...` for available properties.
func SetProperty(property uint, value int) {
	window := getRemoteWindow()

	_, _, err := internal.SendMessage.Call(window, AIMP_WMProperty, uintptr(property|AIMP_NotifyPropertySet), uintptr(value))
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		fmt.Println("AIMP-RemoteApi error when sending Set message:", err.Error())
	}
}
