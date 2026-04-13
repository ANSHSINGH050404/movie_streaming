package main

import ("fmt"
 "github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Failed to start server: ", err)
	}
}