go build --buildmode=c-shared -ldflags="-s -w"  -o ./main.dll main.go
upx.exe -9 ./main.dll