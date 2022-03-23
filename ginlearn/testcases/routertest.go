package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Path variable
func paramString() {

	router := gin.Default()
	router.GET("/user/:name", func(c *gin.Context) {

		name := c.Param("name")

		c.String(200, "get name: %s", name)
	})

	router.Run(":8080")
}

func queryString() {

	router := gin.Default()
	router.GET("/welcome", func(c *gin.Context) {

		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")

		c.String(200, "firstname: %s lastname: %s", firstname, lastname)
	})

	router.Run(":8080")
}

func formDataPost() {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status": "posted",
			"message": message,
			"nick": nick,
		})
	})

	router.Run(":8080")
}

func queryPostForm() {

	router := gin.Default()
	router.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s, message: %s", id, page, name, message)
	})

	router.Run(":8080")
}

func uploadFile(dir string) {

	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8M
	router.POST("/upload", func(c *gin.Context) {

		file, _ := c.FormFile("file")
		fmt.Println(file.Filename)

		dst := fmt.Sprintf("%s/%s", dir, file.Filename)
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			fmt.Println(err)
		}

		c.String(200, fmt.Sprintf("'%s' uploa!", file.Filename))
	})

	router.Run(":8080")
}

func groupRouterWithCustomRecovery() {

	r := gin.New()
	r.Use(gin.Logger())

	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(500, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(500)
	}))

	r.GET("/panic", func(c *gin.Context) {
		panic("foo")
	})

	r.GET("/", func(c *gin.Context) {
		c.String(200, "ohai")
	})

	r.Run(":8080")

}

func AuthRequired() func(ctx *gin.Context) {

	return func(ctx *gin.Context) {
		fmt.Println("mock auth required middleware")
	}
}



func main() {

	//queryString()

	//表单提交
	//formDataPost()

	//queryPostForm()

	//dir := "/Users/zhanghaoxuan/goproject/src/learngo/ginlearn/testcases"
	//uploadFile(dir)

	groupRouterWithCustomRecovery()

}
