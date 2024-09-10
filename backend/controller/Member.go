// package controller

// import (
// 	"SA-67-SongThor-SUT/entity"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// // CreateText รับข้อความจาก Member และเก็บไว้ในฐานข้อมูล
// func CreateText(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var request struct {
// 			MemberID    uint   `json:"member_id"`
// 			SellerID    uint   `json:"seller_id"`
// 			TextDisplay string `json:"text_display"`
// 		}

// 		// bind ข้อมูล JSON เข้าตัวแปร request
// 		if err := c.ShouldBindJSON(&request); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// ตรวจสอบว่า MemberID มีอยู่ในฐานข้อมูล
// 		var member entity.Member
// 		if err := db.First(&member, request.MemberID).Error; err != nil {
// 			c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
// 			return
// 		}

// 		// ตรวจสอบว่า SellerID มีอยู่ในฐานข้อมูล
// 		var seller entity.Seller
// 		if err := db.First(&seller, request.SellerID).Error; err != nil {
// 			c.JSON(http.StatusNotFound, gin.H{"error": "Seller not found"})
// 			return
// 		}

// 		// ตรวจสอบว่า MemberID ไม่ใช่ SellerID
// 		if request.MemberID == seller.MemberID {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Member cannot be the seller"})
// 			return
// 		}

// 		// ค้นหาห้องแชทที่มีอยู่แล้ว
// 		var roomChat entity.RoomChat
// 		result := db.Where("member_id = ? AND seller_id = ?", request.MemberID, request.SellerID).First(&roomChat)
// 		if result.RowsAffected == 0 {
// 			c.JSON(http.StatusNotFound, gin.H{"error": "Room chat not found"})
// 			return
// 		}

// 		// สร้างข้อความใหม่
// 		newText := entity.TextDisplay{
// 			RoomID:   roomChat.RoomID,
// 			MemberID: request.MemberID,
// 			SellerID: request.SellerID,
// 			Content:  request.TextDisplay,
// 		}

// 		// บันทึกข้อความลงในฐานข้อมูล
// 		if err := db.Create(&newText).Error; err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save text"})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"message": "Text created successfully", "data": newText})
// 	}
// }

package controller

import (
	"SA-67-SongThor-SUT/config"
	"SA-67-SongThor-SUT/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST /users
func CreateMember(c *gin.Context) {
	var user entity.Member

	// bind เข้าตัวแปร user
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB()

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
	// 	return
	// }

	// สร้าง User
	m := entity.Member{
		Username:    user.Username,
		Password:    user.Password,
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		PicProfile:  user.PicProfile,
	}

	// บันทึก
	if err := db.Create(&m).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Created success", "data": m})
}

// GET /members/:id
func GetMember(c *gin.Context) {
	ID := c.Param("id")
	var member entity.Member

	db := config.DB()
	result := db.Preload("Seller").First(&member, ID) // Preload Seller related to Member
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, member)
}

// GET /users
func ListMember(c *gin.Context) {

	var user []entity.Member

	db := config.DB()
	results := db.Preload("Seller").Find(&user)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DELETE /members/:id
func DeleteMember(c *gin.Context) { // ลบข้อมูลสมาชิกตาม ID
	id := c.Param("id")
	db := config.DB()
	if tx := db.Delete(&entity.Member{}, id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})
}

// PATCH /members/:id
func UpdateMember(c *gin.Context) { // อัพเดทข้อมูลสมาชิกตาม ID
	var member entity.Member

	MemberID := c.Param("id")

	db := config.DB()
	result := db.First(&member, MemberID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	if err := db.Model(&member).Updates(member).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update member"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})
}
