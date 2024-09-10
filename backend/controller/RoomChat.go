package controller

import (
	"SA-67-SongThor-SUT/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// // CreateRoomChat สร้างห้องแชทใหม่หากยังไม่มีอยู่แล้ว
// func CreateRoomChat(db *gorm.DB, memberID uint, sellerID uint) (*entity.RoomChat, error) {
// 	// ตรวจสอบว่าห้องแชทนี้มีอยู่แล้วหรือยัง
// var existingRoom entity.RoomChat
// result := db.Where("member_id = ? AND seller_id = ?", memberID, sellerID).First(&existingRoom)

// // ถ้ามีอยู่แล้ว return ห้องแชทเดิม
// if result.RowsAffected > 0 {
// 	return &existingRoom, nil
// }

// // ถ้าไม่มี สร้างห้องแชทใหม่
// newRoom := entity.RoomChat{
// 	MemberID: memberID,
// 	SellerID: sellerID,
// }

// // บันทึกห้องแชทใหม่ลงในฐานข้อมูล
// if err := db.Create(&newRoom).Error; err != nil {
// 	return nil, err
// }

// return &newRoom, nil
// }


// POST /products
// func CreateRoomChat(c *gin.Context) {

// 	var roomchat entity.RoomChat

// 	// bind เข้าตัวแปร product
// 	if err := c.ShouldBindJSON(&roomchat); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	db := config.DB()

// 	// ตรวจสอบว่า Seller มีอยู่ในระบบหรือไม่
// 	var member entity.Member
// 	if err := db.First(&member, roomchat.MemberID).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Seller not found"})
// 		return
// 	}

// 		// ถ้าไม่มี สร้างห้องแชทใหม่
// 	newRoom := entity.RoomChat{
// 		MemberID: memberID,
// 		SellerID: sellerID,
// 	}

// 	// บันทึก Product
// 	if err := db.Create(&p).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// โหลดข้อมูล Seller ที่เชื่อมโยงกับ Product นี้
// 	db.Preload("Seller").First(&p, p.ID)

// 	c.JSON(http.StatusCreated, gin.H{"message": "Created success", "data": p})
// }

// CreateRoomChat สร้างห้องแชทใหม่หากยังไม่มีอยู่แล้ว
func CreateRoomChat(c *gin.Context, db *gorm.DB) (*entity.RoomChat, error) {
	var request struct {
		MemberID uint `json:"member_id"`
		SellerID uint `json:"seller_id"`
	}

	// Bind JSON Request
	if err := c.BindJSON(&request); err != nil {
		return nil, err
	}

	var existingRoom entity.RoomChat
	result := db.Where("member_id = ? AND seller_id = ?", request.MemberID, request.SellerID).First(&existingRoom)

	// ถ้ามีอยู่แล้ว return ห้องแชทเดิม
	if result.RowsAffected > 0 {
		return &existingRoom, nil
	}

	// ถ้าไม่มี สร้างห้องแชทใหม่
	newRoom := entity.RoomChat{
		MemberID: request.MemberID,
		SellerID: request.SellerID,
	}

	// บันทึกห้องแชทใหม่ลงในฐานข้อมูล
	if err := db.Create(&newRoom).Error; err != nil {
		return nil, err
	}

	return &newRoom, nil
}
