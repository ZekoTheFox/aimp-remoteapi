package aimpremoteapi

import (
	"errors"
	"fmt"
	"strings"
	"unsafe"

	"gitbub.com/zekothefox/aimp-remoteapi/pkgs/internal"
	"golang.org/x/sys/windows"
)

var blank = AIMPTrackInfo{
	Title:    "Unavailable",
	Album:    "Unavailable",
	FileName: "Unavailable",
	Genre:    "Unavailable",
	Date:     "Unavailable",
}

func readByte(pointer unsafe.Pointer) byte {
	b := *(*byte)(pointer) // scallop activity
	return b
}

func getString(pointer unsafe.Pointer, length int) string {
	bytestring := []byte{}

	for i := 0; i < length; i++ {
		bytestring = append(bytestring, readByte(unsafe.Add(pointer, i)))
	}

	return strings.TrimSpace(string(bytestring))
}

func cleanup(handle uintptr, view uintptr) {
	internal.UnmapViewOfFile.Call(view)
	internal.CloseHandle.Call(handle)
}

// Gets the current track's info from AIMP
// Returns a pointer to a struct that holds the current track info.
func GetCurrentTrack() (*AIMPTrackInfo, error) {
	className, _ := windows.UTF16PtrFromString(AIMPRemoteAccessClass)
	handle, _, err := internal.OpenFileMapping.Call(windows.FILE_MAP_READ, 1, uintptr(unsafe.Pointer(className)))
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		fmt.Println("AIMP-RemoteApi error on acceessing file mapping:", err.Error())
		return nil, errors.New("Unable to access AIMP IPC file map.")
	}

	view, _, _ := internal.MapViewOfFile.Call(handle, windows.FILE_MAP_READ, 0, 0, uintptr(AIMPRemoteAccessMapFileSize))
	if view == 0 {
		fmt.Println("AIMP-RemoteApi error on mapping file map handle:", err.Error())
		return nil, errors.New("Unable to map file mapping from AIMP.")
	}
	defer cleanup(handle, view)

	// the data stored as a file map is stored in the result of the map call above, which is an address
	// unsafe.Pointer normally says converting uintptr back to Pointer isn't correct, but the result isn't usable otherwise
	rawFileInfo := *(*[AIMPRemoteAccessMapFileSize]byte)(unsafe.Pointer(view))
	fileInfo := unpackFileInfo(rawFileInfo[:])

	lengths := []int{
		int(fileInfo.AlbumLength),
		int(fileInfo.ArtistLength),
		int(fileInfo.DateLength),
		int(fileInfo.FileNameLength),
		int(fileInfo.GenreLength),
		int(fileInfo.TitleLength),
	}

	fmt.Println(lengths)

	values := []string{}

	data := rawFileInfo[getStructSize():]
	for i := range lengths {
		offset := 0
		if i > 0 {
			for v := 0; v < i; v++ {
				// multiply by 2 since we're dealing with utf16 (2 bytes / 16 bits)
				offset += lengths[v] * 2
			}
		}

		widestring := []uint16{}
		for w := 2; w <= (lengths[i] * 2); w += 2 {
			char := (uint16(data[offset+w]) << 8) + uint16(data[offset+w+1])
			widestring = append(widestring, char)
		}

		str := windows.UTF16ToString(widestring)
		values = append(values, str)
	}

	file := AIMPTrackInfo{
		Album:    values[0],
		Artist:   values[1],
		Date:     values[2],
		FileName: values[3],
		Genre:    values[4],
		Title:    values[5],
	}

	return &file, nil
}
