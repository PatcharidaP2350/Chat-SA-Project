package controller

import (
	"SA-67-SongThor-SUT/config"
	"SA-67-SongThor-SUT/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateMessage สร้างข้อความใหม่ใน RoomChat
func CreateMessage(db *gorm.DB, roomChatID uint, memberID uint, sellerID uint, content string) (*entity.Message, error) {
	newMessage := entity.Message{
		RoomChatID: roomChatID,
		MemberID:   memberID,
		SellerID:   sellerID,
		Content:    content,
	}

	// บันทึกข้อความลงในฐานข้อมูล
	if err := db.Create(&newMessage).Error; err != nil {
		return nil, err
	}

//		return &newMessage, nil
//	}
func GetMessages(c *gin.Context) {
	roomID := c.Param("room_id") // รับ room_id จาก URL

	var messages []entity.Message // สร้างตัวแปรเก็บข้อมูล Messages

	db := config.DB() // เชื่อมต่อกับฐานข้อมูล

	// ดึงข้อความทั้งหมดใน RoomChat ที่มี room_id ตามที่ระบุ
	result := db.Where("room_chat_id = ?", roomID).Find(&messages)

	// ตรวจสอบว่าพบข้อมูลหรือไม่
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	// ส่งข้อมูล Messages กลับไปในรูปแบบ JSON
	c.JSON(http.StatusOK, messages)
}

func SetMessage(c *gin.Context) {
	var request struct {
		RoomChatID uint   `json:"room_chat_id"` // รหัสของห้องแชท
		Content    string `json:"content"`      // เนื้อหาของข้อความ
		MemberID   uint   `json:"member_id"`    // รหัสของผู้ส่งข้อความ
	}

	// Bind JSON Request เพื่อรับข้อมูลจาก Body ของ request
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// สร้างข้อความใหม่
	newMessage := entity.Message{
		RoomChatID: request.RoomChatID,
		Content:    request.Content,
		MemberID:   request.MemberID,
	}

	// เชื่อมต่อฐานข้อมูล
	db := config.DB()

	// บันทึกข้อความใหม่ลงในตาราง Message
	if err := db.Create(&newMessage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save message"})
		return
	}

	// ส่งข้อมูลข้อความที่ถูกบันทึกกลับไป
	c.JSON(http.StatusOK, gin.H{
		"message": "Message saved successfully",
		"data":    newMessage,
	})
}
