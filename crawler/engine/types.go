package engine

// 请求结构体
type Request struct {
	Url        string                            // 需要解析的url
	ParserFunc func(contents []byte) ParseResult // 不同url有不同的解析器
}

// 解析结果结构体
type ParseResult struct {
	Requests []Request     // 请求slice
	Items    []interface{} // 结果数组
}

// 返回空的Parser
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
