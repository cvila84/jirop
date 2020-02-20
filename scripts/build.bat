@echo off
REM pkg/fyne/icon/jira.go is generated with fyne bundle jira.png > jira.go
REM jirop.syso is generated with rcsc -icon jira.ico
set CC=C:\Programs\Cygwin\bin\x86_64-w64-mingw32-gcc.exe
pushd cmd\jirop
go build -ldflags -H=windowsgui -o ..\..\jirop.exe
popd
