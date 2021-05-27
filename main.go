package main

import (
	"embed"
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
	"time"
	//"log"
	"net/http"
	//导入echo包
	"github.com/labstack/echo"
)

type Files struct {
	Name    string    `json:"name" xml:"name"`
	IsDir   bool      `json:"is_dir" xml:"is_dir"`
	Url     string    `json:"url" xml:"url"`
	Size    int64     `json:"size" xml:"size"`
	ModTime time.Time `json:"mod_time" xml:"mod_time"`
}

// 获取文件列表
func listFiles(dirname string) []Files {
	var _files []Files
	fileInfos, err := ioutil.ReadDir("./" + dirname)
	if err != nil {
		//log.Fatal(err)
		errors.New("目录错误")
	}
	for _, fi := range fileInfos {
		_files = append(_files, Files{
			Name:    fi.Name(),
			IsDir:   fi.IsDir(),
			Url:     dirname + "/" + fi.Name(),
			Size:    fi.Size(),
			ModTime: fi.ModTime(),
		})
	}
	return _files
}

//go:embed assets
var f_assets embed.FS

//go:embed template/index.html
var index_content []byte

func main() {
	//实例化echo对象。
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.Blob(http.StatusOK, "text/html", index_content)
	})
	e.GET("/files", func(c echo.Context) error {
		name := c.FormValue("dir")
		name = strings.Replace(name, "../", "", -1)
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusOK)
		return json.NewEncoder(c.Response()).Encode(listFiles(name))
	})
	e.GET("/assets/*", echo.WrapHandler(http.FileServer(http.FS(f_assets))))
	//e.GET("/css/*", echo.WrapHandler(http.FileServer(http.FS(f_css))))
	e.Static("/file", "./")

	//启动http server, 并监听8080端口，冒号（:）前面为空的意思就是绑定网卡所有Ip地址，本机支持的所有ip地址都可以访问。
	e.Start(":8080")
}
