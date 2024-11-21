package aimpremoteapi

import (
	"encoding/binary"
	"fmt"
)

// hardcoded order, because go seems to move the struct's order around
var order = map[string]int{
	"Deprecated1":    1,
	"Active":         2,
	"BitRate":        3,
	"Channels":       4,
	"Duration":       5,
	"FileSize":       6,
	"FileMark":       7,
	"SampleRate":     8,
	"TrackNumber":    9,
	"AlbumLength":    10,
	"ArtistLength":   11,
	"DateLength":     12,
	"FileNameLength": 13,
	"GenreLength":    14,
	"TitleLength":    15,
	"Deprecated2":    16,
}

// in bytes
var sizes = map[int]int{
	1:  4,
	2:  1,
	3:  4,
	4:  4,
	5:  4,
	6:  8,
	7:  4,
	8:  4,
	9:  4,
	10: 4,
	11: 4,
	12: 4,
	13: 4,
	14: 4,
	15: 4,
	16: 4 * 6,
}

func getStructSize() int {
	total := 0
	for i := range sizes {
		total += sizes[i]
	}
	return total
}

func getPackedOffset(field string) int {
	position := order[field]
	if position <= 0 {
		return 0
	}

	offset := 0
	for i := 0; i <= position-1; i++ {
		offset += sizes[i]
	}

	return offset
}

func setField(info *AIMPRemoteFileInfo, field string, parsable any) {
	if field == "Deprecated2" {
		fmt.Println("AIMP-RemoteApi Warning: Deprecated2 field parsing not implemented")
		return
	}

	size := sizes[order[field]]

	// hardcoding the size isn't ideal, but not quite sure how to deal with it "properly"
	// for now, the struct shouldn't change hugely enough that it matters though
	switch size {
	case 4:
		i, ok := parsable.(uint32)
		if !ok {
			fmt.Println("AIMP-RemoteApi failed type assertion for field:", field)
			return
		}

		switch field {
		case "Deprecated1":
			info.Deprecated1 = i
		case "BitRate":
			info.BitRate = i
		case "Channels":
			info.Channels = i
		case "Duration":
			info.Duration = i
		case "SampleRate":
			info.SampleRate = i
		case "TrackNumber":
			info.TrackNumber = i
		case "AlbumLength":
			info.AlbumLength = i
		case "ArtistLength":
			info.ArtistLength = i
		case "DateLength":
			info.DateLength = i
		case "FileNameLength":
			info.FileNameLength = i
		case "GenreLength":
			info.GenreLength = i
		case "TitleLength":
			info.TitleLength = i
		}
	case 1:
		b, ok := parsable.(bool)
		if !ok {
			fmt.Println("AIMP-RemoteApi failed type assertion for field:", field)
		}

		// only Active is a bool
		if field == "Active" {
			info.Active = b
		}
	case 8:
		i, ok := parsable.(int64)
		if !ok {
			fmt.Println("AIMP-RemoteApi failed type assertion for field:", field)
		}

		// only FileSize is an int64
		if field == "FileSize" {
			info.FileSize = i
		}
	}
}

func unpackFileInfo(rawInfo []byte) *AIMPRemoteFileInfo {
	info := AIMPRemoteFileInfo{}

	for k := range order {
		offset := getPackedOffset(k)
		size := sizes[order[k]]

		// reading directly from the file map is iffy
		// for now, its better to just query the specific data using the message cmds
		// the endianness might matter...? the structure doesn't seem to contain data that is usable
		//   outside the metadata string lengths, so not sure whats going on yet
		switch size {
		case 4:
			data := uint32(0)
			binary.Decode(rawInfo[offset:offset+size], binary.BigEndian, &data)

			setField(&info, k, data)
		case 1:
			data := uint8(0)
			binary.Decode(rawInfo[offset:offset+size], binary.LittleEndian, &data)

			setField(&info, k, data == 1)
		case 8:
			data := int64(0)
			binary.Decode(rawInfo[offset:offset+size], binary.LittleEndian, &data)

			setField(&info, k, data)
		case 24:
			// only Deprecated2 uses this effectively and idk what the field is used for
			// its also at the end, so i guess we'll just leave it as is lol
		}
	}

	return &info
}
