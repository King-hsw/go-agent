GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -H=windowsgui" -o pc-agent.exe


GOOS=windows GOARCH=amd64 go build  -o pc-agent.exe