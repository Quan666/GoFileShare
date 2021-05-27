# go文件分享

将打包好的程序放到需要分享的文件目录运行即可

## 打包

go build -trimpath -ldflags "-s -w" GoFileShare.go

upx GoFileShare.exe