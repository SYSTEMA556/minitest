package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// ★ r をちゃんと使う
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

/*
 * ──↓ もしお試しコードを残したいなら↓──
 *       ※ func main() ではなく別名に
 */

func DemoHello() {
	println("Hello, sample!")
}
