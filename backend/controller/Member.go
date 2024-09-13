package controller

import (
	"SA-67-SongThor-SUT/config"
	"SA-67-SongThor-SUT/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMember(c *gin.Context) {
	memberID := c.Param("member_id")
	var member []entity.Member// สร้างตัวแปรเก็บข้อมูล Messages

	db := config.DB() // เชื่อมต่อกับฐานข้อมูล

	// ดึงข้อความทั้งหมดใน RoomChat ที่มี room_id ตามที่ระบุ
	result := db.Where("member_id = ?", memberID).Find(&member)

	// ตรวจสอบว่าพบข้อมูลหรือไม่
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	// ส่งข้อมูล Messages กลับไปในรูปแบบ JSON
	c.JSON(http.StatusOK, member)
}