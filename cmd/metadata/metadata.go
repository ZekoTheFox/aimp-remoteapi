package main

import (
	"fmt"

	aimpremoteapi "github.com/zekothefox/aimp-remoteapi/pkgs/aimp-remote-api"
)

func main() {
	fmt.Println("-- metadata.go current track info example")

	info, err := aimpremoteapi.GetCurrentTrack()
	if err != nil {
		fmt.Println("Unable to get current track due to error:", err.Error())
		return
	}

	fmt.Println("Album:", info.Album)
	fmt.Println("Artist:", info.Artist)
	fmt.Println("Date:", info.Date)
	fmt.Println("File Name:", info.FileName)
	fmt.Println("Genre:", info.Genre)
	fmt.Println("Title:", info.Title)
}
