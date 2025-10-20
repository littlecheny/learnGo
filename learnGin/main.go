package main

import(
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	type student struct{
		Name string
		Age int8
	}

	r.LoadHTMLGlob("templates/*")

	stu1 := &student{Name: "Geektut", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/arr", func(c *gin.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gin.H{
			"title": "Gin",
			"stuArr": [2]*student{stu1, sut2},
		})
	})

	r.Run(":8080")
}