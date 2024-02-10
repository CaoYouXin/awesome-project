@echo off

cd ../

go generate
go build -ldflags "-H windowsgui" -o "awesome.exe"
