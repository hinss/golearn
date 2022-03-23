package templatemethod

import "fmt"

/**
	结构型: 模板方法模式
 */

// 模板上传
type UploadInterface interface {

	// 上传公共的字段
	UploadCommon()

	// 上传自定义的字段
	UploadOthers()
}

// 暴露出去的公共上传模板方法
func UploadTemplate(u UploadInterface) {
	u.UploadCommon()
	u.UploadOthers()
}

type CommonUploader struct {
}

func (c CommonUploader) UploadCommon() {
	fmt.Println("公共的字段上传逻辑处理中...")
}

func (c CommonUploader) UploadOthers() {
	// 钩子方法 空实现
}


type InterestUploader struct {
	CommonUploader
}

func (iu InterestUploader) UploadOthers() {
	fmt.Println("执行interest 上传逻辑")
}

type MaturityUploader struct {
	CommonUploader
}

func (iu MaturityUploader) UploadOthers() {
	fmt.Println("执行Maturity 上传逻辑")
}



