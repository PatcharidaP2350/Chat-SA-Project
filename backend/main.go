package main

import (
	"SA-67-SongThor-SUT/config"
	"SA-67-SongThor-SUT/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	const PORT = "8080" // กำหนดหมายเลขพอร์ต
	// open connection database
	config.ConnectionDB()

	// Generate databases
	config.SetupDatabase()

	r := gin.Default()

	r.Use(CORSMiddleware())

	router := r.Group("")
	{

		// Room Routes
		router.POST("/roomchat", controller.CreateRoomChat)
		router.GET("/roomchat/:room_id", controller.GetMessages)
		router.POST("/message", controller.SetMessage)
		// 	router.PATCH("/message", controller.createMessege)
		// 	router.DELETE("/users/:id", controller.DeleteUser)
		// 	// Gender Routes
		// 	router.GET("/genders", controller.ListGenders)
		// }

		r.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "API RUNNING... PORT: %s", PORT)
		})

		
	}
	r.Run("localhost:" + PORT)  // Run the server
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

