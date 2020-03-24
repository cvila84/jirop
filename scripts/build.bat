@echo off
fyne bundle ..\assets\jira.png > ..\pkg\fyne\icon\jira.go
rsrc -ico ..\assets\jira.ico -o ..\cmd\jirop\jirop.syso
set CC=C:\Programs\Cygwin\bin\x86_64-w64-mingw32-gcc.exe
pushd ..\cmd\jirop
go install -ldflags -H=windowsgui
popd
