package main

import "github.com/gin-gonic/gin"

func main() {

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



	err := r.Run(":9000")
	print(err)


}