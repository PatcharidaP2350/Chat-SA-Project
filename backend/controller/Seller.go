package controller

import (
	"SA-67-SongThor-SUT/config"
	"SA-67-SongThor-SUT/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSeller(c *gin.Context) {
	sellerID := c.Param("id") // รับ room_id จาก URL

	var seller []entity.Seller // สร้างตัวแปรเก็บข้อมูล Messages

	db := config.DB() // เชื่อมต่อกับฐานข้อมูล

	// ดึงข้อความทั้งหมดใน RoomChat ที่มี room_id ตามที่ระบุ
	result := db.Where("id = ?", sellerID).Find(&seller)

	// ตรวจสอบว่าพบข้อมูลหรือไม่
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	// ส่งข้อมูล Messages กลับไปในรูปแบบ JSON
	c.JSON(http.StatusOK, seller)
}