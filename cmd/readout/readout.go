package main

import (
	"fmt"

	aimpremoteapi "gitbub.com/zekothefox/aimp-remoteapi/pkgs/aimp-remote-api"
)

func main() {
	fmt.Println("-- readout.exe aimp remote api example")

	properties := map[string]uint{
		"Duration":              aimpremoteapi.AIMP_NotifyPropertyPlayerDuration,
		"Position":              aimpremoteapi.AIMP_NotifyPropertyPlayerPosition,
		"Player State":          aimpremoteapi.AIMP_NotifyPropertyPlayerState,
		"Volume":                aimpremoteapi.AIMP_NotifyPropertyVolume,
		"Mute":                  aimpremoteapi.AIMP_NotifyPropertyMute,
		"Track Repeat":          aimpremoteapi.AIMP_NotifyPropertyTrackRepeat,
		"Track Shuffle":         aimpremoteapi.AIMP_NotifyPropertyTrackShuffle,
		"Radio Capture":         aimpremoteapi.AIMP_NotifyPropertyRadioCapture,
		"Visualizer Fullscreen": aimpremoteapi.AIMP_NotifyPropertyVisualFullscreen,
	}

	for key, value := range properties {
		property, err := aimpremoteapi.GetProperty(value)
		if err != nil {
			fmt.Println("Failed to get", key, "property:", err.Error())
			continue
		}

		fmt.Println(key+":", property)
	}
}
