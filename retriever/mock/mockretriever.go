package mock

type Retriever struct {
	Contents string
}

func (r *Retriever) Post(url string, form map[string]string) string {
	return form["contents"]
}

// 这里是值接收者，可以接受值或者指针接收者
func (r *Retriever) Get(url string) string {
	return r.Contents
}
