package main

import "github.com/gin-gonic/gin"

func main() {
	//userAuth := jwtgo.NewJwtGoAuth("This_is_an_key")
	//repo := mongo.NewUserRepository("mongodb://localhost:27017")
	//defer repo.Disconnect()

	r := gin.Default()
	r.GET("ping", ping)

	//artistGroup := r.Group("/artist")
	//{
	//	ctr := InitArtistController()
	//	artistGroup.POST("/getArtist", ctr.GetArtist)
	//}
	//
	//artworkGroup := r.Group("/artwork")
	//{
	//	ctr := InitArtworkController()
	//	artworkGroup.POST("/getArtwork", ctr.GetArtwork)
	//}



	err := r.Run(":9002")
	print(err)
}

//============================================
func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}