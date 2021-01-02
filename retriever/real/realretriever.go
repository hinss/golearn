package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retreiver struct {
	UsersAgent string
	TimeOut    time.Duration
}

// 指针接收者实现只能以指针方式使用
func (r *Retreiver) Get(url string) string {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	res, err := httputil.DumpResponse(resp, true)

	// close
	resp.Body.Close()

	if err != nil {
		panic(err)
	}

	return string(res)
}
