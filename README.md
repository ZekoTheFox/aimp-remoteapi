# aimp-remoteapi

Library for interfacing with [AIMP](https://aimp.ru)'s IPC Info, more specifically its Remote File Track Info API.

Written in/for Go.

# Usage

AIMP needs to be open in order to send messages and access player information.

`GetCurrentTrack()` lets you access the following information from the player:

-   Track title
-   Album name
-   Artist name(s)
-   Date
-   Genre
-   File location/path

More information about the file can be queried using `GetProperty(...)`, see `constants.go` for commands list. (`AIMP_NotifyCmd*`)

Currently only tested for Windows; AIMP for Linux support is not specifically included and unknown if works.

# To-do

-   AIMP has a track change notification via messages; create an implmentation for it

# License

This repository is offered under the GNU LGPLv3 License.

See `LICENSE` for more information regarding the licensing of the code in this repository.
