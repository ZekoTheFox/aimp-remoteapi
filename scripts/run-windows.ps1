# assuming we're running in the parent directory
.\scripts\build-windows.ps1
echo '> built binaries, running'
# run binaries
.\.bin\readout.exe
.\.bin\metadata.exe
