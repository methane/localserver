all:
	GOOS=windows GOARCH=386 go build -o localserver.exe
	zip localserver-win32.zip localserver.exe
