cd ./frp-dev/cmd/frpc/
go build --buildmode=c-shared -ldflags="-s -w"  -o ../../../Costura64/main.dll main.go
upx.exe -9 ../../../Costura64/main.dll