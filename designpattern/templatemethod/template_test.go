package templatemethod

import "testing"

func TestTemplateMethod(t *testing.T) {

	interestUploader := InterestUploader{}
	UploadTemplate(interestUploader)

	maturityUploader := MaturityUploader{}
	UploadTemplate(maturityUploader)


}
