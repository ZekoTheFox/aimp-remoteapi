package aimpremoteapi

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/zekothefox/aimp-remoteapi/pkgs/internal"
	"golang.org/x/sys/windows"
)

var blank = AIMPTrackInfo{
	Title:    "Unavailable",
	Album:    "Unavailable",
	FileName: "Unavailable",
	Genre:    "Unavailable",
	Date:     "Unavailable",
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
		fmt.Println("AIMP-RemoteApi error on opening file mapping:", err.Error())
		return nil, errors.New("Unable to open AIMP's remote file map. (" + err.Error() + ")")
	}

	view, _, _ := internal.MapViewOfFile.Call(handle, windows.FILE_MAP_READ, 0, 0, uintptr(AIMPRemoteAccessMapFileSize))
	if view == 0 {
		fmt.Println("AIMP-RemoteApi error on file mapping handle:", err.Error())
		return nil, errors.New("Unable to view file mapping. (" + err.Error() + ")")
	}

	defer cleanup(handle, view)

	// the result stored in `view` is the address to where the file map is
	// unsafe.Pointer tells us that this usage probably isn't correct,
	//   though i believe this is pretty much how it works in c/c++
	rawFileInfo := *(*[AIMPRemoteAccessMapFileSize]byte)(unsafe.Pointer(view))
	fileInfo := unpackFileInfo(rawFileInfo[:])

	// the order is always the same from what i've found
	lengths := []int{
		int(fileInfo.AlbumLength),
		int(fileInfo.ArtistLength),
		int(fileInfo.DateLength),
		int(fileInfo.FileNameLength),
		int(fileInfo.GenreLength),
		int(fileInfo.TitleLength),
	}

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
			// assemble uint16 out of 2 bytes since it'll need to be decoded as a utf16 string
			char := (uint16(data[offset+w]) << 8) + uint16(data[offset+w+1])
			widestring = append(widestring, char)
		}

		str := windows.UTF16ToString(widestring)
		values = append(values, str)
	}

	// there probably is a better way of doing this
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
