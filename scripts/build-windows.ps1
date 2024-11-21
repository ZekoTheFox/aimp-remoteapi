echo '> building cmd/readout/readout.go'
go build -o .bin\readout.exe cmd\readout\readout.go

echo '> building cmd/metadata/metadata.go'
go build -o .bin\metadata.exe cmd\metadata\metadata.go
