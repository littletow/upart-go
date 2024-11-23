@echo off
rem package zip file
echo "start"
set zipfilename="gart.zip"
set binfilename="gart.exe"
set txtfilename="sha256.txt"
go build -o %binfilename% -ldflags="-s -w" 
certutil -hashfile %binfilename% SHA256 > %txtfilename%
zip %zipfilename% %binfilename% %txtfilename%
echo "done"