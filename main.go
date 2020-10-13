package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()


	err := r.Run(":9002")
	print(err)


}