package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

//type userError string

type userError struct {
	msg string
}

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return e.msg
}

func HandleFileListing(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError{msg: "path must start " +
			"with " + prefix}
	}
	// 1. 拿到/list/后面的文件名 /list/fib.txt
	path := request.URL.Path[len(prefix):]
	// 2. 调用系统函数打开文件
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// 3.读出所有的内容 返回的是[]byte
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	// 4.通过resWriter写回到浏览器中
	writer.Write(all)
	return nil
}
