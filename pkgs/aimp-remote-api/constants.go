package aimpremoteapi

import "github.com/zekothefox/aimp-remoteapi/pkgs/internal"

// This file is heavily based off of `apiRemote.h` from the AIMP Developer SDK.
//
// The Developer SDK also includes more information on how to use things from here as well, so you
// should probably look at that too.
//
// You can find the AIMP Developer SDK at:
// https://www.aimp.ru/?do=download&os=windows&cat=sdk

// AIMP Remote API Constants
const (
	// Library specific; refers to the AIMP SDK version this is based on.
	// "vX.XX" = version, "-XXXX" = specific build number
	AIMPRemoteApiVersion string = "v5.30-2500"

	// Window class name that AIMP uses for IPC messaging
	AIMPRemoteAccessClass string = "AIMP2_RemoteInfo"
	// File map size, see: https://learn.microsoft.com/en-us/windows/win32/memory/file-mapping
	AIMPRemoteAccessMapFileSize int = 2048
)

// AIMP Remote API Message Constants
//
// Normally would follow Windows API naming conventions, i.e. `WM_...`, but screw it we ball
// (also those would just look kinda ugly to read and type)
const (
	// WM_AIMP_COMMAND
	AIMP_WMCommand = internal.WM_USER + 0x75
	// WM_AIMP_NOTIFY
	AIMP_WMNotify = internal.WM_USER + 0x76
	// WM_AIMP_PROPERTY
	AIMP_WMProperty = internal.WM_USER + 0x77

	// WM_AIMP_COPYDATA_ALBUMART_ID
	// Source header mentions to look at AIMP_RA_CMD_GET_ALBUMART command
	AIMP_WMCopyDataAlbumArtId = 0x41495043
)

/*
	AIMP Remote Api Notification Constants
*/
const (
	// AIMP_RA_PROPVALUE_GET
	AIMP_NotifyPropertyGet = 0
	// AIMP_RA_PROPVALUE_SET
	AIMP_NotifyPropertySet = 1
	// AIMP_RA_PROPERTY_MASK
	AIMP_NotifyPropertyMask = 0xFFFFFFF0

	// AIMP_RA_PROPERTY_VERSION
	// (Read-only)
	//
	// High = Version ID, Low = Build number
	AIMP_NotifyPropertyVersion = 0x10
	// AIMP_RA_PROPERTY_PLAYER_POSITION
	//
	// Get = Returns current player position in milliseconds
	// Set = Desired position in milliseconds
	AIMP_NotifyPropertyPlayerPosition = 0x20
	// AIMP_RA_PROPERTY_PLAYER_DURATION
	// (Read-only)
	//
	// Returns duration of current track in milliseconds
	AIMP_NotifyPropertyPlayerDuration = 0x30
	// AIMP_RA_PROPERTY_PLAYER_STATE
	// (Read-only)
	//
	// Returns the current player state
	// - `0` is Stopped
	// - `1` is Paused
	// - `2` is Playing
	AIMP_NotifyPropertyPlayerState = 0x40
	// AIMP_RA_PROPERTY_VOLUME
	//
	// Get = Returns current volume levels from 0 .. 100 (%, in percentage)
	// Set = Desired volume level from 0 .. 100
	//
	// Returns 0 on fail
	AIMP_NotifyPropertyVolume = 0x50
	// AIMP_RA_PROPERTY_MUTE
	//
	// Get = Returns current mute from 0 .. 1
	// Set = Desired mute state from 0 .. 1
	//
	// Returns 0 on fail
	AIMP_NotifyPropertyMute = 0x60
	// AIMP_RA_PROPERTY_TRACK_REPEAT
	//
	// Get = Returns current repeat state from 0 .. 1
	// Set = Desired repeat state from 0 .. 1
	//
	// Returns 0 on fail
	AIMP_NotifyPropertyTrackRepeat = 0x70
	// AIMP_RA_PROPERTY_TRACK_SHUFFLE
	//
	// Get = Returns player shuffle state from 0 .. 1
	// Set = Desired shuffle state from 0 .. 1
	//
	// Returns 0 on fail
	AIMP_NotifyPropertyTrackShuffle = 0x80
	// AIMP_RA_PROPERTY_RADIOCAP
	//
	// Get = Returns radio capture state from 0 .. 1
	// Set = Desired radio capture state from 0 .. 1
	//
	// Returns 0 on fail
	AIMP_NotifyPropertyRadioCapture = 0x90
	// AIMP_RA_PROPERTY_VISUAL_FULLSCREEN
	//
	// Get = Returns fullscreen visualization mode state from 0 .. 1
	// Set = Desired fullscreen state from 0 .. 1
	//
	// Returns 0 on fail
	AIMP_NotifyPropertyVisualFullscreen = 0xA0

	// AIMP_RA_CMD_BASE
	AIMP_NotifyCmdBase = 10
	// AIMP_RA_CMD_REGISTER_NOTIFY
	// LParam accepts a Window Handle, which will receive AIMP_WMNotify / WM_AIMP_NOTIFY messages
	// Source mentions to look at WM_AIMP_NOTIFY message for more information
	AIMP_NotifyCmdRegisterNotify   = AIMP_NotifyCmdBase + 1
	AIMP_NotifyCmdUnregisterNotify = AIMP_NotifyCmdBase + 2
	AIMP_NotifyCmdPlay             = AIMP_NotifyCmdBase + 3
	AIMP_NotifyCmdPlayPause        = AIMP_NotifyCmdBase + 4
	AIMP_NotifyCmdPause            = AIMP_NotifyCmdBase + 5
	AIMP_NotifyCmdStop             = AIMP_NotifyCmdBase + 6
	AIMP_NotifyCmdNext             = AIMP_NotifyCmdBase + 7
	AIMP_NotifyCmdPrevious         = AIMP_NotifyCmdBase + 8
	AIMP_NotifyCmdVisualNext       = AIMP_NotifyCmdBase + 9
	AIMP_NotifyCmdVisualPrevious   = AIMP_NotifyCmdBase + 10
	AIMP_NotifyCmdQuit             = AIMP_NotifyCmdBase + 11
	AIMP_NotifyCmdAddFiles         = AIMP_NotifyCmdBase + 12
	AIMP_NotifyCmdAddFolders       = AIMP_NotifyCmdBase + 13
	AIMP_NotifyCmdAddPlaylists     = AIMP_NotifyCmdBase + 14
	AIMP_NotifyCmdAddUrl           = AIMP_NotifyCmdBase + 15
	AIMP_NotifyCmdOpenFiles        = AIMP_NotifyCmdBase + 16
	AIMP_NotifyCmdOpenFolders      = AIMP_NotifyCmdBase + 17
	AIMP_NotifyCmdOpenPlaylists    = AIMP_NotifyCmdBase + 18
	AIMP_NotifyCmdGetAlbumArt      = AIMP_NotifyCmdBase + 19
	AIMP_NotifyCmdVisualStart      = AIMP_NotifyCmdBase + 20
	AIMP_NotifyCmdVisualStop       = AIMP_NotifyCmdBase + 21

	AIMP_NotifyBase       = 0
	AIMP_NotifyTrackInfo  = AIMP_NotifyBase + 1
	AIMP_NotifyTrackStart = AIMP_NotifyBase + 2
	AIMP_NotifyProperty   = AIMP_NotifyBase + 3
)

type AIMPRemoteFileInfo = struct {
	Deprecated1    uint32
	Active         bool
	BitRate        uint32
	Channels       uint32
	Duration       uint32
	FileSize       int64
	FileMark       uint32
	SampleRate     uint32
	TrackNumber    uint32
	AlbumLength    uint32
	ArtistLength   uint32
	DateLength     uint32
	FileNameLength uint32
	GenreLength    uint32
	TitleLength    uint32
	Deprecated2    [6]uint32
}

type AIMPTrackInfo = struct {
	Album    string
	Artist   string
	Date     string
	FileName string
	Genre    string
	Title    string
}
