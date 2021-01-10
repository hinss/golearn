package main

import (
	"learngo/exceptionhandler/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

// 这里使用到了闭包
func errWrapper(
	handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// defer 匿名函数
		defer func() {
			// panic 遇到recover() 后会继续执行自定义的处理逻辑
			r := recover()
			if r != nil {
				// 内部log下错误
				log.Panicf("Panic: %v", r)
				// 那么这里往后执行会输出InternalServerError
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)

				return
			}
		}()

		err := handler(writer, request)
		if err != nil {
			log.Panicf("Error occured "+
				"handling request:%s",
				err.Error())

			// 如果error是自定义的userError接口类型时(userError interface 定义了error interface)
			if userError, ok := err.(userError); ok {

				http.Error(writer,
					userError.Message(),
					http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,
				http.StatusText(code),
				code)
		}
	}
}

// 自定义error接口
type userError interface {
	error
	Message() string
}

func main() {

	// web.go里做一件事情启动一个web服务器，根据url获得文件的文件服务器
	http.HandleFunc("/", errWrapper((filelisting.HandleFileListing)))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

}
