package controller

import (
	"SA-67-SongThor-SUT/config"
	"SA-67-SongThor-SUT/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRoomChat(c *gin.Context) {
	var room entity.RoomChat

	// bind เข้าตัวแปร user
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := config.DB()

	var seller entity.Seller
	db.First(&seller, room.SellerID)
	if seller.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "seller not found"})
		return
	}

	r := entity.RoomChat{
		MemberID: room.MemberID,
		SellerID: room.SellerID,
		Seller:   seller,
		// RoomChatID: Room.RoomChatID,
	}

	// บันทึก
	if err := db.Create(&r).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Created success", "data": r})
}

