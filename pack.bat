@echo off 
rem package zip file
echo "start"
set zipfilename="gart.zip"
set binfilename="gart.exe"
set txtfilename="sha256.txt"
go build -o %binfilename% -ldflags="-s -w" 
certutil -hashfile %binfilename% SHA256 > %txtfilename%

set target=''
set line=0
setlocal enabledelayedexpansion
for /f %%a in (sha256.txt) do (
    set /a line+=1
    if !line! == 2 (
        set target=%%a
        echo !target! > %txtfilename%
    )
)
zip %zipfilename% %binfilename% %txtfilename%
echo "done"