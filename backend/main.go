// package main

// import (
// 	"SA-67-SongThor-SUT/controller"
// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// 	"net/http"
// )

// var db *gorm.DB

// func main() {
// 	// ตั้งค่า Gin
// 	r := gin.Default()

// 	// สร้าง route สำหรับ CreateRoomChat และบันทึกข้อความลง Message
// 	r.POST("/roomchat", func(c *gin.Context) {
// 		var request struct {
// 			MemberID uint   `json:"member_id"`
// 			SellerID uint   `json:"seller_id"`
// 			Content  string `json:"content"` // รับข้อความจาก request
// 		}

// 		// Bind JSON Request
// 		if err := c.BindJSON(&request); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// สร้างห้องแชทหรือดึงห้องแชทเดิม
// 		roomChat, err := controller.CreateRoomChat(db, request.MemberID, request.SellerID)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// สร้างข้อความใหม่และบันทึกลงใน Message
// 		newMessage, err := controller.CreateMessage(db, roomChat.RoomID, request.MemberID, request.SellerID, request.Content)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save message"})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{
// 			"message":     "Chat and message created successfully",
// 			"room_chat":   roomChat,
// 			"new_message": newMessage,
// 		})
// 	})

// 	// เริ่มเซิร์ฟเวอร์
// 	r.Run(":8080")
// }

/* ---------------------------------------------------------*/

// package main

// import (
// 	"SA-67-SongThor-SUT/entity"

// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
// )

// func main() {
// 	db, err := gorm.Open(sqlite.Open("songthorsut.db"), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	db.AutoMigrate(&entity.Member{}, &entity.Seller{}, &entity.RoomChat{}, &entity.Message{})
// }

/* ---------------------------------------------------------*/

package main

import (
	"SA-67-SongThor-SUT/config" // สมมติว่ามี package นี้สำหรับการตั้งค่าการเชื่อมต่อ DB และ config อื่นๆ
	"SA-67-SongThor-SUT/controller"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const PORT = "8080"

var db *gorm.DB

func main() {
	// ตั้งค่าการเชื่อมต่อฐานข้อมูล
	config.ConnectionDB()
	config.SetupDatabase()

	// ตั้งค่า Gin
	r := gin.Default()

	// ตั้งค่า Middleware (เช่น CORS)
	r.Use(CORSMiddleware()) // ตรวจสอบว่าคุณได้สร้างฟังก์ชัน CORSMiddleware() หรือไม่

	// สร้าง route สำหรับ CreateRoomChat และบันทึกข้อความลง Message
	r.POST("/roomchat", func(c *gin.Context) {
		var request struct {
			MemberID uint   `json:"member_id"`
			SellerID uint   `json:"seller_id"`
			Content  string `json:"content"` // รับข้อความจาก request
		}

		// Bind JSON Request
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// สร้างห้องแชทหรือดึงห้องแชทเดิม
		roomChat, err := controller.CreateRoomChat(c, db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// สร้างข้อความใหม่และบันทึกลงใน Message
		newMessage, err := controller.CreateMessage(db, roomChat.RoomID, request.MemberID, request.SellerID, request.Content)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save message"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":     "Chat and message created successfully",
			"room_chat":   roomChat,
			"new_message": newMessage,
		})
	})

	// ตั้งค่าเส้นทาง API อื่นๆ
	router := r.Group("")
	{
		// router.GET("/member/:id", controller.GetMember)
		// router.GET("/roomchat", controller.ListMember)
		router.POST("/roomchat", controller.CreateRoomChat)
		// router.PATCH("/member/:id", controller.UpdateMember)
		// router.DELETE("/member/:id", controller.DeleteMember)

	}

	// ตรวจสอบสถานะการทำงานของ API
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "API RUNNING... PORT: %s", PORT)
	})

	// เริ่มเซิร์ฟเวอร์
	r.Run("localhost:" + PORT)
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