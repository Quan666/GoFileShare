# go文件分享服务器


## 打包

go build -trimpath -ldflags "-s -w" httpServer.go
upx httpServer.exe